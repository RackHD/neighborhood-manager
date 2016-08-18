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
	}
	return proxyResponse, err
}
