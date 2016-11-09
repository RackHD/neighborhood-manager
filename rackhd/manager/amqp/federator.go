package federator

import (
	"fmt"
	"github.com/michaelklishin/rabbit-hole"
	"github.com/streadway/amqp"
	"net/url"
)

var exchangeList = []ExchangeConfig{
	ExchangeConfig{
		ExchangeName: "on.events",
		ExchangeType: "topic",
		Durable:      true,
	},
	ExchangeConfig{
		ExchangeName: "on.heartbeat",
		ExchangeType: "topic",
		Durable:      true,
	},
}

// ExchangeConfig is ...
type ExchangeConfig struct {
	ExchangeName string
	ExchangeType string
	Durable      bool
}

// AmqpFed is the AMQP Federation & Monitoring Service
type AmqpFed struct {
	URI     amqp.URI
	MgmtURI url.URL
	conn    *amqp.Connection
	tag     string
}

// NewAmqpFed creates a new federation
func NewAmqpFed(amqpURI, mgmtPort, ctag string) (*AmqpFed, error) {
	// take in list of exchanges by type and variable
	uri, err := amqp.ParseURI(amqpURI)
	mgmt, err := url.Parse(fmt.Sprintf("http://%s:%s", uri.Host, mgmtPort))
	if err != nil {
		return nil, fmt.Errorf("Could not parse AMQP URI")
	}
	a := &AmqpFed{
		URI:     uri,
		MgmtURI: *mgmt,
		conn:    nil,
		tag:     ctag,
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
	if err := a.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}
	return nil
}

//AddRackHD takes in an amqpURI address
func (a *AmqpFed) AddRackHD(amqpURI amqp.URI, uuid string) error {

	if err := a.CreateFedUpstream(amqpURI, uuid); err != nil {
		return fmt.Errorf("Unable to create upstream connection")
	}

	return nil
}

// RemoveRackHD is ...
func (a *AmqpFed) RemoveRackHD(address string) error {
	// a.monLock.Lock()
	// defer a.monLock.Unlock()
	//
	// r, ok := a.monitors[address]
	// if ok {
	// 	//already exists, stop the goroutine
	// 	close(r.StopCh)
	// }
	//
	// delete(a.monitors, address)

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
