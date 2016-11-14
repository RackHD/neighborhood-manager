package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NodesPutObmsByNodeIDHandlerFunc turns a function with the right signature into a nodes put obms by node Id handler
type NodesPutObmsByNodeIDHandlerFunc func(context.Context, NodesPutObmsByNodeIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NodesPutObmsByNodeIDHandlerFunc) Handle(ctx context.Context, params NodesPutObmsByNodeIDParams) middleware.Responder {
	return fn(ctx, params)
}

// NodesPutObmsByNodeIDHandler interface for that can handle valid nodes put obms by node Id params
type NodesPutObmsByNodeIDHandler interface {
	Handle(context.Context, NodesPutObmsByNodeIDParams) middleware.Responder
}

// NewNodesPutObmsByNodeID creates a new http.Handler for the nodes put obms by node Id operation
func NewNodesPutObmsByNodeID(ctx *middleware.Context, handler NodesPutObmsByNodeIDHandler) *NodesPutObmsByNodeID {
	return &NodesPutObmsByNodeID{Context: ctx, Handler: handler}
}

/*NodesPutObmsByNodeID swagger:route PUT /nodes/{identifier}/obm /api/2.0 nodesPutObmsByNodeId

Put an OBM service

Create or update the specified OBM service with the associated Node ID

*/
type NodesPutObmsByNodeID struct {
	Context *middleware.Context
	Handler NodesPutObmsByNodeIDHandler
}

func (o *NodesPutObmsByNodeID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewNodesPutObmsByNodeIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// NodesPutObmsByNodeIDCreatedBody nodes put obms by node ID created body
// swagger:model NodesPutObmsByNodeIDCreatedBody
type NodesPutObmsByNodeIDCreatedBody interface{}
