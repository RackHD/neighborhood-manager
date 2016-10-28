package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WorkflowsDeleteGraphsByNameHandlerFunc turns a function with the right signature into a workflows delete graphs by name handler
type WorkflowsDeleteGraphsByNameHandlerFunc func(WorkflowsDeleteGraphsByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WorkflowsDeleteGraphsByNameHandlerFunc) Handle(params WorkflowsDeleteGraphsByNameParams) middleware.Responder {
	return fn(params)
}

// WorkflowsDeleteGraphsByNameHandler interface for that can handle valid workflows delete graphs by name params
type WorkflowsDeleteGraphsByNameHandler interface {
	Handle(WorkflowsDeleteGraphsByNameParams) middleware.Responder
}

// NewWorkflowsDeleteGraphsByName creates a new http.Handler for the workflows delete graphs by name operation
func NewWorkflowsDeleteGraphsByName(ctx *middleware.Context, handler WorkflowsDeleteGraphsByNameHandler) *WorkflowsDeleteGraphsByName {
	return &WorkflowsDeleteGraphsByName{Context: ctx, Handler: handler}
}

/*WorkflowsDeleteGraphsByName swagger:route DELETE /workflows/graphs/{injectableName} /api/2.0 workflowsDeleteGraphsByName

Delete a workflow graph

Delete the workflow graph with the specified value of the injectableName property.

*/
type WorkflowsDeleteGraphsByName struct {
	Context *middleware.Context
	Handler WorkflowsDeleteGraphsByNameHandler
}

func (o *WorkflowsDeleteGraphsByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWorkflowsDeleteGraphsByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// WorkflowsDeleteGraphsByNameNoContentBody workflows delete graphs by name no content body
// swagger:model WorkflowsDeleteGraphsByNameNoContentBody
type WorkflowsDeleteGraphsByNameNoContentBody interface{}