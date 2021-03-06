package consul

import (
	"crypto/tls"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/RackHD/neighborhood-manager/libreg"
	"github.com/RackHD/neighborhood-manager/libreg/registry"
	"github.com/hashicorp/consul/api"
)

const (
	// DefaultWatchWaitTime is how long we will block when watching
	// for services or nodes to change/update
	DefaultWatchWaitTime = 15 * time.Second
)

var (
	// ErrMultipleEndpointsUnsupported is thrown when there are
	// multiple endpoints specified for Consul
	ErrMultipleEndpointsUnsupported = errors.New("consul does not support multiple endpoints")

	// ErrSessionRenew is thrown when the session can't be
	// renewed because the Consul version does not support sessions
	ErrSessionRenew = errors.New("cannot set or renew session for ttl, unable to operate on sessions")
)

// Consul is the receiver type for the
// Store interface
type Consul struct {
	sync.Mutex
	config *api.Config
	client *api.Client
}

// Register registers consul to libreg
func Register() {
	libreg.AddRegistry(registry.CONSUL, New)
}

// New creates a new Consul client given a list
// of endpoints and optional tls config
func New(endpoints []string, options *registry.Config) (registry.Registry, error) {
	if len(endpoints) > 1 {
		return nil, ErrMultipleEndpointsUnsupported
	}

	s := &Consul{}

	// Create Consul client
	config := api.DefaultConfig()
	s.config = config
	config.HttpClient = http.DefaultClient
	config.Address = endpoints[0]
	config.Scheme = "http"

	// Set options
	if options != nil {
		if options.TLS != nil {
			s.setTLS(options.TLS)
		}
		if options.ConnectionTimeout != 0 {
			s.setTimeout(options.ConnectionTimeout)
		}
	}

	// Creates a new client
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	s.client = client

	return s, nil
}

// SetTLS sets Consul TLS options
func (s *Consul) setTLS(tls *tls.Config) {
	s.config.HttpClient.Transport = &http.Transport{
		TLSClientConfig: tls,
	}
	s.config.Scheme = "https"
}

// SetTimeout sets the timeout for connecting to Consul
func (s *Consul) setTimeout(time time.Duration) {
	s.config.WaitTime = time
}

// ServiceRegister add a local agent service and its check
func (s *Consul) ServiceRegister(serv *registry.AgentServiceRegistration) error {

	err := s.client.Agent().ServiceRegister(
		&api.AgentServiceRegistration{
			ID:                serv.ID,
			Name:              serv.Name,
			Tags:              serv.Tags,
			Port:              serv.Port,
			Address:           serv.Address,
			EnableTagOverride: serv.EnableTagOverride,
			Check: &api.AgentServiceCheck{
				HTTP:                           serv.Check.HTTP,
				Interval:                       serv.Check.Interval,
				DeregisterCriticalServiceAfter: serv.Check.DeregisterCriticalServiceAfter,
			},
		},
	)
	return err
}

// Register adds an entry into the consul backend
func (s *Consul) Register(reg *registry.CatalogRegistration, options *registry.WriteOptions) error {
	catalog := s.client.Catalog()
	writeOps := s.getWriteOptions(options)
	_, err := catalog.Register(
		&api.CatalogRegistration{
			Node:       reg.Node,
			Address:    reg.Address,
			Datacenter: reg.Datacenter,
			Service: &api.AgentService{
				ID:                reg.Service.ID,
				Service:           reg.Service.Service,
				Tags:              reg.Service.Tags,
				Port:              reg.Service.Port,
				Address:           reg.Service.Address,
				EnableTagOverride: reg.Service.EnableTagOverride,
			},
			Check: &api.AgentCheck{
				Node:        reg.Check.Node,
				CheckID:     reg.Check.CheckID,
				Name:        reg.Check.Name,
				Status:      reg.Check.Status,
				Notes:       reg.Check.Notes,
				Output:      reg.Check.Output,
				ServiceID:   reg.Check.ServiceID,
				ServiceName: reg.Check.ServiceName,
			},
		},
		writeOps,
	)
	return err
}

