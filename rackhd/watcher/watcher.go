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
	AddrMap      map[string]Instance
	MsgChan      chan *Message
	AddrMapMutex sync.Mutex
}

// Instance is an struct with instance data
type Instance struct {
	UUID    string
	Address string
}

// Message is a message about an instance
type Message struct {
	Instance Instance
	Action   string
}

// NewMonitor initializes a new Monitor object
// creates backend store (CONSUL, ETCD, etc)
func NewMonitor(serviceName, datacenter, backendAddr string, backend regStore.Backend, msgChan chan *Message) (*Monitor, error) {
	n := &Monitor{}
	r, err := libreg.NewRegistry(backend, []string{backendAddr}, nil)
	if err != nil {
		log.Printf("Error creating backend store: %s\n", err)
		return nil, err
	}
	n.Store = r
	n.Datacenter = datacenter
	n.ServiceName = serviceName
	n.MsgChan = msgChan
	n.AddrMapMutex = sync.Mutex{}
	go n.WatchService()
	return n, err
}

// GetAddresses returns the address map
func (m *Monitor) GetAddresses() (map[string]Instance, error) {
	m.AddrMapMutex.Lock()
	defer m.AddrMapMutex.Unlock()
	return m.AddrMap, nil
}

// UpdateAddresses ...
func (m *Monitor) UpdateAddresses(service []*regStore.CatalogService) error {
	m.AddrMapMutex.Lock()
	defer m.AddrMapMutex.Unlock()
	updateAddrMap := make(map[string]Instance)

	for _, entry := range service {
		address := fmt.Sprintf("%s:%d", entry.ServiceAddress, entry.ServicePort)
		if v, ok := m.AddrMap[entry.Node]; !ok {
			updateAddrMap[entry.Node] = Instance{
				UUID:    entry.Node,
				Address: address,
			}
			m.MsgChan <- &Message{
				Instance: updateAddrMap[entry.Node],
				Action:   "create",
			}
			continue
		} else if v.Address == address {
			updateAddrMap[entry.Node] = Instance{
				UUID:    entry.Node,
				Address: address,
			}
			continue
		} else {
			updateAddrMap[entry.Node] = Instance{
				UUID:    entry.Node,
				Address: address,
			}
			m.MsgChan <- &Message{
				Instance: updateAddrMap[entry.Node],
				Action:   "update",
			}
			continue
		}
	}

	for k := range m.AddrMap {
		if _, ok := updateAddrMap[k]; !ok {
			m.MsgChan <- &Message{
				Instance: m.AddrMap[k],
				Action:   "delete",
			}
		}
	}

	m.AddrMap = updateAddrMap
	return nil
}

// WatchService fetches catalog entries for the given serviceName
func (m *Monitor) WatchService() {
	opts := &regStore.QueryOptions{
		Datacenter: m.Datacenter,
	}
	stopCh := make(<-chan struct{})

	events, err := m.Store.ServiceWatch(m.ServiceName, "", opts, stopCh)
	if err != nil {
		return
	}

	for entries := range events {
		err := m.UpdateAddresses(entries)
		if err != nil {
			return
		}
	}
}
