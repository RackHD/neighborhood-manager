package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/RackHD/neighborhood-manager/libreg"
	regStore "github.com/RackHD/neighborhood-manager/libreg/registry"
	"github.com/RackHD/neighborhood-manager/libreg/registry/consul"
)

// Server is the proxy server struct
type Server struct {
	Address string
	Port    int
	Store   regStore.Registry
}

//NodeObject is a object struct
type NodeObject struct {
	Name    string    `json:"name"`
	Thing1  string    `json:"thing1"`
	Thing2  string    `json:"thing2"`
	Numbers int       `json:"numbers"`
	Time    time.Time `json:"time"`
}

// Serve starts the Server on address:port and handles the routes
func (e *Server) Serve() {
	m := http.NewServeMux()
	m.HandleFunc("/test", e.HandleTest)
	m.HandleFunc("/object", e.HandleServeObject)
	m.HandleFunc("/array", e.HandleServeArray)
	http.ListenAndServe(fmt.Sprintf("%s:%d", e.Address, e.Port), m)
}

// NewServer initializes a new Server
func NewServer(endpointIP, serviceName, datacenter, backendAddr string, backend regStore.Backend, endpointPort int) (*Server, error) {
	consul.Register()
	s := &Server{}
	r, err := libreg.NewRegistry(backend, []string{backendAddr}, nil)
	if err != nil {
		log.Printf("Error creating backend store: %s\n", err)
		return nil, err
	}
	s.Store = r
	s.Address = endpointIP
	s.Port = endpointPort

	for i := 0; i < 5; i++ {
		leader, err := s.Store.Leader()
		if err == nil && leader != "" {
			return s, nil
		}
		time.Sleep(5 * time.Second)
	}

	return nil, errors.New("Unable to find backend cluster")
}

// HandleTest is....well a test
func (e *Server) HandleTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

// HandleServeArray serves an array back to the Http Endpoint
func (e *Server) HandleServeArray(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var nodes []NodeObject
	n := rand.Intn(100)
	for i := 1; i <= n; i++ {
		node := NodeObject{
			Name:    fmt.Sprintf("TESTNODE: %+v", rand.Intn(100)),
			Thing1:  "arrays for days",
			Thing2:  e.Address,
			Numbers: i,
		}
		nodes = append(nodes, node)
	}

	json.NewEncoder(w).Encode(nodes)
}

// HandleServeObject serves an object back to the Http Endpint
func (e *Server) HandleServeObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	name := fmt.Sprintf("TESTNODE: %+v", rand.Intn(100))
	object := NodeObject{
		Name:    name,
		Thing1:  "objects FTW",
		Thing2:  e.Address,
		Numbers: rand.Intn(90000),
	}
	json.NewEncoder(w).Encode(object)
}

// Register is...
func (e *Server) Register(datacenter, serviceName string) {
	rGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := fmt.Sprintf("%d", rGen.Int())
	if err := e.Store.Register(&regStore.CatalogRegistration{
		Node:       n,
		Address:    e.Address,
		Datacenter: datacenter,
		Service: &regStore.AgentService{
			ID:      serviceName,
			Service: serviceName,
			Port:    e.Port,
			Address: e.Address,
		},
		Check: &regStore.AgentCheck{
			Node:        n,
			CheckID:     "service:" + serviceName,
			Name:        "Service '" + serviceName + "' check",
			Status:      "passing",
			ServiceID:   serviceName,
			ServiceName: serviceName,
		},
	}, nil); err != nil {
		log.Printf("Error registering serviceName: %s\n", err)
	}

}
