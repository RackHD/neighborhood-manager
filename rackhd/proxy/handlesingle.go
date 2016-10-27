package proxy

import (
	"net/http"
	"sync"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// HandleSingle is a struct for the http objects
type HandleSingle struct {
	Request *http.Request
	//	Body       interface{}
	Identifier string
	Response   *Response
	wg         *sync.WaitGroup
}

// WriteResponse is ...
func (s *HandleSingle) WriteResponse(rw http.ResponseWriter, rp runtime.Producer) {
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
	s.HandleSingleEndpoint(s.Request, rw)
	//some other function that returns my json interface message

	if err := rp.Produce(rw, s.Response.Body); err != nil {
		panic(err)
	}
}

// HandleSingleMiddleware is ...
func HandleSingleMiddleware(identifier string, r *http.Request) middleware.Responder {
	return &HandleSingle{
		Request:    r,
		Identifier: identifier,
	}
}
