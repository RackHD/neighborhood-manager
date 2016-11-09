package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PollersIDGetHandlerFunc turns a function with the right signature into a pollers Id get handler
type PollersIDGetHandlerFunc func(PollersIDGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PollersIDGetHandlerFunc) Handle(params PollersIDGetParams) middleware.Responder {
	return fn(params)
}

// PollersIDGetHandler interface for that can handle valid pollers Id get params
type PollersIDGetHandler interface {
	Handle(PollersIDGetParams) middleware.Responder
}

// NewPollersIDGet creates a new http.Handler for the pollers Id get operation
func NewPollersIDGet(ctx *middleware.Context, handler PollersIDGetHandler) *PollersIDGet {
	return &PollersIDGet{Context: ctx, Handler: handler}
}

/*PollersIDGet swagger:route GET /pollers/{identifier} /api/2.0 pollersIdGet

Get a specific poller

Get information associated with the specified poller, including type, run interval, command, and whether the poller is paused.


*/
type PollersIDGet struct {
	Context *middleware.Context
	Handler PollersIDGetHandler
}

func (o *PollersIDGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPollersIDGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PollersIDGetOKBody pollers ID get o k body
// swagger:model PollersIDGetOKBody
type PollersIDGetOKBody interface{}
