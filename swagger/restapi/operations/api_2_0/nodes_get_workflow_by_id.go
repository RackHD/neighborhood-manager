package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NodesGetWorkflowByIDHandlerFunc turns a function with the right signature into a nodes get workflow by Id handler
type NodesGetWorkflowByIDHandlerFunc func(NodesGetWorkflowByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NodesGetWorkflowByIDHandlerFunc) Handle(params NodesGetWorkflowByIDParams) middleware.Responder {
	return fn(params)
}

// NodesGetWorkflowByIDHandler interface for that can handle valid nodes get workflow by Id params
type NodesGetWorkflowByIDHandler interface {
	Handle(NodesGetWorkflowByIDParams) middleware.Responder
}

// NewNodesGetWorkflowByID creates a new http.Handler for the nodes get workflow by Id operation
func NewNodesGetWorkflowByID(ctx *middleware.Context, handler NodesGetWorkflowByIDHandler) *NodesGetWorkflowByID {
	return &NodesGetWorkflowByID{Context: ctx, Handler: handler}
}

/*NodesGetWorkflowByID swagger:route GET /nodes/{identifier}/workflows /api/2.0 nodesGetWorkflowById

Get all workflows for a node

Get a list of all workflows that have run against the specified node, or are currently running against the node.


*/
type NodesGetWorkflowByID struct {
	Context *middleware.Context
	Handler NodesGetWorkflowByIDHandler
}

func (o *NodesGetWorkflowByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewNodesGetWorkflowByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// NodesGetWorkflowByIDOKBody nodes get workflow by ID o k body
// swagger:model NodesGetWorkflowByIDOKBody
type NodesGetWorkflowByIDOKBody interface{}