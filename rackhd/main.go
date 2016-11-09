package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	loads "github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	regStore "github.com/RackHD/neighborhood-manager/libreg/registry"
	"github.com/RackHD/neighborhood-manager/libreg/registry/consul"
	"github.com/RackHD/neighborhood-manager/rackhd/manager"
	"github.com/RackHD/neighborhood-manager/rackhd/models"
	"github.com/RackHD/neighborhood-manager/rackhd/watcher"
	"github.com/RackHD/neighborhood-manager/swagger/restapi"
	"github.com/RackHD/neighborhood-manager/swagger/restapi/operations"
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

	// Temp disabled for testing (hard code flag vars)
	//	flag.Parse()

	consul.Register()
	models.InitBackend()
	msgChan := make(chan *watcher.Message)
	_, err := watcher.NewMonitor(serviceName, datacenter, backendAddr, regStore.CONSUL, msgChan)
	if err != nil {
		fmt.Printf("Error creating monitor %s\n", err)
	}
	_, err = rhdman.NewRackHDMan(msgChan)

	// New manager
	// pass channel into manager
	// manager writes into consul k,v

	// for entry := range msgChan {
	// 	fmt.Printf("Channel Data & Events => %+v\n\n", entry)
	// }

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewRackHD20API(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = `RackHD 2.0`
	parser.LongDescription = swaggerSpec.Spec().Info.Description

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