// Deregister removes a node, service or check
func (s *Consul) Deregister(dereg *registry.CatalogDeregistration, options *registry.WriteOptions) error {
	catalog := s.client.Catalog()
	writeOps := s.getWriteOptions(options)
	_, err := catalog.Deregister(
		&api.CatalogDeregistration{
			Node:       dereg.Node,
			Address:    dereg.Address,
			Datacenter: dereg.Datacenter,
			ServiceID:  dereg.ServiceID,
			CheckID:    dereg.CheckID,
		},
		writeOps,
	)
	return err
}

// Datacenters lists known datacenters
func (s *Consul) Datacenters() ([]string, error) {
	catalog := s.client.Catalog()
	dc, err := catalog.Datacenters()
	return dc, err
}

// Nodes lists all nodes in a given DC
func (s *Consul) Nodes(options *registry.QueryOptions) ([]*registry.Node, error) {
	catalog := s.client.Catalog()
	queryOps := s.getQueryOptions(options)
	nodes, _, err := catalog.Nodes(queryOps)
	var retNodes []*registry.Node
	for _, v := range nodes {
		retNodes = append(retNodes, &registry.Node{
			Node:    v.Node,
			Address: v.Address,
		})
	}
	return retNodes, err
}

// NodesWatch watches for changes to the nodes in a given DC
func (s *Consul) NodesWatch(options *registry.QueryOptions, stopChan <-chan struct{}) (<-chan []*registry.Node, error) {
	catalog := s.client.Catalog()
	queryOps := s.getQueryOptions(options)
	watchCh := make(chan []*registry.Node)

	go func() {
		defer close(watchCh)

		// Override the wait time option to create the watch
		queryOps.WaitTime = DefaultWatchWaitTime

		for {
			select {
			case <-stopChan:
				return
			default:
			}

			nodes, meta, err := catalog.Nodes(queryOps)
			if err != nil {
				return
			}

			// If LastIndex didn't change then it means `Get` returned
			// because of the WaitTime and the key didn't change.
			if queryOps.WaitIndex == meta.LastIndex {
				continue
			}

			queryOps.WaitIndex = meta.LastIndex

			var retNodes []*registry.Node
			for _, v := range nodes {
				retNodes = append(retNodes, &registry.Node{
					Node:    v.Node,
					Address: v.Address,
				})
			}

			watchCh <- retNodes
		}
	}()

	return watchCh, nil
}

// Services lists all services in a given DC
func (s *Consul) Services(options *registry.QueryOptions) (map[string][]string, error) {
	catalog := s.client.Catalog()
	queryOps := s.getQueryOptions(options)
	services, _, err := catalog.Services(queryOps)
	return services, err
}

// ServicesWatch watches for changes to the list of services in a given DC
func (s *Consul) ServicesWatch(options *registry.QueryOptions, stopChan <-chan struct{}) (<-chan map[string][]string, error) {
	catalog := s.client.Catalog()
	queryOps := s.getQueryOptions(options)
	watchCh := make(chan map[string][]string)

	go func() {
		defer close(watchCh)

		// Override the wait time option to create the watch
		queryOps.WaitTime = DefaultWatchWaitTime

		for {
			select {
			case <-stopChan:
				return
			default:
			}

			services, meta, err := catalog.Services(queryOps)
			if err != nil {
				return
			}

			// If LastIndex didn't change then it means `Get` returned
			// because of the WaitTime and the key didn't change.
			if queryOps.WaitIndex == meta.LastIndex {
				continue
			}

			queryOps.WaitIndex = meta.LastIndex

			if services != nil {
				watchCh <- services
			}
		}
	}()

	return watchCh, nil
}

