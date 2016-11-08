package rhdman

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/RackHD/neighborhood-manager/rackhd/models"
	"github.com/RackHD/neighborhood-manager/rackhd/watcher"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/streadway/amqp"
)

// RackHDMan is a struct
type RackHDMan struct {
	WatcherChan chan *watcher.Message
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
func NewRackHDMan(msgChan chan *watcher.Message) RackHDMan {
	r := RackHDMan{
		WatcherChan: msgChan,
	}
	return r
}

// Start starts things
func (r *RackHDMan) Start() {
	go r.ProcessChan()
}

// ProcessChan is ...
func (r *RackHDMan) ProcessChan() error {
	for watcherMsg := range r.WatcherChan {
		switch watcherMsg.Action {
		case "create":
			if err := r.CreateNewRackHD(watcherMsg); err != nil {
				return err
			}
		case "update":
			if err := r.UpdateRackHD(watcherMsg); err != nil {
				return err
			}
		case "delete":
			if err := r.DeleteRackHD(watcherMsg); err != nil {
				return err
			}
		}
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
	uri, _, err := GetRackHDamqpURI(watcherMsg.Instance.Address)
	if err != nil {
		return err
	}
	rHD, err := models.NewRhd(watcherMsg.Instance.UUID, "http", uri.String())
	if err != nil {
		return err
	}
	if err = models.UpdateRhd(rHD); err != nil {
		return err
	}
	return nil
}

// CreateNewRackHD is going to create
func (r *RackHDMan) CreateNewRackHD(watcherMsg *watcher.Message) error {
	uri, _, err := GetRackHDamqpURI(watcherMsg.Instance.Address)
	if err != nil {
		return err
	}
	rHD, err := models.NewRhd(watcherMsg.Instance.UUID, "http", uri.String())
	if err != nil {
		return err
	}
	if err = models.CreateRhd(rHD); err != nil {
		return err
	}
	return nil
}

// InjectConfigParams is a stub to inject creds
func (r *RackHDMan) InjectConfigParams(conf map[string]interface{}) error {
	return nil
}

// GetRackHDamqpURI returns the RHD amqpURI
func GetRackHDamqpURI(address string) (amqp.URI, map[string]interface{}, error) {
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
		return amqp.URI{}, nil, fmt.Errorf("Failed to create /config request")
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return amqp.URI{}, nil, fmt.Errorf("Failed to get RHD configuration")
	}
	var v interface{}
	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return amqp.URI{}, nil, fmt.Errorf("Failed to decode RHD configuration")
	}
	conf := v.(map[string]interface{})
	rhdAmqpURI, ok := conf["amqp"]
	if !ok {
		return amqp.URI{}, nil, fmt.Errorf("Unable to get RHD amqp interface")
	}
	amqpURI, err := amqp.ParseURI(rhdAmqpURI.(string))
	if err != nil {
		return amqp.URI{}, nil, fmt.Errorf("Failed to parse AMQP URI")
	}
	return amqpURI, conf, nil
}
