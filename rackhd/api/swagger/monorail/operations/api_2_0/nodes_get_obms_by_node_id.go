package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NodesGetObmsByNodeIDHandlerFunc turns a function with the right signature into a nodes get obms by node Id handler
type NodesGetObmsByNodeIDHandlerFunc func(context.Context, NodesGetObmsByNodeIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NodesGetObmsByNodeIDHandlerFunc) Handle(ctx context.Context, params NodesGetObmsByNodeIDParams) middleware.Responder {
	return fn(ctx, params)
}

// NodesGetObmsByNodeIDHandler interface for that can handle valid nodes get obms by node Id params
type NodesGetObmsByNodeIDHandler interface {
	Handle(context.Context, NodesGetObmsByNodeIDParams) middleware.Responder
}

// NewNodesGetObmsByNodeID creates a new http.Handler for the nodes get obms by node Id operation
func NewNodesGetObmsByNodeID(ctx *middleware.Context, handler NodesGetObmsByNodeIDHandler) *NodesGetObmsByNodeID {
	return &NodesGetObmsByNodeID{Context: ctx, Handler: handler}
}

/*NodesGetObmsByNodeID swagger:route GET /nodes/{identifier}/obm /api/2.0 nodesGetObmsByNodeId

Get all OBM services

Get all the OBM settings for the specified node.

*/
type NodesGetObmsByNodeID struct {
	Context *middleware.Context
	Handler NodesGetObmsByNodeIDHandler
}

func (o *NodesGetObmsByNodeID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewNodesGetObmsByNodeIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// NodesGetObmsByNodeIDOKBody nodes get obms by node ID o k body
// swagger:model NodesGetObmsByNodeIDOKBody
type NodesGetObmsByNodeIDOKBody interface{}
