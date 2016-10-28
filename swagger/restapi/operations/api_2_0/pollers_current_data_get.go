package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PollersCurrentDataGetHandlerFunc turns a function with the right signature into a pollers current data get handler
type PollersCurrentDataGetHandlerFunc func(PollersCurrentDataGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PollersCurrentDataGetHandlerFunc) Handle(params PollersCurrentDataGetParams) middleware.Responder {
	return fn(params)
}

// PollersCurrentDataGetHandler interface for that can handle valid pollers current data get params
type PollersCurrentDataGetHandler interface {
	Handle(PollersCurrentDataGetParams) middleware.Responder
}

// NewPollersCurrentDataGet creates a new http.Handler for the pollers current data get operation
func NewPollersCurrentDataGet(ctx *middleware.Context, handler PollersCurrentDataGetHandler) *PollersCurrentDataGet {
	return &PollersCurrentDataGet{Context: ctx, Handler: handler}
}

/*PollersCurrentDataGet swagger:route GET /pollers/{identifier}/data/current /api/2.0 pollersCurrentDataGet

Get latest data for a poller

Get latest output data generated by the poller poller with the specified identifier.

*/
type PollersCurrentDataGet struct {
	Context *middleware.Context
	Handler PollersCurrentDataGetHandler
}

func (o *PollersCurrentDataGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPollersCurrentDataGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PollersCurrentDataGetOKBody pollers current data get o k body
// swagger:model PollersCurrentDataGetOKBody
type PollersCurrentDataGetOKBody interface{}

// PollersCurrentDataGetNoContentBody pollers current data get no content body
// swagger:model PollersCurrentDataGetNoContentBody
type PollersCurrentDataGetNoContentBody interface{}