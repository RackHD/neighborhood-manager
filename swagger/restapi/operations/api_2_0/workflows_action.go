package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WorkflowsActionHandlerFunc turns a function with the right signature into a workflows action handler
type WorkflowsActionHandlerFunc func(WorkflowsActionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WorkflowsActionHandlerFunc) Handle(params WorkflowsActionParams) middleware.Responder {
	return fn(params)
}

// WorkflowsActionHandler interface for that can handle valid workflows action params
type WorkflowsActionHandler interface {
	Handle(WorkflowsActionParams) middleware.Responder
}

// NewWorkflowsAction creates a new http.Handler for the workflows action operation
func NewWorkflowsAction(ctx *middleware.Context, handler WorkflowsActionHandler) *WorkflowsAction {
	return &WorkflowsAction{Context: ctx, Handler: handler}
}

/*WorkflowsAction swagger:route PUT /workflows/{identifier}/action /api/2.0 workflowsAction

Perform an action on a workflow

Perform the specified action on the workflow with the specified instance identifier. Currently, the cancel action is supported.


*/
type WorkflowsAction struct {
	Context *middleware.Context
	Handler WorkflowsActionHandler
}

func (o *WorkflowsAction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWorkflowsActionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// WorkflowsActionAcceptedBody workflows action accepted body
// swagger:model WorkflowsActionAcceptedBody
type WorkflowsActionAcceptedBody interface{}
