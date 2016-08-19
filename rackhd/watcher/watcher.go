package watcher

import (
	"fmt"
	"log"

	"github.com/RackHD/neighborhood-manager/libreg"
	regStore "github.com/RackHD/neighborhood-manager/libreg/registry"
)

// Monitor is a struct to initialize a BackendStore
type Monitor struct {
	Store       regStore.Registry
	Datacenter  string
	ServiceName string
}

// NewMonitor initializes a new Monitor object
// creates backend store (CONSUL, ETCD, etc)
func NewMonitor(serviceName, datacenter, backendAddr string, backend regStore.Backend) (*Monitor, error) {
	n := &Monitor{}
	r, err := libreg.NewRegistry(backend, []string{backendAddr}, nil)
	if err != nil {
		log.Printf("Error creating backend store: %s\n", err)
		return nil, err
	}
	n.Store = r
	n.Datacenter = datacenter
	n.ServiceName = serviceName
	return n, err
}

// GetAddresses calls GetService and passes our desired ServiceName.
// It then creates a map with the (ip:port)'s retrieved from the GetService call
func (m *Monitor) GetAddresses() (map[string]struct{}, error) {
	service, err := m.GetService(m.ServiceName)
	if err != nil {
		log.Printf("Error fetching %s catalog entries ==> %s\n", m.ServiceName, err)
		return nil, err
	}
	addresses := make(map[string]struct{})
	for _, entry := range service {
		addr := fmt.Sprintf("%s:%d", entry.ServiceAddress, entry.ServicePort)
		addresses[addr] = struct{}{}
	}
	return addresses, err
}

// GetService fetches catalog entries for the given serviceName
func (m *Monitor) GetService(serviceName string) ([]*regStore.CatalogService, error) {
	entries, err := m.Store.Service(serviceName, "", nil)
	if err != nil {
		log.Printf("Error fetching %s ==> %s\n", serviceName, err)
		return nil, err
	}
	return entries, err
}
