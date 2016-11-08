package federator

import (
	"fmt"
	"github.com/RackHD/neighborhood-manager/rackhd/manager"
	"github.com/michaelklishin/rabbit-hole"
	"github.com/streadway/amqp"
	"log"
	"net"
	"net/url"
	"sync"
)

// ExchangeConfig is ...
type ExchangeConfig struct {
	ExchangeName string
	ExchangeType string
	Durable      bool
}

// RackHD is ....
type RackHD struct {
	URI    amqp.URI
	StopCh chan struct{}
}

// AmqpFed is the AMQP Federation & Monitoring Service
type AmqpFed struct {
	URI      amqp.URI
	MgmtURI  url.URL
	conn     *amqp.Connection
	tag      string
	wg       *sync.WaitGroup
	monLock  *sync.Mutex
	monitors map[string]RackHD
}

// NewAmqpFed creates a new federation
func NewAmqpFed(amqpURI, mgmtPort, ctag string, exchangeList []ExchangeConfig) (*AmqpFed, error) {
	// take in list of exchanges by type and variable
	uri, err := amqp.ParseURI(amqpURI)
	mgmt, err := url.Parse(fmt.Sprintf("http://%s:%s", uri.Host, mgmtPort))
	if err != nil {
		return nil, fmt.Errorf("Could not parse AMQP URI")
	}
	a := &AmqpFed{
		URI:      uri,
		MgmtURI:  *mgmt,
		conn:     nil,
		tag:      ctag,
		wg:       &sync.WaitGroup{},
		monLock:  &sync.Mutex{},
		monitors: make(map[string]RackHD),
	}

	a.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}
	err = a.CreateFedPolicy()

	err = a.CreateDefaultExchanges(exchangeList)

	return a, err
}

// Shutdown starts the Server on address:port and handles the routes
func (a *AmqpFed) Shutdown() error {
	a.monLock.Lock()
	defer a.monLock.Unlock()
	for _, v := range a.monitors {
		close(v.StopCh)
	}
	a.wg.Wait()
	if err := a.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}
	return nil
}

//AddRackHD is ...
func (a *AmqpFed) AddRackHD(address string) error {
	amqpURI, conf, err := rhdman.GetRackHDamqpURI(address)
	if err != nil {
		return err
	}
	if amqpURI.Host == "0.0.0.0" {
		host, _, err := net.SplitHostPort(address)
		if err != nil {
			return err
		}
		amqpURI.Host = host
	} else if amqpURI.Host == "127.0.0.1" {
		return fmt.Errorf("AMQP not exposed externally")
	}
	hostname, ok := conf["HOSTNAME"]
	if !ok {
		hostname = "localhost"
		log.Println("Choosing sane default RHD name")
	}

	if err := a.CreateFedUpstream(amqpURI, hostname.(string)); err != nil {
		return fmt.Errorf("Unable to create upstream connection")
	}

	r, ok := a.monitors[address]
	if ok {
		//already exists, stop the goroutine
		close(r.StopCh)
	}

	newR := RackHD{
		URI:    amqpURI,
		StopCh: make(chan struct{}),
	}

	a.monitors[address] = newR

	return nil
}

// RemoveRackHD is ...
func (a *AmqpFed) RemoveRackHD(address string) error {
	a.monLock.Lock()
	defer a.monLock.Unlock()

	r, ok := a.monitors[address]
	if ok {
		//already exists, stop the goroutine
		close(r.StopCh)
	}

	delete(a.monitors, address)

	return nil
}

// CreateDefaultExchanges is ...
func (a *AmqpFed) CreateDefaultExchanges(exchangeList []ExchangeConfig) error {
	//TODO: CHANGE THIS TO HTTP APIS
	channel, err := a.conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}
	for _, conf := range exchangeList {
		if err = channel.ExchangeDeclare(
			conf.ExchangeName,
			conf.ExchangeType,
			conf.Durable,
			false,
			false,
			false,
			nil,
		); err != nil {
			return fmt.Errorf("Exchange Declare: %s", err)
		}
	}
	if err := channel.Cancel(a.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}
	return nil
}

// CreateFedUpstream is ...
func (a *AmqpFed) CreateFedUpstream(addr amqp.URI, name string) error {
	// PUT /api/parameters/federation-upstream/%2f/my-upstream
	rmqc, err := rabbithole.NewClient(a.MgmtURI.String(), "guest", "guest")
	if err != nil {
		return fmt.Errorf("Failed to create RMQ HTTP Client")
	}
	fedDefinition := rabbithole.FederationDefinition{
		Uri:     addr.String(),
		Expires: 36000000,
		MaxHops: 1,
		AckMode: "on-confirm",
	}
	_, err = rmqc.PutFederationUpstream("/", name, fedDefinition)
	if err != nil {
		return fmt.Errorf("Failed to create upstream")
	}
	return nil
}

// CreateFedPolicy is ...
func (a *AmqpFed) CreateFedPolicy() error {
	// PUT /api/policies/%2f/federate-rabbitmq
	// PUT /api/policies/%2f/federate-rackhd
	rmqc, err := rabbithole.NewClient(a.MgmtURI.String(), "guest", "guest")
	if err != nil {
		return fmt.Errorf("Failed to create RMQ HTTP Client")
	}
	// default optional policy set for the federation policies
	defaultPolicySet := map[string]interface{}{
		"federation-upstream-set": "all",
	}

	// policy to federate the default amq. exchanges from all hosts
	rabbit := rabbithole.Policy{
		Pattern:    "^amq\\.",
		Definition: defaultPolicySet,
		ApplyTo:    "exchanges",
	}
	_, err = rmqc.PutPolicy(a.URI.Vhost, "federate-rabbitmq", rabbit)
	if err != nil {
		return fmt.Errorf("Failed to create federate-rabbitmq policy")
	}

	// policy to federate the default on. exchange from all RackHD hosts
	rackhd := rabbithole.Policy{
		Pattern:    "^on\\.",
		Definition: defaultPolicySet,
		ApplyTo:    "exchanges",
	}

	_, err = rmqc.PutPolicy(a.URI.Vhost, "federate-rackhd", rackhd)
	if err != nil {
		return fmt.Errorf("Failed to create federate-rackhd policy")
	}

	return nil
}

// NewRackhdMonitor is ...
func (a *AmqpFed) NewRackhdMonitor() error {
	return nil
}
