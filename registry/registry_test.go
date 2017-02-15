package registry_test

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	libreg "github.com/RackHD/neighborhood-manager/libreg/registry"
	"github.com/RackHD/neighborhood-manager/libreg/registry/mock"
	"github.com/RackHD/neighborhood-manager/registry"
	"github.com/king-jam/gossdp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
			urnIS           string
			urnRHD          string
			ipRHD           string
			ipIS            string
			rGen            *rand.Rand
			uuidRHD         string
			uuidIS          string
			err             error
			ssdp            *gossdp.Ssdp
			serverIS        gossdp.AdvertisableServer
			serviceNameIS   string
			serviceTagIS    string
			serverRHD       gossdp.AdvertisableServer
			serviceNameRHD  string
			serviceTagRHD   string
		)

		BeforeEach(func() {
			// Timeout of 5s because default 1s isn't long enough for SSDP message to be seen
			timeout = time.Second * 5

			// Poll at 500ms because default 10ms is too fast to be useful
			pollingInterval = time.Millisecond * 500

			mock.Register()
			r, _ = registry.NewRegistry(libreg.MOCK, datacenter, backendAddr)
			urnIS = "urn:schemas-upnp-org:service:agent:0.1"

			ipIS = "192.168.1.1:65535"
			rGen = rand.New(rand.NewSource(time.Now().UnixNano()))
			uuidIS = strconv.Itoa(rGen.Int())

			urnRHD = "urn:schemas-upnp-org:service:api:2.0"

			ipRHD = "192.168.1.1:65535"
			uuidRHD = strconv.Itoa(rGen.Int())

			ssdp, err = gossdp.NewSsdp(nil)
			if err != nil {
				log.Printf("err: %s\n", err)
			}
			Expect(err).NotTo(HaveOccurred())

			serverIS = gossdp.AdvertisableServer{
				ServiceType: urnIS,
				DeviceUuid:  uuidIS,
				Location:    fmt.Sprintf("%s%s%s", "http://", ipIS, "/fakepath"),
				MaxAge:      2,
			}
			serviceNameIS = "Inservice-service:agent:0.1"
			serviceTagIS = "Inservice"

			serverRHD = gossdp.AdvertisableServer{
				ServiceType: urnRHD,
				DeviceUuid:  uuidRHD,
				Location:    fmt.Sprintf("%s%s%s", "http://", ipRHD, "/fakepath"),
				MaxAge:      2,
			}
			serviceNameRHD = "RackHD-service:api:2.0"
			serviceTagRHD = "RackHD"
		})

		It("should successfully register an RackHD-on-http service", func() {
			ssdp.AdvertiseServer(serverRHD)
			go ssdp.Start()

			// Start the test and wait for SSDP message to be seen
			r.AddSearchTerm(urnRHD, serviceTagRHD)
			go r.Run()

			Eventually(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceNameRHD] != nil
			}, timeout, pollingInterval).Should(BeTrue())

			Consistently(func() bool {
				service, _ := r.Store.Service(serviceNameRHD, "", nil)
				return service[0].ServiceName == serviceNameRHD
			}, timeout, pollingInterval).Should(BeTrue())

			ssdp.Stop()
			r.Stop()
			log.Println("Finished test: Register RHD successfully")

		})

		It("should successfully register an Inservive-Agent service", func() {
			ssdp.AdvertiseServer(serverIS)
			go ssdp.Start()

			// Start the test and wait for SSDP message to be seen
			r.AddSearchTerm(urnIS, serviceTagIS)
			go r.Run()

			Eventually(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceNameIS] != nil
			}, timeout, pollingInterval).Should(BeTrue())

			Consistently(func() bool {
				service, _ := r.Store.Service(serviceNameIS, "", nil)
				return service[0].ServiceName == serviceNameIS
			}, timeout, pollingInterval).Should(BeTrue())

			ssdp.Stop()
			r.Stop()
			log.Println("Finished test: Register IA successfully")

		})

		It("should handle badly formed URNs if they're whitelisted", func() {
			serviceNameIS = "Inservice-urn:MalformedData"
			urnIS = "urn:MalformedData"
			serverIS.ServiceType = urnIS

			ssdp.AdvertiseServer(serverIS)
			go ssdp.Start()

			// Start the test and wait for SSDP message to be seen
			r.AddSearchTerm(urnIS, serviceTagIS)
			go r.Run()

			Eventually(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceNameIS] != nil
			}, timeout, pollingInterval).Should(BeTrue())

			Consistently(func() bool {
				service, _ := r.Store.Service(serviceNameIS, "", nil)
				return service[0].ServiceName == serviceNameIS
			}, timeout, pollingInterval).Should(BeTrue())

			r.Stop()
			ssdp.Stop()
			log.Println("Finished test: Handle Malformed URN")

		})

		It("should ignore SSDP messages with unknown URNs", func() {

			ssdp.AdvertiseServer(serverIS)
			go ssdp.Start()

			// Start the test and wait for SSDP message to be seen
			urnIS = "urn:BAD:URN"
			r.AddSearchTerm(urnIS, serviceTagIS)
			go r.Run()

			Consistently(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceNameIS] == nil
			}, timeout, pollingInterval).Should(BeTrue())

			r.Stop()
			ssdp.Stop()
			log.Println("Finished test: Ignore unknown URN")

		})

		It("should fail to parse SSDP messages with a bad IP address", func() {

			badIP := "192.168.1165535"
			badScheme := "://"

			serverIS.Location = fmt.Sprintf("%s%s%s", "http://", badIP, "/fakepath")
			ssdp.AdvertiseServer(serverIS)

			noSchemeServer := serverIS
			noSchemeServer.Location = fmt.Sprintf("%s%s%s", badScheme, ipIS, "/fakepath")
			ssdp.AdvertiseServer(noSchemeServer)
			go ssdp.Start()

			r.AddSearchTerm(urnIS, serviceTagIS)

			// Start the test and wait for SSDP message to be seen
			go r.Run()

			Consistently(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceNameIS] == nil
			}, timeout, pollingInterval).Should(BeTrue())

			r.Stop()
			ssdp.Stop()
			log.Println("Finished test: Fail to parse bad IP")

		})

		It("should fail to parse SSDP messages with a bad port", func() {

			ipIS := "192.168.1.1:INVALIDPORT"
			serverIS.Location = fmt.Sprintf("%s%s%s", "http://", ipIS, "/fakepath")
			ssdp.AdvertiseServer(serverIS)
			go ssdp.Start()

			// Start the test and wait for SSDP message to be seen
			r.AddSearchTerm(urnIS, serviceTagIS)
			go r.Run()

			Consistently(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceNameIS] == nil
			}, timeout, pollingInterval).Should(BeTrue())

			r.Stop()
			ssdp.Stop()
			log.Println("Finished test: Fail to parse bad port")

		})

		It("should fail to register SSDP messages with no Node UUID", func() {

			serverIS.DeviceUuid = ""
			ssdp.AdvertiseServer(serverIS)
			go ssdp.Start()

			// Start the test and wait for SSDP message to be seen
			r.AddSearchTerm(urnIS, serviceTagIS)
			go r.Run()

			Consistently(func() bool {
				services, _ := r.Store.Services(nil)
				return services[serviceNameIS] == nil
			}, timeout, pollingInterval).Should(BeTrue())

			r.Stop()
			ssdp.Stop()
			log.Println("Finished test: Fail to register message without UUID")

		})
	})
})
