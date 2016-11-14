package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SkusIDGetNodesHandlerFunc turns a function with the right signature into a skus Id get nodes handler
type SkusIDGetNodesHandlerFunc func(context.Context, SkusIDGetNodesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SkusIDGetNodesHandlerFunc) Handle(ctx context.Context, params SkusIDGetNodesParams) middleware.Responder {
	return fn(ctx, params)
}

// SkusIDGetNodesHandler interface for that can handle valid skus Id get nodes params
type SkusIDGetNodesHandler interface {
	Handle(context.Context, SkusIDGetNodesParams) middleware.Responder
}

// NewSkusIDGetNodes creates a new http.Handler for the skus Id get nodes operation
func NewSkusIDGetNodes(ctx *middleware.Context, handler SkusIDGetNodesHandler) *SkusIDGetNodes {
	return &SkusIDGetNodes{Context: ctx, Handler: handler}
}

/*SkusIDGetNodes swagger:route GET /skus/{identifier}/nodes /api/2.0 skusIdGetNodes

Get nodes for specific SKU

Get the nodes associated with the specified SKU definition.

*/
type SkusIDGetNodes struct {
	Context *middleware.Context
	Handler SkusIDGetNodesHandler
}

func (o *SkusIDGetNodes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewSkusIDGetNodesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SkusIDGetNodesOKBody skus ID get nodes o k body
// swagger:model SkusIDGetNodesOKBody
type SkusIDGetNodesOKBody interface{}
