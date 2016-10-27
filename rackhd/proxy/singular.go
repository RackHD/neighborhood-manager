package proxy

import (
	"net/http"

	"github.com/hashicorp/go-cleanhttp"
)

// HandleSingleEndpoint sends, recieves, and processes all the data
func (s *HandleSingle) HandleSingleEndpoint(r *http.Request, rw http.ResponseWriter) Response {
	addrMap, err := s.GetStoredAddresses()
	resp := Response{}
	if len(addrMap) == 0 {
		resp.Body = []byte("[]")
		resp.StatusCode = 200
		// Write an error saying no endpoints under managemnt
		return resp
	}
	if err != nil {
		resp.StatusCode = 500
		//TODO something here with body vs err writing back if an error exists

		// resp.Body, err = json.Marshal(Err{Msg: "Internal error fetching endpoint addresses."})
		// if err != nil {
		//   resp.Error =
		// }
		return resp
	}

	oneResp := s.GetResponse(r, addrMap)
	s.RespHeaderWriter(r, rw, oneResp)
	s.RespCheck(r, oneResp, &resp)
	return resp
}

// GetStoredAddresses calls GetAddresses and returns a map of addresses
func (s *HandleSingle) GetStoredAddresses() (map[string]struct{}, error) {
	// TODO this is where we do something with s.Identifier

	// addresses, err := p.Store.GetAddresses()
	// if err != nil {
	// 	log.Printf("Did not get IP List ==> %s\n", err)
	// }
	return nil, nil
}

// GetResponse makes channels for the response and errors from http.Get.
// A go func is spun up for each http.Get and the responses are fed
// into their respective channels.
func (s *HandleSingle) GetResponse(r *http.Request, addrs map[string]struct{}) Responses {
	cr := make(chan *Response, 1)
	for entry := range addrs {
		s.wg.Add(1)
		go func(entry string, r *http.Request) {
			defer s.wg.Done()
			req, err := NewRequest(r, entry)
			if err != nil {
				cr <- NewResponseFromError(err)
				return
			}
			client := cleanhttp.DefaultClient()
			respGet, err := client.Do(req)
			if err != nil {
				cr <- NewResponseFromError(err)
				return
			}
			defer respGet.Body.Close()
			responseCopy, err := NewResponse(respGet)
			if err != nil {
				cr <- NewResponseFromError(err)
				return
			}
			cr <- responseCopy
		}(entry, r)
	}
	s.wg.Wait()
	close(cr)
	var allResp Responses
	for entry := range cr {
		allResp = append(allResp, *entry)
	}
	return allResp
}

// RespHeaderWriter writes the StatusCode and Headers
func (s *HandleSingle) RespHeaderWriter(r *http.Request, rw http.ResponseWriter, allResp Responses) {
	// TODO change this to pass all headers as its a single response
	var status int
	status = 500
	for _, s := range allResp {
		if s.StatusCode < status {
			status = s.StatusCode
		}
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
}

// RespCheck identifies the type of initialResp.Body and writes to the ResponseWriter.
func (s *HandleSingle) RespCheck(r *http.Request, allResp Responses, resp *Response) {
	var cutSize int
	resp.Write([]byte("["))
	for i, r := range allResp {
		if r.Body == nil || ((r.Body[0] == '[') && (r.Body[1] == ']')) {
			continue
		}
		if r.Body[0] == '[' {
			cutSize = 1
		} else if r.Body[0] == '{' {
			cutSize = 0
		} else {
			continue
		}
		resp.Write(r.Body[cutSize : len(r.Body)-cutSize])
		if i != len(allResp)-1 {
			resp.Write([]byte(","))
		}
	}
	resp.Write([]byte("]"))
}
