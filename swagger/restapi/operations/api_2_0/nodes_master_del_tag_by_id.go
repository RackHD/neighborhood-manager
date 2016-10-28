package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NodesMasterDelTagByIDHandlerFunc turns a function with the right signature into a nodes master del tag by Id handler
type NodesMasterDelTagByIDHandlerFunc func(NodesMasterDelTagByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NodesMasterDelTagByIDHandlerFunc) Handle(params NodesMasterDelTagByIDParams) middleware.Responder {
	return fn(params)
}

// NodesMasterDelTagByIDHandler interface for that can handle valid nodes master del tag by Id params
type NodesMasterDelTagByIDHandler interface {
	Handle(NodesMasterDelTagByIDParams) middleware.Responder
}

// NewNodesMasterDelTagByID creates a new http.Handler for the nodes master del tag by Id operation
func NewNodesMasterDelTagByID(ctx *middleware.Context, handler NodesMasterDelTagByIDHandler) *NodesMasterDelTagByID {
	return &NodesMasterDelTagByID{Context: ctx, Handler: handler}
}

/*NodesMasterDelTagByID swagger:route DELETE /nodes/tags/{tagName} /api/2.0 nodesMasterDelTagById

Delete tags nodes

Delete specified tag from all nodes.

*/
type NodesMasterDelTagByID struct {
	Context *middleware.Context
	Handler NodesMasterDelTagByIDHandler
}

func (o *NodesMasterDelTagByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewNodesMasterDelTagByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}