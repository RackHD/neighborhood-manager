package proxy

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var dat map[string]interface{}

// DecodeBody decodes the request body into json
func DecodeBody(r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dat); err != nil {
		log.Printf("Error decoding Json body => %s\n", err)
	}
	fmt.Printf("%v\n", dat["name"])
	ReplaceValue(dat, "")
}

// ReplaceValue loops over nested maps and replaces the value for the passed key
func ReplaceValue(d map[string]interface{}, value string) {
	net := dat["options"].(map[string]interface{})["defaults"].(map[string]interface{})["networkDevices"].(map[string]interface{})[])

	fmt.Printf("%v\n", net)
}
