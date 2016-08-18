package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"

	"github.com/RackHD/NeighborhoodManager/rackhd/cmd/utils/api"
	"github.com/king-jam/libreg/registry"
)

var binaryName, buildDate, buildUser, commitHash, goVersion, osArch, releaseVersion string
var backendAddr, endpointAddr, serviceName string

// init takes in configurable flags
func init() {
	flag.StringVar(&backendAddr, "backend-address", "127.0.0.1:8500", "address:port of the backend store")
	flag.StringVar(&endpointAddr, "endpoint-address", "http://0.0.0.0:10001", "http://address:port of the endpoint server")
	flag.StringVar(&serviceName, "service-name", "RackHD-service:api:2.0:TEST", "The service being spoofed")
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

	log.Println(binaryName)
	log.Println("  Release version: " + releaseVersion)
	log.Println("  Built On: " + buildDate)
	log.Println("  Build By: " + buildUser)
	log.Println("  Commit Hash: " + commitHash)
	log.Println("  Go version: " + goVersion)
	log.Println("  OS/Arch: " + osArch)

	flag.Parse()

	// Parse proxyAddr
	endpointIP, endpointPort, err := extractIPPort(endpointAddr)
	if err != nil {
		log.Fatalf("Error parsing endpoint address: %s\n", err)
	}

	// Proxy server configuration
	h, err := api.NewServer(endpointIP, serviceName, "dc1", backendAddr, registry.CONSUL, endpointPort)
	if err != nil {
		log.Fatalf("Endpoint server configuration failed: %s\n", err)
	}
	h.Register("dc1", serviceName)

	fmt.Printf("Endpoint is served on => %s:%d\n", h.Address, h.Port)

	// Start the handler server for the proxy endpoint
	h.Serve()
}
