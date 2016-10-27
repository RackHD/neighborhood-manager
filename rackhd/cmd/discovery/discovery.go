package main

import (
	"flag"
	"fmt"
	"log"

	regStore "github.com/RackHD/neighborhood-manager/libreg/registry"
	"github.com/RackHD/neighborhood-manager/libreg/registry/consul"
	"github.com/RackHD/neighborhood-manager/rackhd/watcher"
)

var binaryName, buildDate, buildUser, commitHash, goVersion, osArch, releaseVersion string
var backendAddr, serviceName, datacenter string

// init takes in configurable flags
func init() {
	flag.StringVar(&backendAddr, "backend-address", "127.0.0.1:8500", "address:port of the backend store")
	flag.StringVar(&serviceName, "service-name", "RackHD-service:api:2.0", "The service being proxied")
	flag.StringVar(&datacenter, "datacenter", "dc1", "The consul datacenter string")
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

	consul.Register()
	msgChan := make(chan *watcher.Message)
	_, err := watcher.NewMonitor(serviceName, datacenter, backendAddr, regStore.CONSUL, msgChan)
	if err != nil {
		fmt.Printf("Error creating monitor %s\n", err)
	}

	for entry := range msgChan {
		fmt.Printf("Channel Data & Events => %+v\n\n", entry)
	}

}
