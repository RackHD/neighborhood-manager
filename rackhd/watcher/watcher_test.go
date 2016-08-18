package watcher_test

import (
	"fmt"

	regStore "github.com/RackHD/NeighborhoodManager/libreg/registry"
	"github.com/RackHD/NeighborhoodManager/libreg/registry/mock"
	"github.com/RackHD/NeighborhoodManager/rackhd/watcher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Watcher", func() {
	var (
		serviceName    string
		badServiceName string
		datacenter     string
		backendAddr    string
		badBackend     regStore.Backend
		endpointPort   int
		endpointAddr   string
	)
	BeforeEach(func() {
		serviceName = "RackHD-service:api:2.0"
		badServiceName = "service_error_injection"
		datacenter = "dc-test"
		backendAddr = "127.0.0.1:8500"
		endpointPort = 9090
		endpointAddr = "10.240.16.69"

	})

	Describe("NewMonitor", func() {
		It("should return a Monitor struct using mock backend", func() {
			mock.Register()
			m, err := watcher.NewMonitor(serviceName, datacenter, backendAddr, regStore.MOCK)

			Expect(m).To(BeAssignableToTypeOf(&watcher.Monitor{}))
			Expect(err).ToNot(HaveOccurred())
			Expect(m.Datacenter).To(Equal(datacenter))
			Expect(m.ServiceName).To(Equal(serviceName))
		})

		It("should reject invalid backend", func() {
			mock.Register()
			_, err := watcher.NewMonitor(serviceName, datacenter, backendAddr, badBackend)
			Expect(err).To(HaveOccurred())
		})

	})

	Describe("MonitorFunctions", func() {
		var (
			m *watcher.Monitor
		)
		BeforeEach(func() {
			mock.Register()
			m, _ = watcher.NewMonitor(serviceName, datacenter, backendAddr, regStore.MOCK)
			_ = m.Store.Register(&regStore.CatalogRegistration{
				Node:       "42192294-095f-a13c-8dad-52c27c87ec66",
				Address:    endpointAddr,
				Datacenter: datacenter,
				Service: &regStore.AgentService{
					ID:      serviceName,
					Service: serviceName,
					Port:    endpointPort,
					Address: endpointAddr,
				},
				Check: &regStore.AgentCheck{
					Node:        "42192294-095f-a13c-8dad-52c27c87ec66",
					CheckID:     "service:" + serviceName,
					Name:        "Service '" + serviceName + "' check",
					Status:      "passing",
					ServiceID:   serviceName,
					ServiceName: serviceName,
				},
			}, nil)

		})

		It("should return a map[string]struct{} of addresses from the backend store", func() {
			addresses, err := m.GetAddresses()
			Expect(addresses).To(BeAssignableToTypeOf(map[string]struct{}{}))
			Expect(err).ToNot(HaveOccurred())

		})

		It("should check that addresses[] has a valid address stored in the map", func() {
			addresses, err := m.GetAddresses()
			_, ok := addresses[fmt.Sprintf("%s:%d", endpointAddr, endpointPort)]
			Expect(ok).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})

		It("should return a backend array for the given serviceName", func() {
			s, err := m.GetService(serviceName)

			Expect(s).To(BeAssignableToTypeOf([]*regStore.CatalogService{}))
			Expect(err).ToNot(HaveOccurred())
		})

		It("should return an error for a badServiceName", func() {
			_, err := m.GetService(badServiceName)

			Expect(err).To(HaveOccurred())
		})

		It("should return an error for a badServiceName", func() {
			m.ServiceName = badServiceName
			_, err := m.GetAddresses()

			Expect(err).To(HaveOccurred())
		})
	})

})
