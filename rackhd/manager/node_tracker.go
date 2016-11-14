package manager

import (
	"encoding/json"
	"fmt"
	"github.com/RackHD/neighborhood-manager/rackhd/manager/amqp"
	"github.com/RackHD/neighborhood-manager/rackhd/models"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/streadway/amqp"
	"net/http"
	"net/url"
	"sync"
)

// Node is a struct to unmarshal RHD Node object
type Node struct {
	NodeID string `json:"id"`
}

// Nodes is a slice of Node structs
type Nodes []Node

// NodeEvent is a RHD event generated fron the on.events exchange
type NodeEvent struct {
	Type     string `json:"type"`
	Action   string `json:"action"`
	NodeID   string `json:"nodeId"`
	NodeType string `json:"nodeType"`
}

// NodeTracker stores the state data of managed instances
type NodeTracker struct {
	stop chan struct{}
	wg   sync.WaitGroup
	UUID string
}

// NewNodeTracker creates go routines for days
func NewNodeTracker(uuid string, uri amqp.URI) *NodeTracker {

	nt := &NodeTracker{
		stop: make(chan struct{}),
		wg:   sync.WaitGroup{},
		UUID: uuid,
	}
	nt.wg.Add(1)
	go func(n *NodeTracker) {
		defer n.wg.Done()
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
			case <-n.stop:
				return
			default:
				for event := range events {
					n.processNodeEvent(uuid, &event)
				}
			}
		}
	}(nt)

	return nt
}

func (n *NodeTracker) processNodeEvent(uuid string, m *amqp.Delivery) error {
	var event NodeEvent
	if err := json.Unmarshal(m.Body, &event); err != nil {
		fmt.Printf("Failed to unmarshal")
	}
	node, err := models.NewRhdNode(uuid, event.NodeID)
	if err != nil {
		return err
	}

	switch event.Action {
	case "discovered":
	case "added":
		return models.CreateNode(node)
	case "removed":
		return models.DeleteNode(node)
	default:
	}
	m.Ack(true)
	return nil
}

// InitiateNodeCache uses the RackHD REST API to initiate our index
func (n *NodeTracker) InitiateNodeCache(address string) error {
	client := cleanhttp.DefaultClient()
	uri := &url.URL{
		Scheme: "http",
		Host:   address,
		Path:   "/api/2.0/nodes",
	}

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return fmt.Errorf("Failed to create /nodes request")
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to get RHD nodes")
	}
	var nodes Nodes
	if err = json.NewDecoder(resp.Body).Decode(&nodes); err != nil {
		return fmt.Errorf("Failed to decode RHD configuration")
	}
	for _, node := range nodes {
		rhdNode, err := models.NewRhdNode(n.UUID, node.NodeID)
		if err != nil {
			return err
		}
		models.CreateNode(rhdNode)
	}
	return nil
}
