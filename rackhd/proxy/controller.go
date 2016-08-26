package proxy

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"

	regStore "github.com/RackHD/neighborhood-manager/libreg/registry"
	"github.com/RackHD/neighborhood-manager/libreg/registry/consul"
	"github.com/RackHD/neighborhood-manager/rackhd/watcher"
)

// Server is the proxy server struct
type Server struct {
	Address string
	Port    int
	Store   *watcher.Monitor
	wg      *sync.WaitGroup
}

// Serve starts the Server on address:port and handles the routes
func (p *Server) Serve() {
	m := http.NewServeMux()
	m.HandleFunc("/test", p.HandleTest)
	m.HandleFunc("/", p.HandleNodes)
	http.ListenAndServe(fmt.Sprintf("%s:%d", p.Address, p.Port), m)
}

// NewServer initializes a new Server
func NewServer(proxyIP, serviceName, datacenter, backendAddr string, backend regStore.Backend, proxyPort int) (*Server, error) {
	consul.Register()
	m, err := watcher.NewMonitor(serviceName, datacenter, backendAddr, backend)
	if err != nil {
		return nil, err
	}
	proxyServer := &Server{
		Address: proxyIP,
		Port:    proxyPort,
		Store:   m,
		wg:      &sync.WaitGroup{},
	}
	return proxyServer, nil
}

// HandleTest is....well a test
func (p *Server) HandleTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

// HandleNodes sends, recieves, and processes all the data
func (p *Server) HandleNodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	addrMap := p.GetAddresses(w, r)
	cr, _ := p.GetResp(r, addrMap)
	p.RespCheck(w, cr)
}

// GetResp makes channels for the response and errors from http.Get.
// A go func is spun up for each http.Get and the responses are fed
// into their respective channels.
func (p *Server) GetResp(r *http.Request, addrs map[string]struct{}) (chan *Response, chan error) {
	cr := make(chan *Response, len(addrs))
	errs := make(chan error, len(addrs))
	defer close(cr)
	defer close(errs)
	for entry := range addrs {
		p.wg.Add(1)
		go func(entry string) {
			defer p.wg.Done()
			r.URL.Host = entry
			r.URL.Scheme = "http"
			fmt.Printf("url string %s\n", r.URL.String())
			respGet, err := http.Get(r.URL.String())
			if err != nil {
				errs <- fmt.Errorf("Could not send any HTTP Get requests: %s\n", err)
				return
			}
			responseCopy, err := NewResponse(respGet)
			if err != nil {
				log.Printf("Error copying response => %s\n", err)
			}
			cr <- responseCopy
			respGet.Body.Close()
		}(entry)
	}
	p.wg.Wait()
	return cr, errs
}

// GetAddresses decides from where to retrieve the addresses
func (p *Server) GetAddresses(w http.ResponseWriter, r *http.Request) map[string]struct{} {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	querySlice := r.URL.Query()
	if len(querySlice["ip"]) > 0 {
		addrMap := p.GetQueryAddresses(querySlice["ip"])
		return addrMap
	}
	addrMap := p.GetStoredAddresses()
	return addrMap
}

// GetStoredAddresses calls GetAddresses and returns a map of addresses
func (p *Server) GetStoredAddresses() map[string]struct{} {
	addresses, err := p.Store.GetAddresses()
	if err != nil {
		log.Printf("Did not get IP List ==> %s\n", err)
	}
	return addresses
}

// GetQueryAddresses retrives a url flag and returns a map of address(es)
func (p *Server) GetQueryAddresses(querySlice []string) map[string]struct{} {
	fmt.Println(querySlice)
	queryMap := make(map[string]struct{})
	for _, elem := range querySlice {
		ip, port, err := net.SplitHostPort(elem)
		if err != nil {
			log.Printf("Invalid port => %s\n", err)
			return nil
		}
		if net.ParseIP(ip) != nil {
			queryMap[ip+":"+port] = struct{}{}
		}
	}
	return queryMap
}

// RespCheck identifies the type of initialResp.Body and calls the appropriate
// helper function to write to the ResponseWriter.
func (p *Server) RespCheck(w http.ResponseWriter, cr chan *Response) {
	initialResp := <-cr
	if initialResp.Body[0] != '[' {
		w.Write(initialResp.Body)
		return
	}
	w.Write(initialResp.Body[0 : len(initialResp.Body)-2])
	for r := range cr {
		w.Write([]byte(","))
		w.Write(r.Body[1 : len(r.Body)-2])
	}
	w.Write([]byte("]"))
}
