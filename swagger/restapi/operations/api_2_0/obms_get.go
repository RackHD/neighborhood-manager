package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ObmsGetHandlerFunc turns a function with the right signature into a obms get handler
type ObmsGetHandlerFunc func(ObmsGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ObmsGetHandlerFunc) Handle(params ObmsGetParams) middleware.Responder {
	return fn(params)
}

// ObmsGetHandler interface for that can handle valid obms get params
type ObmsGetHandler interface {
	Handle(ObmsGetParams) middleware.Responder
}

// NewObmsGet creates a new http.Handler for the obms get operation
func NewObmsGet(ctx *middleware.Context, handler ObmsGetHandler) *ObmsGet {
	return &ObmsGet{Context: ctx, Handler: handler}
}

/*ObmsGet swagger:route GET /obms /api/2.0 obmsGet

Get list of all OBM service instances

Get a list of all OBM settings that have been associated with nodes. OBM settings allow RackHD to communicate with the BMC of a node.


*/
type ObmsGet struct {
	Context *middleware.Context
	Handler ObmsGetHandler
}

func (o *ObmsGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewObmsGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ObmsGetOKBody obms get o k body
// swagger:model ObmsGetOKBody
type ObmsGetOKBody interface{}