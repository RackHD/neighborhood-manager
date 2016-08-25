package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"

	"github.com/RackHD/neighborhood-manager/libreg/registry"
	"github.com/RackHD/neighborhood-manager/rackhd/cmd/utils/api"
)

var binaryName, buildDate, buildUser, commitHash, goVersion, osArch, releaseVersion string
var backendAddr, endpointAddr, serviceName, datacenter string

// init takes in configurable flags
func init() {
	flag.StringVar(&backendAddr, "backend-address", "127.0.0.1:8500", "address:port of the backend store")
	flag.StringVar(&endpointAddr, "endpoint-address", "http://0.0.0.0:10002", "http://address:port of the endpoint server")
	flag.StringVar(&serviceName, "service-name", "RackHD-service:api:2.0:TEST", "The service being spoofed")
	flag.StringVar(&datacenter, "datacenter", "dc1", "The consul datacenter string")
}

// extractIPPort splits the Address flag into an ip string anf port int
func extractIPPort(location string) (string, int, error) {
	addr, err := url.Parse(location)
	if err != nil {
		return "", 0, err
	}

	agentIP, portStr, err := net.SplitHostPort(addr.Host)
	if err != nil {
		return "", 0, err
	}

	agentPort, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, err
	}

	return agentIP, agentPort, nil
}

// Main
func main() {

	flag.Parse()

	// Parse proxyAddr
	endpointIP, endpointPort, err := extractIPPort(endpointAddr)
	if err != nil {
		log.Fatalf("Error parsing endpoint address: %s\n", err)
	}

	// Proxy server configuration
	h, err := api.NewServer(endpointIP, serviceName, datacenter, backendAddr, registry.CONSUL, endpointPort)
	if err != nil {
		log.Fatalf("Endpoint server configuration failed: %s\n", err)
	}

	h.Register(datacenter, serviceName)

	fmt.Printf("Endpoint is served on => %s:%d\n", h.Address, h.Port)

	// Start the handler server for the proxy endpoint
	h.Serve()
}
