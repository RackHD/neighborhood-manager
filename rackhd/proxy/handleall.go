package proxy

import (
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

// WriteResponse is ...
func (a *HandleAll) WriteResponse(rw http.ResponseWriter, rp runtime.Producer) {
	// for k, v := range e.headers {
	// 	for _, val := range v {
	// 		rw.Header().Add(k, val)
	// 	}
	// }
	// if e.code > 0 {
	// 	rw.WriteHeader(e.code)
	// } else {
	// 	rw.WriteHeader(http.StatusInternalServerError)
	// }
	a.HandleAllEndpoints(a.Request, rw)
	//some other function that returns my json interface message

	if err := rp.Produce(rw, a.Response.Body); err != nil {
		panic(err)
	}
}

// HandleAllMiddleware is ...
func HandleAllMiddleware(r *http.Request) middleware.Responder {
	return &HandleAll{
		Request: r,
	}
}
