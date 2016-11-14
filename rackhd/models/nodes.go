package models

import (
	"encoding/json"
)

const (
	// nodesPrefix is the area under rhd where nodes are stored
	nodesPrefix = "nodes/"
)

// RhdNode stores relevant cache data about a node
type RhdNode struct {
	ID    string
	RhdID string
}

// NewRhdNode creates a new RhdNode object for storage
func NewRhdNode(rhdID string, nodeID string) (*RhdNode, error) {
	return &RhdNode{
		RhdID: rhdID,
		ID:    nodeID,
	}, nil
}

// CreateNode allows creating a node on the backend
func CreateNode(node *RhdNode) error {
	return UpdateNode(node)
}

// UpdateNode updates a RhdNode on the backend
func UpdateNode(node *RhdNode) error {
	nodePath := rhdPrefix + node.RhdID + "/" + nodesPrefix + node.ID
	if err := CreateNodeCache(node); err != nil {
		return err
	}
	b, err := json.Marshal(node)
	if err != nil {
		return err
	}
	return db.Put(nodePath, b, nil)
}

// DeleteNode removes RhdNode on the backend
func DeleteNode(node *RhdNode) error {
	nodePath := rhdPrefix + node.RhdID + "/" + nodesPrefix + node.ID
	if err := DeleteNodeCache(node.ID); err != nil {
		return err
	}
	return db.DeleteTree(nodePath)
}

// GetAllNodes is currently stubbed out but unsupported
func GetAllNodes() ([]*RhdNode, error) {
	var nodes []*RhdNode
	return nodes, nil
}

// GetAllNodesByRhdID is currently stubbed out but unsupported
func GetAllNodesByRhdID(id string) ([]*RhdNode, error) {
	var rhds []*RhdNode
	return rhds, nil
}

// GetNodeByRhdIDByNodeID is currently stubbed out but unsupported
func GetNodeByRhdIDByNodeID(rid string, nid string) (*RhdNode, error) {
	return &RhdNode{}, nil
}

// GetNodesByRhdIDByNodeIDs is currently stubbed out but unsupported
func GetNodesByRhdIDByNodeIDs(rid string, ids []string) ([]*RhdNode, error) {
	var nodes []*RhdNode
	return nodes, nil
}

// CreateNodeCache stores the cache layer lookup functionality
func CreateNodeCache(node *RhdNode) error {
	nodePath := nodesPrefix + node.ID
	return db.Put(nodePath, []byte(node.RhdID), nil)
}

// GetRhdIDByNodeID returns the RHD ID for a node
func GetRhdIDByNodeID(id string) (string, error) {
	nodesPath := nodesPrefix + id
	pair, err := db.Get(nodesPath)
	return string(pair.Value), err
}

// DeleteNodeCache removes a cache index
func DeleteNodeCache(id string) error {
	nodePath := nodesPrefix + id
	return db.DeleteTree(nodePath)
}
