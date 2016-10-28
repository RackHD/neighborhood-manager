package proxy

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// HandleAll is a struct for the http objects
type HandleAll struct {
	Request *http.Request
	//	Body       interface{}
	//	Identifier string
	Response *Response
	wg       *sync.WaitGroup
}

// HandleAllMiddleware handles the route call
func HandleAllMiddleware(r *http.Request) middleware.Responder {
	return &HandleAll{
		Request: r,
	}
}

// WriteResponse is ...
func (a *HandleAll) WriteResponse(rw http.ResponseWriter, rp runtime.Producer) {
	addrMap, err := GetStoredAddresses("")
	resp := Response{}
	if len(addrMap) == 0 {

		resp.Body = []byte("{No endpoints under management.}")
		resp.StatusCode = http.StatusOK

	} else if err != nil {

		resp.StatusCode = http.StatusInternalServerError
		resp.Body = []byte(fmt.Sprintf("{%s}", err))

	} else {

		oneResp := GetResponses(a.Request, addrMap)
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
