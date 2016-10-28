package proxy

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// HandleSingle is a struct for the http objects
type HandleSingle struct {
	Request *http.Request
	//	Body       interface{}
	Identifier string
	Response   *Response
}

// HandleSingleMiddleware handles the rotue call
func HandleSingleMiddleware(identifier string, r *http.Request) middleware.Responder {
	return &HandleSingle{
		Request:    r,
		Identifier: identifier,
	}
}

// WriteResponse creates and writes back the response to the middleware handler
func (s *HandleSingle) WriteResponse(rw http.ResponseWriter, rp runtime.Producer) {
	addrMap, err := GetStoredAddresses(s.Identifier)
	resp := Response{}
	if len(addrMap) == 0 {

		resp.Body = []byte("{No endpoints under management.}")
		resp.StatusCode = http.StatusOK

	} else if err != nil {

		resp.StatusCode = http.StatusInternalServerError
		resp.Body = []byte(fmt.Sprintf("{%s}", err))

	} else {

		oneResp := GetResponses(s.Request, addrMap)
		RespCheck(oneResp, &resp)

	}

	for k, v := range resp.Header {
		for _, val := range v {
			rw.Header().Add(k, val)
		}
	}
	// rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if resp.StatusCode > 0 {
		rw.WriteHeader(resp.StatusCode)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	if err := rp.Produce(rw, resp.Body); err != nil {
		panic(err)
	}
}
