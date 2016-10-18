package watcher

import (
	"fmt"
	"github.com/RackHD/neighborhood-manager/libreg"
	regStore "github.com/RackHD/neighborhood-manager/libreg/registry"
	"log"
	"sync"
)

// Monitor is a struct to initialize a BackendStore
type Monitor struct {
	Store        regStore.Registry
	Datacenter   string
	ServiceName  string
	AddrMap      map[string]struct{}
	AddrMapMutex sync.Mutex
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
	n.AddrMapMutex = sync.Mutex{}
	go n.WatchService()
	return n, err
}

// GetAddresses calls GetService and passes our desired ServiceName.
// It then creates a map with the (ip:port)'s retrieved from the GetService call
func (m *Monitor) GetAddresses() (map[string]struct{}, error) {
	m.AddrMapMutex.Lock()
	defer m.AddrMapMutex.Unlock()
	return m.AddrMap, nil
}

// UpdateAddresses ...
func (m *Monitor) UpdateAddresses(service []*regStore.CatalogService) error {
	updateAddrMap := make(map[string]struct{})
	for _, entry := range service {
		addr := fmt.Sprintf("%s:%d", entry.ServiceAddress, entry.ServicePort)
		updateAddrMap[addr] = struct{}{}
	}
	m.AddrMapMutex.Lock()
	defer m.AddrMapMutex.Unlock()
	m.AddrMap = updateAddrMap
	return nil
}

// WatchService fetches catalog entries for the given serviceName
func (m *Monitor) WatchService() {
	opts := &regStore.QueryOptions{
		Datacenter: m.Datacenter,
	}
	// TODO: Handle the no service exist case
	//entries, err := m.Store.Service(serviceName, "", opts)
	stopCh := make(<-chan struct{})
	events, err := m.Store.ServiceWatch(m.ServiceName, "", opts, stopCh)
	if err != nil {
		return
	}

	for entries := range events {
		for _, e := range entries {
			fmt.Printf("+%v\n\n\n\n", e)
		}
		err := m.UpdateAddresses(entries)
		if err != nil {
			return
		}
	}
}
