package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	regStore "github.com/RackHD/NeighborhoodManager/libreg/registry"
	"github.com/RackHD/NeighborhoodManager/libreg/registry/consul"
	"github.com/RackHD/NeighborhoodManager/rackhd/watcher"
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
	addresses, err := p.Store.GetAddresses()
	if err != nil {
		log.Printf("Did not get IP List ==> %s\n", err)
	}
	cr, _ := p.GetResp(r, addresses)

	for range cr {
		p.RespCheck(w, cr)
	}
	return
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
			fmt.Printf("%+v\n\n", respGet)
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

// RespCheck identifies the type of initialResp.Body and calls the appropriate
// helper function to write to the ResponseWriter.
func (p *Server) RespCheck(w http.ResponseWriter, c chan *Response) {
	initialResp := <-c
	w.WriteHeader(initialResp.StatusCode)
	if initialResp.Body[0] == '[' {
		p.CombineResp(c, initialResp, w)
	} else {
		p.PassResp(initialResp, w)
	}
}

// CombineResp combines many response.Body objects and formats them correctly
// with json structuring.
func (p *Server) CombineResp(cr chan *Response, initialResp *Response, w http.ResponseWriter) {
	w.Write(initialResp.Body[0 : len(initialResp.Body)-1])
	for r := range cr {
		w.Write([]byte(","))
		w.Write(r.Body[1 : len(r.Body)-2])
	}
	w.Write([]byte("]"))
}

// PassResp takes a response.Body object and writes it to the ResponseWriter
func (p *Server) PassResp(initialResp *Response, w http.ResponseWriter) {
	w.Write(initialResp.Body)
}
