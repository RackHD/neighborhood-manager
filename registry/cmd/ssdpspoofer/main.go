package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/king-jam/gossdp"
	"github.com/spf13/viper"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	viper.SetConfigName("ssdpSpoof")
	viper.SetConfigType("json")
	viper.AddConfigPath(dir)
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("SSDP Spoofer Configuration Error: %s\n", err)
	}

	log.Printf("SSDP Spoofer Configuration: %s\n", viper.ConfigFileUsed())

	ssdp, err := gossdp.NewSsdp(nil)
	if err != nil {
		log.Fatalf("Error creating SSDP Server: %s\n", err)
	}

	rGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	nodes := viper.GetStringMapString("ssdp")
	for node := range nodes {
		IPs := viper.GetStringSlice("ssdp." + node + ".broadcastIP")
		urns := viper.GetStringSlice("ssdp." + node + ".urnList")
		ip := IPs[0]
		uuid := strconv.Itoa(rGen.Int())
		log.Println(ip)
		if node == "NOUUID" {
			uuid = ""
		}

		for _, urn := range urns {
			server := gossdp.AdvertisableServer{
				ServiceType: urn,
				DeviceUuid:  uuid,
				Location:    fmt.Sprintf("%s%s%s", "http://", ip, "/fakepath"),
				MaxAge:      30,
			}
			ssdp.AdvertiseServer(server)

		}
	}

	defer ssdp.Stop()
	ssdp.Start()
}
