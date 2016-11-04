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
	monLock     *sync.Mutex
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
	// range over channels
	// switch statement over msg
	// do some
}

// ProcessChan is ...
func (r *RackHDMan) ProcessChan() error {
	for watcherMsg := range r.WatcherChan {
		uri, err := r.GetRackHDamqpURI(watcherMsg.Instance.Address)
		if err != nil {
			return err
		}
		switch watcherMsg.Action {
		case "create":
      models.NewRhd(watcherMsg.Instance.UUID, "http", amqpURI string)
			//			models.CreateRhd(rhd * models.RackHD)
		case "read":

		case "update":

		case "delete":

		}
	}
	return nil
}

// GetRackHDamqpURI returns the RHD amqpURI
func (r *RackHDMan) GetRackHDamqpURI(address string) (amqp.URI, error) {
	r.monLock.Lock()
	defer r.monLock.Unlock()
	// Get RackHD instance details
	client := cleanhttp.DefaultClient()
	uri := &url.URL{
		Scheme: "http",
		Host:   address,
		Path:   "/api/2.0/config",
	}
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return amqp.URI{}, fmt.Errorf("Failed to create /config request")
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return amqp.URI{}, fmt.Errorf("Failed to get RHD configuration")
	}
	var v interface{}
	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return amqp.URI{}, fmt.Errorf("Failed to decode RHD configuration")
	}
	conf := v.(map[string]interface{})
	rhdAmqpURI, ok := conf["amqp"]
	if !ok {
		return amqp.URI{}, fmt.Errorf("Unable to get RHD amqp interface")
	}
	amqpURI, err := amqp.ParseURI(rhdAmqpURI.(string))
	if err != nil {
		return amqp.URI{}, fmt.Errorf("Failed to parse AMQP URI")
	}
	return amqpURI, nil
}
