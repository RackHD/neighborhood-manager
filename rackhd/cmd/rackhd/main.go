package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"

	"github.com/RackHD/NeighborhoodManager/rackhd/proxy"
	"github.com/king-jam/libreg/registry"
)

var binaryName, buildDate, buildUser, commitHash, goVersion, osArch, releaseVersion string
var backendAddr, proxyAddr, serviceName string

// init takes in configurable flags
func init() {
	flag.StringVar(&backendAddr, "backend-address", "127.0.0.1:8500", "address:port of the backend store")
	flag.StringVar(&proxyAddr, "proxy-address", "http://0.0.0.0:10001", "http://address:port of the proxy server")
	flag.StringVar(&serviceName, "service-name", "RackHD-service:api:2.0", "The service being proxied")
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
	proxyIP, proxyPort, err := extractIPPort(proxyAddr)
	if err != nil {
		log.Fatalf("Error parsing proxy address: %s\n", err)
	}

	// Proxy server configuration
	h, err := proxy.NewServer(proxyIP, serviceName, "dc-docker", backendAddr, registry.CONSUL, proxyPort)
	if err != nil {
		log.Fatalf("Proxy server configuration failed: %s\n", err)
	}

	fmt.Printf("Service name is => %s\n", serviceName)
	fmt.Printf("Proxy is served on => %s:%d\n", h.Address, h.Port)

	// Start the handler server for the proxy endpoint
	h.Serve()
}
