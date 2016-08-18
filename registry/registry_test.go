package registry_test

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/king-jam/gossdp"
	libreg "github.com/king-jam/libreg/registry"
	"github.com/king-jam/libreg/registry/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/skunkworxs/nm-registry"
)

var _ = Describe("Registry", func() {

	var (
		backendAddr string
		datacenter  string
	)
	BeforeEach(func() {
		backendAddr = "127.0.0.1:8500"
		datacenter = "dc-test"
	})

	Describe("NewRegistry", func() {
		var (
			badBackend libreg.Backend
		)

		It("should return a Registry struct using mock backend", func() {
			var r *registry.Registry

			mock.Register()
			r, _ = registry.NewRegistry(libreg.MOCK, datacenter, backendAddr)
			Expect(r).To(BeAssignableToTypeOf(&registry.Registry{}))
		})
		It("should reject invalid backends", func() {
			badBackend = "BADBACKEND"

			mock.Register()
			_, err := registry.NewRegistry(badBackend, datacenter, backendAddr)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("SSDP URN Management", func() {

		It("should add a new URN to the list", func() {
			var r *registry.Registry
			mock.Register()
			r, _ = registry.NewRegistry(libreg.MOCK, datacenter, backendAddr)

			r.RemoveSearchTerm("URN:DUMMY:INFO")
			r.AddSearchTerm("URN:DUMMY:INFO", "TEST DATA")
			st := r.GetSearchTerms()
			Expect(st["URN:DUMMY:INFO"]).To(Equal("TEST DATA"))
		})

		It("should remove a URN from the list", func() {
			var r *registry.Registry
			mock.Register()
			r, _ = registry.NewRegistry(libreg.MOCK, datacenter, backendAddr)

			r.AddSearchTerm("URN:DUMMY:INFO", "TEST DATA")
			r.RemoveSearchTerm("URN:DUMMY:INFO")
			st := r.GetSearchTerms()
			Expect(st["URN:DUMMY:INFO"]).To(Equal(""))
		})
	})

	Describe("Service Registration", func() {
		var (
			r               *registry.Registry
			timeout         time.Duration
			pollingInterval time.Duration
			urn             string
			ip              string
			rGen            *rand.Rand
			uuid            string
			err             error
			ssdp            *gossdp.Ssdp
			server          gossdp.AdvertisableServer
			serviceName     string
			serviceTag      string
		)

		BeforeEach(func() {
			// Timeout of 5s because default 1s isn't long enough for SSDP message to be seen
			timeout = time.Second * 5

			// Poll at 500ms because default 10ms is too fast to be useful
			pollingInterval = time.Millisecond * 500

			mock.Register()
			r, _ = registry.NewRegistry(libreg.MOCK, datacenter, backendAddr)
			urn = "urn:schemas-upnp-org:service:agent:0.1"

			ip = "192.168.1.1:65535"
			rGen = rand.New(rand.NewSource(time.Now().UnixNano()))
			uuid = strconv.Itoa(rGen.Int())

			ssdp, err = gossdp.NewSsdp(nil)
			if err != nil {
				log.Printf("err: %s\n", err)
			}
			Expect(err).NotTo(HaveOccurred())

			server = gossdp.AdvertisableServer{
				ServiceType: urn,
				DeviceUuid:  uuid,
				Location:    fmt.Sprintf("%s%s%s", "http://", ip, "/fakepath"),
				MaxAge:      2,
			}
			serviceName = "Inservice-service:agent:0.1"
			serviceTag = "Inservice"
		})

		It("should successfully register an Inservive-Agent service", func() {
			ssdp.AdvertiseServer(server)
			go ssdp.Start()

			// Start the test and wait for SSDP message to be seen
			r.AddSearchTerm(urn, serviceTag)
			go r.Run()

			Eventually(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceName] != nil
			}, timeout, pollingInterval).Should(BeTrue())

			Consistently(func() bool {
				service, _ := r.Store.Service(serviceName, "", nil)
				return service[0].ServiceName == serviceName
			}, timeout, pollingInterval).Should(BeTrue())

			ssdp.Stop()
			r.Stop()
			log.Println("Finished test: Register IA successfully")

		})

		It("should handle badly formed URNs if they're whitelisted", func() {
			serviceName = "Inservice-urn:MalformedData"
			urn = "urn:MalformedData"
			server.ServiceType = urn

			ssdp.AdvertiseServer(server)
			go ssdp.Start()

			// Start the test and wait for SSDP message to be seen
			r.AddSearchTerm(urn, serviceTag)
			go r.Run()

			Eventually(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceName] != nil
			}, timeout, pollingInterval).Should(BeTrue())

			Consistently(func() bool {
				service, _ := r.Store.Service(serviceName, "", nil)
				return service[0].ServiceName == serviceName
			}, timeout, pollingInterval).Should(BeTrue())

			r.Stop()
			ssdp.Stop()
			log.Println("Finished test: Handle Malformed URN")

		})

		It("should ignore SSDP messages with unknown URNs", func() {

			ssdp.AdvertiseServer(server)
			go ssdp.Start()

			// Start the test and wait for SSDP message to be seen
			urn = "urn:BAD:URN"
			r.AddSearchTerm(urn, serviceTag)
			go r.Run()

			Consistently(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceName] == nil
			}, timeout, pollingInterval).Should(BeTrue())

			r.Stop()
			ssdp.Stop()
			log.Println("Finished test: Ignore unknown URN")

		})

		It("should fail to parse SSDP messages with a bad IP address", func() {

			badIP := "192.168.1165535"
			badScheme := "://"

			server.Location = fmt.Sprintf("%s%s%s", "http://", badIP, "/fakepath")
			ssdp.AdvertiseServer(server)

			noSchemeServer := server
			noSchemeServer.Location = fmt.Sprintf("%s%s%s", badScheme, ip, "/fakepath")
			ssdp.AdvertiseServer(noSchemeServer)
			go ssdp.Start()

			r.AddSearchTerm(urn, serviceTag)

			// Start the test and wait for SSDP message to be seen
			go r.Run()

			Consistently(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceName] == nil
			}, timeout, pollingInterval).Should(BeTrue())

			r.Stop()
			ssdp.Stop()
			log.Println("Finished test: Fail to parse bad IP")

		})

		It("should fail to parse SSDP messages with a bad port", func() {

			ip := "192.168.1.1:INVALIDPORT"
			server.Location = fmt.Sprintf("%s%s%s", "http://", ip, "/fakepath")
			ssdp.AdvertiseServer(server)
			go ssdp.Start()

			// Start the test and wait for SSDP message to be seen
			r.AddSearchTerm(urn, serviceTag)
			go r.Run()

			Consistently(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceName] == nil
			}, timeout, pollingInterval).Should(BeTrue())

			r.Stop()
			ssdp.Stop()
			log.Println("Finished test: Fail to parse bad port")

		})

		It("should fail to register SSDP messages with no Node UUID", func() {

			server.DeviceUuid = ""
			ssdp.AdvertiseServer(server)
			go ssdp.Start()

			// Start the test and wait for SSDP message to be seen
			r.AddSearchTerm(urn, serviceTag)
			go r.Run()

			Consistently(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceName] == nil
			}, timeout, pollingInterval).Should(BeTrue())

			r.Stop()
			ssdp.Stop()
			log.Println("Finished test: Fail to register message without UUID")

		})
	})
})
