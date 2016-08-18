package main

import (
	"log"

	"github.com/king-jam/libreg"
	"github.com/king-jam/libreg/registry"
	"github.com/king-jam/libreg/registry/consul"
)

func main() {

	consul.Register()

	store, err := libreg.NewRegistry(registry.CONSUL, []string{"localhost:8500"}, nil)
	if err != nil {
		log.Printf("Could not connect to consul: %s\n", err)
		return
	}

	deregisterServices(store)
	deregisterNodes(store)
}

func deregisterServices(store registry.Registry) error {

	// Dereg services
	services, err := store.Services(nil)
	if err != nil {
		log.Printf("Could not get services: %s\n", err)
		return err
	}

	for service := range services {

		// Dont delete the consul service
		if service == "consul" {
			continue
		}

		nodes, err := store.Service(service, "", nil)
		if err != nil {
			log.Printf("Could not get service nodes: %s\n", err)
			return err
		}

		for _, node := range nodes {

			err := store.Deregister(
				&registry.CatalogDeregistration{
					Node:       node.Node,
					Address:    node.Address,
					Datacenter: "dc1",
					ServiceID:  node.ServiceID,
					CheckID:    node.ServiceID,
				}, nil)
			if err != nil {
				log.Printf("Error Deregistering node: %s\n", err)
			}

		}
	}

	return nil
}

func deregisterNodes(store registry.Registry) error {
	// Dereg Nodes
	nodes, err := store.Nodes(nil)
	if err != nil {
		log.Printf("Could not get nodes: %s\n", err)
		return err
	}

	for _, node := range nodes {

		// Don't delete consul_server nodes or localhost (local consul client)
		if node.Node == "consul_server1" ||
			node.Node == "consul_server2" ||
			node.Node == "consul_server3" ||
			node.Address == "127.0.0.1" {
			continue
		}

		err := store.Deregister(
			&registry.CatalogDeregistration{
				Node:       node.Node,
				Address:    node.Address,
				Datacenter: "dc1",
				ServiceID:  "",
				CheckID:    "",
			}, nil)
		if err != nil {
			log.Printf("Error Deregistering node: %s\n", err)
		}
	}

	return nil
}
