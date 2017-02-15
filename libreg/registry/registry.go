package registry

import (
	"crypto/tls"
	"errors"
	"time"
)

// Backend represents a Service Registry Backend
type Backend string

const (
	// CONSUL backend
	CONSUL Backend = "consul"

	// MOCK backend
	MOCK Backend = "mock"
)

var (
	// ErrBackendNotSupported is thrown when the backend registry is not supported by libreg
	ErrBackendNotSupported = errors.New("Backend registry not supported yet, please choose one of")
	// ErrCallNotSupported is thrown when a method is not implemented/supported by the current registry
	ErrCallNotSupported = errors.New("The current call is not supported with this backend")
	// ErrNotReachable is thrown when the API cannot be reached for issuing common registry operations
	ErrNotReachable = errors.New("API not reachable")
	// ErrNodeNotFound is thrown when the node is not found int he registry
	ErrNodeNotFound = errors.New("Node not found in registry")
	// ErrSvcNotFound is thrown when the service is not found in the registry
	ErrSvcNotFound = errors.New("Service not found in registry")
)

// Config contains the options for a registry client
type Config struct {
	ClientTLS         *ClientTLSConfig
	TLS               *tls.Config
	ConnectionTimeout time.Duration
	PersistConnection bool
	Username          string
	Password          string
}

// ClientTLSConfig contains data for a Client TLS configuration in the form
// the etcd client wants it.  Eventually we'll adapt it for ZK and Consul.
type ClientTLSConfig struct {
	CertFile   string
	KeyFile    string
	CACertFile string
}

// AgentServiceRegistration is used to register a new service
type AgentServiceRegistration struct {
	ID                string
	Name              string
	Tags              []string
	Port              int
	Address           string
	EnableTagOverride bool
	Check             *AgentServiceCheck
	Checks            AgentServiceChecks
}

// AgentServiceCheck is used to define a node or service level check
type AgentServiceCheck struct {
	Script            string
	DockerContainerID string
	Shell             string // Only supported for Docker.
	Interval          string
	Timeout           string
	TTL               string
	HTTP              string
	TCP               string
	Status            string
	Notes             string
	TLSSkipVerify     bool

	// In Consul 0.7 and later, checks that are associated with a service
	// may also contain this optional DeregisterCriticalServiceAfter field,
	// which is a timeout in the same Go time format as Interval and TTL. If
	// a check is in the critical state for more than this configured value,
	// then its associated service (and all of its associated checks) will
	// automatically be deregistered.
	DeregisterCriticalServiceAfter string
}

// AgentServiceChecks is an array of AgentServiceCheck
type AgentServiceChecks []*AgentServiceCheck

// Registry represents the backend registry storage
// Each registry should support every call listed
// here or it cannot be utilized as a registry backend
type Registry interface {
	// Registers a new local service
	ServiceRegister(serv *AgentServiceRegistration) error
	// Register creates a new node, service or check
	Register(reg *CatalogRegistration, options *WriteOptions) error

	// Deregister removes a node, service or check
	Deregister(dereg *CatalogDeregistration, options *WriteOptions) error

	// Datacenters lists known datacenters
	Datacenters() ([]string, error)

	// Nodes lists all nodes in a given DC
	Nodes(options *QueryOptions) ([]*Node, error)

	// NodesWatch watches for changes to the nodes list
	NodesWatch(options *QueryOptions, stopCh <-chan struct{}) (<-chan []*Node, error)

	// Services lists all services in a given DC
	Services(options *QueryOptions) (map[string][]string, error)

	// ServicesWatch watches for changes to the list of services in a given DC
	ServicesWatch(options *QueryOptions, stopCh <-chan struct{}) (<-chan map[string][]string, error)

	// Service lists the nodes in a given service
	Service(service, tag string, options *QueryOptions) ([]*CatalogService, error)

	// ServiceWatch watches for changes to a service in a given DC
	ServiceWatch(service, tag string, options *QueryOptions, stopChan <-chan struct{}) (<-chan []*CatalogService, error)

	// Node lists the services provided by a given node
	Node(node string, options *QueryOptions) (*CatalogNode, error)

	// Leader returns the leader of the cluster
	Leader() (string, error)
}

// Node is ...
type Node struct {
	Node    string
	Address string
}

// CatalogService is ...
type CatalogService struct {
	Node                     string
	Address                  string
	ServiceID                string
	ServiceName              string
	ServiceAddress           string
	ServiceTags              []string
	ServicePort              int
	ServiceEnableTagOverride bool
}

// CatalogNode is ...
type CatalogNode struct {
	Node     *Node
	Services map[string]*AgentService
}

// CatalogRegistration is ...
type CatalogRegistration struct {
	Node       string
	Address    string
	Datacenter string
	Service    *AgentService
	Check      *AgentCheck
}

//CatalogDeregistration is ...
type CatalogDeregistration struct {
	Node       string
	Address    string
	Datacenter string
	ServiceID  string
	CheckID    string
}

// WriteOptions are used to parameterize a write
type WriteOptions struct {
	// Providing a datacenter overwrites the DC provided
	// by the Config
	Datacenter string

	// Token is used to provide a per-request ACL token
	// which overrides the agent's default token.
	Token string
}

// QueryOptions are used to parameterize a query
type QueryOptions struct {
	// Providing a datacenter overwrites the DC provided
	// by the Config
	Datacenter string

	// AllowStale allows any Consul server (non-leader) to service
	// a read. This allows for lower latency and higher throughput
	AllowStale bool

	// RequireConsistent forces the read to be fully consistent.
	// This is more expensive but prevents ever performing a stale
	// read.
	RequireConsistent bool

	// WaitIndex is used to enable a blocking query. Waits
	// until the timeout or the next index is reached
	WaitIndex uint64

	// WaitTime is used to bound the duration of a wait.
	// Defaults to that of the Config, but can be overridden.
	WaitTime time.Duration

	// Token is used to provide a per-request ACL token
	// which overrides the agent's default token.
	Token string

	// Near is used to provide a node name that will sort the results
	// in ascending order based on the estimated round trip time from
	// that node. Setting this to "_agent" will use the agent's node
	// for the sort.
	Near string
}

// AgentCheck represents a check known to the agent
type AgentCheck struct {
	Node        string
	CheckID     string
	Name        string
	Status      string
	Notes       string
	Output      string
	ServiceID   string
	ServiceName string
}

// AgentService represents a service known to the agent
type AgentService struct {
	ID                string
	Service           string
	Tags              []string
	Port              int
	Address           string
	EnableTagOverride bool
}
