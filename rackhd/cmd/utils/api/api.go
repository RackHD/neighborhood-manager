package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/king-jam/libreg"
	regStore "github.com/king-jam/libreg/registry"
	"github.com/king-jam/libreg/registry/consul"
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

	return s, nil
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
			Thing2:  "moar things go here",
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
		Thing2:  "moar things go here :)",
		Numbers: rand.Intn(90000),
	}

	json.NewEncoder(w).Encode(object)
}

// Register is...
func (e *Server) Register(datacenter, serviceName string) {
	n := fmt.Sprintf("%d", rand.Int())
	_ = e.Store.Register(&regStore.CatalogRegistration{
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
	}, nil)

}