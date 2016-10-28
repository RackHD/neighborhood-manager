package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WorkflowsPutTaskHandlerFunc turns a function with the right signature into a workflows put task handler
type WorkflowsPutTaskHandlerFunc func(WorkflowsPutTaskParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WorkflowsPutTaskHandlerFunc) Handle(params WorkflowsPutTaskParams) middleware.Responder {
	return fn(params)
}

// WorkflowsPutTaskHandler interface for that can handle valid workflows put task params
type WorkflowsPutTaskHandler interface {
	Handle(WorkflowsPutTaskParams) middleware.Responder
}

// NewWorkflowsPutTask creates a new http.Handler for the workflows put task operation
func NewWorkflowsPutTask(ctx *middleware.Context, handler WorkflowsPutTaskHandler) *WorkflowsPutTask {
	return &WorkflowsPutTask{Context: ctx, Handler: handler}
}

/*WorkflowsPutTask swagger:route PUT /workflows/tasks /api/2.0 workflowsPutTask

Put a workflow task

Create or update a workflow task in the library of tasks.

*/
type WorkflowsPutTask struct {
	Context *middleware.Context
	Handler WorkflowsPutTaskHandler
}

func (o *WorkflowsPutTask) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWorkflowsPutTaskParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// WorkflowsPutTaskCreatedBody workflows put task created body
// swagger:model WorkflowsPutTaskCreatedBody
type WorkflowsPutTaskCreatedBody interface{}