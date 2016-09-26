package proxy

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Response is the internal proxy response object
type Response struct {
	StatusCode int
	Body       []byte
	RequestURL string
	Error      error
}

// NewResponse copies a http.Response into a proxy Response
func NewResponse(resp *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading Response.Body %s\n", err)
		return nil, err
	}
	proxyResponse := &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
		RequestURL: resp.Request.URL.String(),
		Error:      nil,
	}
	return proxyResponse, err
}

// NewResposeFromError sets errors
func NewResposeFromError(err error) *Response {
	proxyRespnse := &Response{
		StatusCode: 500,
		Error:      err,
	}
	return proxyRespnse
}

// NewRequest copies a http.Request & Header and sets the new host
func NewRequest(r *http.Request, host string) (*http.Request, error) {
	req, err := http.NewRequest(r.Method, "http://"+host+r.URL.Path, r.Body)
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
