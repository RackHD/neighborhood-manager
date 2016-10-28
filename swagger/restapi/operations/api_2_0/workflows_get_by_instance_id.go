package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WorkflowsGetByInstanceIDHandlerFunc turns a function with the right signature into a workflows get by instance Id handler
type WorkflowsGetByInstanceIDHandlerFunc func(WorkflowsGetByInstanceIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WorkflowsGetByInstanceIDHandlerFunc) Handle(params WorkflowsGetByInstanceIDParams) middleware.Responder {
	return fn(params)
}

// WorkflowsGetByInstanceIDHandler interface for that can handle valid workflows get by instance Id params
type WorkflowsGetByInstanceIDHandler interface {
	Handle(WorkflowsGetByInstanceIDParams) middleware.Responder
}

// NewWorkflowsGetByInstanceID creates a new http.Handler for the workflows get by instance Id operation
func NewWorkflowsGetByInstanceID(ctx *middleware.Context, handler WorkflowsGetByInstanceIDHandler) *WorkflowsGetByInstanceID {
	return &WorkflowsGetByInstanceID{Context: ctx, Handler: handler}
}

/*WorkflowsGetByInstanceID swagger:route GET /workflows/{identifier} /api/2.0 workflowsGetByInstanceId

Get a workflow

Get the workflow with the specified instance identifier.

*/
type WorkflowsGetByInstanceID struct {
	Context *middleware.Context
	Handler WorkflowsGetByInstanceIDHandler
}

func (o *WorkflowsGetByInstanceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWorkflowsGetByInstanceIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// WorkflowsGetByInstanceIDOKBody workflows get by instance ID o k body
// swagger:model WorkflowsGetByInstanceIDOKBody
type WorkflowsGetByInstanceIDOKBody interface{}