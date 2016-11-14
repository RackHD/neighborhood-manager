package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PollersLibGetHandlerFunc turns a function with the right signature into a pollers lib get handler
type PollersLibGetHandlerFunc func(context.Context, PollersLibGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PollersLibGetHandlerFunc) Handle(ctx context.Context, params PollersLibGetParams) middleware.Responder {
	return fn(ctx, params)
}

// PollersLibGetHandler interface for that can handle valid pollers lib get params
type PollersLibGetHandler interface {
	Handle(context.Context, PollersLibGetParams) middleware.Responder
}

// NewPollersLibGet creates a new http.Handler for the pollers lib get operation
func NewPollersLibGet(ctx *middleware.Context, handler PollersLibGetHandler) *PollersLibGet {
	return &PollersLibGet{Context: ctx, Handler: handler}
}

/*PollersLibGet swagger:route GET /pollers/library /api/2.0 pollersLibGet

Get a list of possible pollers

Get a list of all available poller definitions in the poller library.

*/
type PollersLibGet struct {
	Context *middleware.Context
	Handler PollersLibGetHandler
}

func (o *PollersLibGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPollersLibGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PollersLibGetOKBody pollers lib get o k body
// swagger:model PollersLibGetOKBody
type PollersLibGetOKBody interface{}
