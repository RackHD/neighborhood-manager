package rhdman

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync"

	"github.com/RackHD/neighborhood-manager/rackhd/manager/amqp"
	"github.com/RackHD/neighborhood-manager/rackhd/models"
	"github.com/RackHD/neighborhood-manager/rackhd/watcher"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/streadway/amqp"
)

// RackHDMan is a struct
type RackHDMan struct {
	WatcherChan chan *watcher.Message
	Federator   *federator.AmqpFed
	RHDS        map[string]*RackHD
}

var monLock *sync.Mutex

// HTTPEndpoints is a struct to inject creds into the URI
type HTTPEndpoints struct {
	Address string
	Port    int
	HTTPS   bool
	Proxies bool
	Auth    bool
	Routers string
}

// NewRackHDMan instantiates a new RackHDMan obj
func NewRackHDMan(msgChan chan *watcher.Message) (RackHDMan, error) {
	fed, err := federator.NewAmqpFed("amqp://guest:guest@localhost:5672", "8080", "nm")
	if err != nil {
		return RackHDMan{}, err
	}

	rhd := make(map[string]*RackHD)
	r := RackHDMan{
		WatcherChan: msgChan,
		Federator:   fed,
		RHDS:        rhd,
	}
	return r, nil
}

// Start starts things
func (r *RackHDMan) Start() {
	go r.ProcessChan()
}

// ProcessChan is ...
func (r *RackHDMan) ProcessChan() error {
	for watcherMsg := range r.WatcherChan {
		fmt.Printf("%+v", *watcherMsg)
		switch watcherMsg.Action {
		case "create":
			fmt.Println("HERE 1")
			if err := r.CreateNewRackHD(watcherMsg); err != nil {
				return err
			}
		case "update":
			fmt.Println("HERE 2")
			if err := r.UpdateRackHD(watcherMsg); err != nil {
				return err
			}
		case "delete":
			fmt.Println("HERE 3")
			if err := r.DeleteRackHD(watcherMsg); err != nil {
				return err
			}
		}
		fmt.Println("HERE 4")
	}
	return nil
}

// DeleteRackHD is going to delete
func (r *RackHDMan) DeleteRackHD(watcherMsg *watcher.Message) error {
	err := models.DeleteRhdByID(watcherMsg.Instance.UUID)
	return err
}

// UpdateRackHD is going to update
func (r *RackHDMan) UpdateRackHD(watcherMsg *watcher.Message) error {
	// conf, err := GetRackHDConfig(watcherMsg.Instance.Address)
	// if err != nil {
	// 	return err
	// }
	// uri, err := r.GetRackHDamqpURI(conf)
	// if err != nil {
	// 	return err
	// }
	// rHD, err := models.NewRhd(watcherMsg.Instance.UUID, "http", uri.String())
	// if err != nil {
	// 	return err
	// }
	// if err = models.UpdateRhd(rHD); err != nil {
	// 	return err
	// }
	return nil
}

// CreateNewRackHD is going to create
func (r *RackHDMan) CreateNewRackHD(watcherMsg *watcher.Message) error {
	conf, err := GetRackHDConfig(watcherMsg.Instance.Address)
	if err != nil {
		return err
	}
	uri, err := r.GetRackHDamqpURI(conf, watcherMsg.Instance.Address)
	if err != nil {
		return err
	}
	if err = r.Federator.AddRackHD(uri, watcherMsg.Instance.UUID); err != nil {
		return err
	}
	trackingObj := NewRackHDTracker(watcherMsg.Instance.UUID, uri)
	r.RHDS[watcherMsg.Instance.UUID] = trackingObj

	rHD, err := models.NewRhd(watcherMsg.Instance.UUID, "http", uri.String())
	if err != nil {
		return err
	}

	return models.CreateRhd(rHD)
}

// InjectConfigParams is a stub to inject creds
func (r *RackHDMan) InjectConfigParams(conf map[string]interface{}) error {
	return nil
}

// GetRackHDConfig returns the RHD amqpURI
func GetRackHDConfig(address string) (map[string]interface{}, error) {
	monLock.Lock()
	defer monLock.Unlock()
	// Get RackHD instance details
	client := cleanhttp.DefaultClient()
	uri := &url.URL{
		Scheme: "http",
		Host:   address,
		Path:   "/api/2.0/config",
	}
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create /config request")
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to get RHD configuration")
	}
	var v interface{}
	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, fmt.Errorf("Failed to decode RHD configuration")
	}
	conf := v.(map[string]interface{})
	return conf, nil
}

// GetRackHDamqpURI gets the amqp URI
func (r *RackHDMan) GetRackHDamqpURI(conf map[string]interface{}, address string) (amqp.URI, error) {
	rhdAmqpURI, ok := conf["amqp"]
	if !ok {
		return amqp.URI{}, fmt.Errorf("Unable to get RHD amqp interface")
	}
	amqpURI, err := amqp.ParseURI(rhdAmqpURI.(string))
	if err != nil {
		return amqp.URI{}, fmt.Errorf("Failed to parse AMQP URI")
	}
	if amqpURI.Host == "0.0.0.0" {
		host, _, err := net.SplitHostPort(address)
		if err != nil {
			return amqp.URI{}, err
		}
		amqpURI.Host = host
	} else if amqpURI.Host == "127.0.0.1" {
		// TODO: figure out if we are on the same host as RHD RabbitMQ instance
		return amqp.URI{}, fmt.Errorf("AMQP not exposed externally")
	}
	return amqpURI, nil
}

// RackHD stores the state data of managed instances
type RackHD struct {
	stop chan struct{}
	wg   sync.WaitGroup
}

// NewRackHDTracker creates go routines for days
func NewRackHDTracker(uuid string, uri amqp.URI) *RackHD {

	tObj := &RackHD{
		stop: make(chan struct{}),
		wg:   sync.WaitGroup{},
	}
	tObj.wg.Add(1)
	go func(t *RackHD) {
		defer t.wg.Done()
		MQ := federator.AMQPClient{}
		err := MQ.Initialize(uri.String())
		if err != nil {
			return
		}
		defer MQ.Close()
		events, err := MQ.AmqpListen(
			"on.events",
			"topic",
			uuid+"-events",
			"event.node",
			uuid+"-nm")
		if err != nil {
			return
		}
		for {
			select {
			case <-t.stop:
				return
			default:
				for event := range events {
					t.processEvents(&event)
				}
			}
		}
	}(tObj)

	return tObj
}

func (rhd *RackHD) processEvents(m *amqp.Delivery) error {
	switch m.Exchange {

	case "on.events":
		return rhd.processNodeEvent(m)

	default:
		return fmt.Errorf("Not my message, not my problem")
	}
}

func (rhd *RackHD) processNodeEvent(m *amqp.Delivery) error {
	fmt.Printf("%+v", m)
	m.Ack(true)
	return nil
}
