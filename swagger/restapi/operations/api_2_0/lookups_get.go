package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// LookupsGetHandlerFunc turns a function with the right signature into a lookups get handler
type LookupsGetHandlerFunc func(LookupsGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LookupsGetHandlerFunc) Handle(params LookupsGetParams) middleware.Responder {
	return fn(params)
}

// LookupsGetHandler interface for that can handle valid lookups get params
type LookupsGetHandler interface {
	Handle(LookupsGetParams) middleware.Responder
}

// NewLookupsGet creates a new http.Handler for the lookups get operation
func NewLookupsGet(ctx *middleware.Context, handler LookupsGetHandler) *LookupsGet {
	return &LookupsGet{Context: ctx, Handler: handler}
}

/*LookupsGet swagger:route GET /lookups /api/2.0 lookupsGet

Get a list of lookups

Get a list of all lookups currently stored. Lookups relate mac addresses to ip addresses.

*/
type LookupsGet struct {
	Context *middleware.Context
	Handler LookupsGetHandler
}

func (o *LookupsGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewLookupsGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
