package registry

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/RackHD/neighborhood-manager/libreg"
	regStore "github.com/RackHD/neighborhood-manager/libreg/registry"
	"github.com/king-jam/gossdp"
)

// Registry is a service registry object
type Registry struct {
	start       chan bool
	stop        chan bool
	wg          *sync.WaitGroup
	ssdpServer  *gossdp.Ssdp
	Store       regStore.Registry
	searchTerms map[string]string
	datacenter  string
	mut         sync.RWMutex
}

// AddSearchTerm adds a new tagged URN to the whitelist of SSDP URNs to act on
func (p *Registry) AddSearchTerm(urn, tag string) {
	p.mut.Lock()
	defer p.mut.Unlock()

	p.searchTerms[urn] = tag
}

// RemoveSearchTerm removes an entry from the whitelist of SSDP URNs to act on
func (p *Registry) RemoveSearchTerm(urn string) {
	p.mut.Lock()
	defer p.mut.Unlock()

	delete(p.searchTerms, urn)
}

// GetSearchTerms returns the whitelist of SSDP URNs to act on
func (p *Registry) GetSearchTerms() map[string]string {
	p.mut.RLock()
	defer p.mut.RUnlock()

	st := p.searchTerms
	return st
}

func extractIPPort(location string) (string, int, error) {
	addr, err := url.Parse(location)
	if err != nil {
		return "", 0, err
	}

	agentIP, portStr, err := net.SplitHostPort(addr.Host)
	if err != nil {
		return "", 0, err
	}

	agentPort, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, err
	}

	return agentIP, agentPort, nil
}

// NotifyBye handles a NotifyBye message from the SSDP listener
func (p *Registry) NotifyBye(message gossdp.ByeMessage) {
	// Should never hit this in normal execution.
	// Implemented to satisfy interface
	log.Printf("%+v\n", message)
}

// Response handles an M-Search Response message
func (p *Registry) Response(message gossdp.ResponseMessage) {
	// Should never hit this in normal execution.
	// Implemented to satisfy interface
	log.Printf("%+v\n", message)
}

// NotifyAlive handles a NotifyAlive message from the SSDP listener
func (p *Registry) NotifyAlive(message gossdp.AliveMessage) {
	// Check to see if URN is in whitelist, and get tag if yes
	st := p.GetSearchTerms()
	tag, goodUrn := st[message.Urn]

	if !goodUrn {
		return
	}

	//Expects URN for a service to be in form:
	// urn:dmtf-org:service:ServiceName:VersionNum
	//Extract "service:ServiceName:VersionNum"
	var serviceName string
	if strings.Count(message.Urn, ":") < 3 {
		serviceName = tag + "-" + message.Urn
	} else {
		parsedUrn := strings.SplitN(message.Urn, ":", 3)
		serviceName = tag + "-" + parsedUrn[2]
	}

	// Check to see if service is already registered for this node
	node, err := p.Store.Node(message.DeviceId, nil)
	if err != nil {
		log.Printf("Error checking for node: %s\n", err)
		return
	}
	if node != nil {
		for service := range node.Services {
			if service == serviceName {
				return

			}
		}

	}

	// Parse the IP broadcasted for IP and port
	agentIP, agentPort, err := extractIPPort(message.Location)
	if err != nil {
		log.Printf("Error parsing SSDP Message IP: %s\n", err)
		return
	}

	// Register the service with backend store
	err = p.register(agentIP, message.DeviceId, serviceName, agentPort)
	if err != nil {
		log.Printf("Error: Could not register service %s: %s\n", serviceName, err)
		return
	}

	if err = p.agentServiceRegister(serviceName, agentIP, "15s", agentPort); err != nil {
		log.Printf("Error: Could not register agent service check %s: %s\n", serviceName, err)
		return
	}

	log.Printf("New Service <%s> was registered for Node %s\n", serviceName, message.DeviceId)
}

// agentServiceRegister registers a local agent service and its check
func (p *Registry) agentServiceRegister(pluginName, agentIP, interval string, pluginPort int) error {
	err := p.Store.ServiceRegister(
		&regStore.AgentServiceRegistration{
			ID:                pluginName,
			Name:              pluginName,
			Port:              pluginPort,
			Address:           agentIP,
			EnableTagOverride: false,
			Check: &regStore.AgentServiceCheck{
				HTTP:                           "http://" + fmt.Sprintf("%s:%d/api/2.0/nodes", agentIP, pluginPort),
				Interval:                       interval,
				DeregisterCriticalServiceAfter: "5m",
			},
		},
	)
	return err
}

// register adds a service to the backend store
func (p *Registry) register(agentIP, deviceID, pluginName string, pluginPort int) error {

	// Register node with backend store
	err := p.Store.Register(
		&regStore.CatalogRegistration{
			Node:       deviceID,
			Address:    agentIP,
			Datacenter: p.datacenter,
			Service: &regStore.AgentService{
				ID:      pluginName,
				Service: pluginName,
				Port:    pluginPort,
				Address: agentIP,
			},
			Check: &regStore.AgentCheck{
				Node:        deviceID,
				CheckID:     "service:" + pluginName,
				Name:        "Service '" + pluginName + "' check",
				Status:      "passing",
				ServiceID:   pluginName,
				ServiceName: pluginName,
			},
		},
		nil,
	)
	return err
}

// NewRegistry creates a registry object
func NewRegistry(backend regStore.Backend, dc, backendAddr string) (*Registry, error) {

	r := &Registry{}

	// Create backend store (CONSUL, ETCD, etc)
	s, err := libreg.NewRegistry(backend, []string{backendAddr}, nil)
	if err != nil {
		log.Printf("Error creating backend store: %s\n", err)
		return nil, err
	}
	r.Store = s

	// Create client for capturing SSDP broadcasts
	client, err := gossdp.NewSsdp(r)
	if err != nil {
		log.Println("Failed to start client: ", err)
		return nil, err
	}
	r.ssdpServer = client

	r.start = make(chan bool)
	r.wg = &sync.WaitGroup{}
	r.searchTerms = make(map[string]string)
	r.datacenter = dc

	return r, nil
}

// Start blocks until Stop is called
func (p *Registry) Start() {

	for {
		select {
		case <-p.start:
		case <-p.stop:
			p.wg.Done()
			log.Println("Registry stopped listening")
			return
		}
	}
}

// Run handles Service Registry startup
func (p *Registry) Run() {
	p.stop = make(chan bool)

	p.wg.Add(1)
	go p.ssdpServer.Start()

	p.wg.Add(1)
	go p.Start()

	p.start <- true

	log.Println("Service Registry Starting.")
	p.wg.Wait()
}

// Stop closes a channel that should stop all capture
func (p *Registry) Stop() {
	log.Println("Service Registry Stopping.")

	p.ssdpServer.Stop()
	p.wg.Done()
	close(p.stop)

	p.wg.Wait()
}