// Service lists the nodes in a given service
func (s *Consul) Service(service, tag string, options *registry.QueryOptions) ([]*registry.CatalogService, error) {
	catalog := s.client.Catalog()
	queryOps := s.getQueryOptions(options)
	services, _, err := catalog.Service(
		service,
		tag,
		queryOps)
	var retServices []*registry.CatalogService
	for _, v := range services {
		retServices = append(retServices, &registry.CatalogService{
			Node:                     v.Node,
			Address:                  v.Address,
			ServiceID:                v.ServiceID,
			ServiceName:              v.ServiceName,
			ServiceAddress:           v.ServiceAddress,
			ServiceTags:              v.ServiceTags,
			ServicePort:              v.ServicePort,
			ServiceEnableTagOverride: v.ServiceEnableTagOverride,
		})
	}
	return retServices, err
}

// ServiceWatch watches for changes to the list of nodes under a given service
func (s *Consul) ServiceWatch(service, tag string, options *registry.QueryOptions, stopChan <-chan struct{}) (<-chan []*registry.CatalogService, error) {
	catalog := s.client.Catalog()
	queryOps := s.getQueryOptions(options)
	watchCh := make(chan []*registry.CatalogService)

	go func() {
		defer close(watchCh)

		// Override the wait time option to create the watch
		queryOps.WaitTime = DefaultWatchWaitTime

		for {
			select {
			case <-stopChan:
				return
			default:
			}

			services, meta, err := catalog.Service(
				service,
				tag,
				queryOps)
			if err != nil {
				return
			}

			// If LastIndex didn't change then it means `Get` returned
			// because of the WaitTime and the key didn't change.
			if queryOps.WaitIndex == meta.LastIndex {
				continue
			}

			queryOps.WaitIndex = meta.LastIndex

			var retServices []*registry.CatalogService
			for _, v := range services {
				retServices = append(retServices, &registry.CatalogService{
					Node:                     v.Node,
					Address:                  v.Address,
					ServiceID:                v.ServiceID,
					ServiceName:              v.ServiceName,
					ServiceAddress:           v.ServiceAddress,
					ServiceTags:              v.ServiceTags,
					ServicePort:              v.ServicePort,
					ServiceEnableTagOverride: v.ServiceEnableTagOverride,
				})
			}

			watchCh <- retServices
		}
	}()

	return watchCh, nil
}

// Node lists the services provided by a given node
func (s *Consul) Node(node string, options *registry.QueryOptions) (*registry.CatalogNode, error) {
	catalog := s.client.Catalog()
	queryOps := s.getQueryOptions(options)
	n, _, err := catalog.Node(
		node,
		queryOps,
	)
	if err != nil {
		return nil, err
	}
	if n == nil {
		return nil, nil
	}

	var retNode = &registry.Node{
		Node:    n.Node.Node,
		Address: n.Node.Address,
	}
	var retService = make(map[string]*registry.AgentService)
	for k, v := range n.Services {
		retService[k] = &registry.AgentService{
			ID:                v.ID,
			Service:           v.Service,
			Tags:              v.Tags,
			Port:              v.Port,
			Address:           v.Address,
			EnableTagOverride: v.EnableTagOverride,
		}
	}
	return &registry.CatalogNode{
		Node:     retNode,
		Services: retService,
	}, err
}

func (s *Consul) getWriteOptions(options *registry.WriteOptions) *api.WriteOptions {
	ops := &api.WriteOptions{}
	if options != nil {
		if options.Datacenter != "" {
			ops.Datacenter = options.Datacenter
		}
		if options.Token != "" {
			ops.Token = options.Token
		}
	}
	return ops

}

func (s *Consul) getQueryOptions(options *registry.QueryOptions) *api.QueryOptions {
	if options == nil {
		return nil
	}
	ops := &api.QueryOptions{}
	if options != nil {
		ops.Datacenter = options.Datacenter
		ops.AllowStale = options.AllowStale
		ops.RequireConsistent = options.RequireConsistent
		ops.WaitIndex = options.WaitIndex
		ops.WaitTime = options.WaitTime
		ops.Token = options.Token
		ops.Near = options.Near
	}
	return ops
}

// Leader returns the cluster leader
func (s *Consul) Leader() (string, error) {
	status := s.client.Status()
	return status.Leader()
}
