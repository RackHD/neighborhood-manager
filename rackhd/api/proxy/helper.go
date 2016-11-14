package proxy

import (
	"bytes"
	//	"encoding/json"
	//	"io"
	"io/ioutil"
	"log"
	//	"net"
	"fmt"
	"net/http"
	"sync"

	"github.com/RackHD/neighborhood-manager/rackhd/models"
	"github.com/hashicorp/go-cleanhttp"
)

// Responses is an array of Response structs
type Responses []Response

// Response is the internal proxy response object
type Response struct {
	Header     http.Header
	StatusCode int
	Body       []byte
	RequestURL string
	Error      error
}

// Err creates an error message to print
type Err struct {
	Msg string `json:"msg"`
}

var wg *sync.WaitGroup

func init() {
	wg = &sync.WaitGroup{}
}

func (r *Response) Write(in []byte) (int, error) {
	r.Body = append(r.Body[:], in[:]...)
	return len(in), nil
}

// NewResponse copies a http.Response into a proxy Response
func NewResponse(resp *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading Response.Body %s\n", err)
		return nil, err
	}
	proxyResponse := &Response{
		Header:     resp.Header,
		StatusCode: resp.StatusCode,
		Body:       body,
		RequestURL: resp.Request.URL.String(),
		Error:      nil,
	}
	return proxyResponse, err
}

// NewResponseFromError sets errors
func NewResponseFromError(err error) *Response {
	proxyRespnse := &Response{
		StatusCode: 500,
		Error:      err,
	}
	return proxyRespnse
}

// NewRequest copies a http.Request & Header and sets the new host
func NewRequest(r *http.Request, host string) (*http.Request, error) {
	buff, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading Request.Body %s\n", err)
		return nil, err
	}
	reader := bytes.NewReader(buff)
	req, err := http.NewRequest(r.Method, "http://"+host+r.URL.Path, reader)
	if err != nil {
		return nil, err
	}
	for k, v := range r.Header {
		for _, value := range v {
			req.Header.Set(k, value)
		}
	}
	return req, nil
}

// GetResponses makes channels for the response and errors from http.Get.
// A go func is spun up for each http.Get and the responses are fed
// into their respective channels.
func GetResponses(r *http.Request, addrs map[string]struct{}) Responses {
	cr := make(chan *Response, len(addrs))
	for entry := range addrs {
		wg.Add(1)
		go func(entry string, r *http.Request) {
			defer wg.Done()
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
	wg.Wait()
	close(cr)
	var allResp Responses
	for entry := range cr {
		allResp = append(allResp, *entry)
	}
	return allResp
}

// RespCheck identifies the type of initialResp.Body and writes to the ResponseWriter.
func RespCheck(allResp Responses, resp *Response) {
	var cutSize int
	if len(allResp) <= 1 {
		resp.Header = allResp[0].Header
	}
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
		if resp.StatusCode > r.StatusCode {
			resp.StatusCode = r.StatusCode
		}
	}
	resp.Write([]byte("]"))
}

// GetStoredAddresses calls GetAddresses and returns a map of addresses
func GetStoredAddresses(identifier string) (map[string]struct{}, error) {
	if identifier != "" {
		//GO fetch one endpoint
		return nil, nil
	}
	fmt.Printf("HERE=====>>> %s\n\n", identifier)
	addrMap := make(map[string]struct{})
	rHDs, err := models.GetAllRhd()
	if err != nil {
		fmt.Printf("HERE1 ===>> %s\n", err)

		return nil, err
	}
	for _, object := range rHDs {
		addrMap[object.HTTPConf.URL.Host] = struct{}{}
	}
	fmt.Printf("HERE=====>>> ADDRMAP %+v\n\n", addrMap)
	return addrMap, nil
}

// GetQuery checks if there is a query and retrives it
func GetQuery(queryKey string, r *http.Request, rw http.ResponseWriter) (map[string]struct{}, error) {
	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	querySlice := r.URL.Query()
	if len(querySlice[queryKey]) > 0 {
		queryMap := make(map[string]struct{})
		for _, elem := range querySlice[queryKey] {
			// ip, port, err := net.SplitHostPort(elem)
			// if err != nil {
			// 	return nil, err
			// }
			// if net.ParseIP(ip) != nil {
			// 	queryMap[ip+":"+port] = struct{}{}
			// }
			queryMap[elem] = struct{}{}
		}
		return queryMap, nil
	}
	return nil, nil
}

//
// // DecodeBody is used to JSON decode a body
// func DecodeBody(resp *Response, out interface{}) error {
// 	dec := json.NewDecoder(resp.Body)
// 	return dec.Decode(out)
// }
//
// // EncodeBody is used to encode a request body
// func EncodeBody(obj interface{}) (io.Reader, error) {
// 	if obj == nil {
// 		return nil, nil
// 	}
// 	buf := bytes.NewBuffer(nil)
// 	enc := json.NewEncoder(buf)
// 	if err := enc.Encode(obj); err != nil {
// 		return nil, err
// 	}
// 	return buf, nil
// }
