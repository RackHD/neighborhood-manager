package models

import (
	"fmt"
	"github.com/streadway/amqp"
	"net/url"
	"regexp"
	"strings"
)

const (
	// managePrefix is a directory structure that the data will be kept under
	managerPrefix = "rhdman/"

	// rhdPrefix is a directory shortcut to where all RHD objects are stored
	rhdPrefix = managerPrefix + "rhd/"

	// httpPrefix is a directory under a RHD ID where the http URL is stored
	httpPrefix = "httpConf"

	// amqpPrefix is a directory under a RHD ID where the amqp URL is stroed
	amqpPrefix = "amqpConf"
)

var (
	//
	rx *regexp.Regexp
)

func init() {
	rx = regexp.MustCompile(`rhdman\/rhd\/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`)
}

// RackHD stores the relevant data about RackHD nodes under management
type RackHD struct {
	ID       string
	HTTPConf HTTPConfig
	AmqpConf AmqpConfig
	Nodes    []*RhdNode // can be found in the nodes.go file
}

// HTTPConfig struct wrapper of URL for future expansion
type HTTPConfig struct {
	URL *url.URL
}

// AmqpConfig struct wrapper of amqp URI for future expansion
type AmqpConfig struct {
	URI amqp.URI
}

// NewRhd creates a RackHD struct for storage
func NewRhd(id string, httpURL string, amqpURI string) (*RackHD, error) {
	rhd := &RackHD{}
	rhd.ID = id
	url, err := url.Parse(httpURL)
	if err != nil {
		return rhd, fmt.Errorf("Failed to parse HTTP URL")
	}
	rhd.HTTPConf = HTTPConfig{
		URL: url,
	}
	amqp, err := amqp.ParseURI(amqpURI)
	if err != nil {
		return rhd, fmt.Errorf("Failed to parse AMQP URI: %s", err)
	}
	rhd.AmqpConf = AmqpConfig{
		URI: amqp,
	}
	return rhd, nil
}

// CreateRhd stores a RackHD instance on the backend
func CreateRhd(rhd *RackHD) error {
	return UpdateRhd(rhd)
}

// UpdateRhd updates a RackHD instance on the backend
func UpdateRhd(rhd *RackHD) error {
	basePath := rhdPrefix + rhd.ID + "/"
	if err := db.Put(basePath+httpPrefix, []byte(rhd.HTTPConf.URL.String()), nil); err != nil {
		return err
	}
	return db.Put(basePath+amqpPrefix, []byte(rhd.AmqpConf.URI.String()), nil)
}

// GetAllRhd returns all RackHD object stored in the backend
func GetAllRhd() ([]*RackHD, error) {
	var rackhds []*RackHD
	entries, err := db.List(rhdPrefix)
	if err != nil {
		return rackhds, fmt.Errorf("failed")
	}
	lastID := ""
	for _, pair := range entries {
		key := rx.FindString(pair.Key)
		if key != "" {
			tempArr := strings.Split(key, rhdPrefix)
			id := tempArr[len(tempArr)-1]
			// TODO handle 0 length
			if id == lastID {
				// we already saw this one
				continue
			}
			instance, err := getRhdInternal(id)
			if err != nil {
				fmt.Printf("Failed to get RHD")
			}
			rackhds = append(rackhds, instance)
			lastID = id
		}
	}
	return rackhds, nil
}

func getRhdInternal(id string) (*RackHD, error) {
	basePath := rhdPrefix + id + "/"
	rhd := &RackHD{
		ID: id,
	}
	entries, err := db.List(basePath)
	if err != nil {
		return rhd, fmt.Errorf("failed")
	}
	for _, pair := range entries {
		switch pair.Key {
		case basePath + httpPrefix:
			u, _ := url.Parse(string(pair.Value))
			// TODO handle err
			rhd.HTTPConf = HTTPConfig{
				URL: u,
			}
		case basePath + amqpPrefix:
			a, _ := amqp.ParseURI(string(pair.Value))
			// TODO handle err
			rhd.AmqpConf = AmqpConfig{
				URI: a,
			}
		default:
		}
	}
	return rhd, nil
}

// GetRhdByID returns a RackHD based on its unique ID
func GetRhdByID(id string) (*RackHD, error) {
	instance, err := getRhdInternal(id)
	if err != nil {
		return &RackHD{}, err
	}
	return instance, nil
}

// GetRhdsByIDs returns a group of RackHD structs based on an array of unique IDs
func GetRhdsByIDs(ids []string) ([]*RackHD, error) {
	var rackhds []*RackHD
	for _, id := range ids {
		instance, err := getRhdInternal(id)
		if err != nil {
			continue
		}
		rackhds = append(rackhds, instance)
	}
	return rackhds, nil
}

// DeleteRhdByID removes a RackHD from the backend
func DeleteRhdByID(id string) error {
	basePath := rhdPrefix + id
	return db.DeleteTree(basePath)
}

// DeleteRhdsByIDs removes multiple RackHDs from the backend
func DeleteRhdsByIDs(ids []string) error {
	for _, id := range ids {
		if err := DeleteRhdByID(id); err != nil {
			return err
		}
	}
	return nil
}
