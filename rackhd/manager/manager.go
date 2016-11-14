package manager

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/RackHD/neighborhood-manager/rackhd/manager/amqp"
	"github.com/RackHD/neighborhood-manager/rackhd/models"
	"github.com/RackHD/neighborhood-manager/rackhd/watcher"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/streadway/amqp"
)

// Manager is a struct
type Manager struct {
	WatcherChan chan *watcher.Message
	Federator   *federator.AmqpFed
	RHDS        map[string]*NodeTracker
}

// NewManager instantiates a new Manager obj
func NewManager(msgChan chan *watcher.Message) (Manager, error) {
	fed, err := federator.NewAmqpFed("amqp://guest:guest@localhost:5672", "8080", "nm")
	if err != nil {
		return Manager{}, err
	}
	rhd := make(map[string]*NodeTracker)
	m := Manager{
		WatcherChan: msgChan,
		Federator:   fed,
		RHDS:        rhd,
	}
	return m, nil
}

// RunManager handles incoming messages
func (m *Manager) RunManager() {
	for watcherMsg := range m.WatcherChan {
		switch watcherMsg.Action {
		case "create":
			go m.CreateNewRackHD(watcherMsg)
		case "update":
			go m.UpdateRackHD(watcherMsg)
		case "delete":
			go m.DeleteRackHD(watcherMsg)
		default:
		}
	}
}

// DeleteRackHD is going to delete
func (m *Manager) DeleteRackHD(watcherMsg *watcher.Message) error {
	err := models.DeleteRhdByID(watcherMsg.Instance.UUID)
	return err
}

// UpdateRackHD is going to update
func (m *Manager) UpdateRackHD(watcherMsg *watcher.Message) error {
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
func (m *Manager) CreateNewRackHD(watcherMsg *watcher.Message) error {
	conf, err := getRackHDConfig(watcherMsg.Instance.Address)
	if err != nil {
		return err
	}
	uri, err := getRackHDamqpURI(conf, watcherMsg.Instance.Address)
	if err != nil {
		return err
	}
	if err = m.Federator.AddRackHD(uri, watcherMsg.Instance.UUID); err != nil {
		return err
	}
	nodeTracker := NewNodeTracker(watcherMsg.Instance.UUID, uri)
	if err = nodeTracker.InitiateNodeCache(watcherMsg.Instance.Address); err != nil {
		fmt.Println("Failed to init the node cache")
	}
	m.RHDS[watcherMsg.Instance.UUID] = nodeTracker
	rHD, err := models.NewRhd(watcherMsg.Instance.UUID, watcherMsg.Instance.Address, uri.String())
	if err != nil {
		return err
	}

	return models.CreateRhd(rHD)
}

// GetRackHDConfig returns the RHD amqpURI
func getRackHDConfig(address string) (map[string]interface{}, error) {
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
func getRackHDamqpURI(conf map[string]interface{}, address string) (amqp.URI, error) {
	rhdAmqpURIinterface, ok := conf["amqp"]
	if !ok {
		return amqp.URI{}, fmt.Errorf("Unable to get RHD amqp interface")
	}
	rhdAmqpURI := rhdAmqpURIinterface.(string)
	amqpURI, err := amqp.ParseURI(rhdAmqpURI)
	if err != nil {
		return amqp.URI{}, fmt.Errorf("Failed to parse AMQP URI")
	}
	if amqpURI.Host == "0.0.0.0" || amqpURI.Host == "localhost" || amqpURI.Host == "127.0.0.1" {
		host, _, err := net.SplitHostPort(address)
		if err != nil {
			return amqp.URI{}, err
		}
		amqpURI.Host = host
	}
	return amqpURI, nil
}
