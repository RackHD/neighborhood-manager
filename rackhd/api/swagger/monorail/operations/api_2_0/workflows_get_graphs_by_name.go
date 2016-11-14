package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WorkflowsGetGraphsByNameHandlerFunc turns a function with the right signature into a workflows get graphs by name handler
type WorkflowsGetGraphsByNameHandlerFunc func(context.Context, WorkflowsGetGraphsByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WorkflowsGetGraphsByNameHandlerFunc) Handle(ctx context.Context, params WorkflowsGetGraphsByNameParams) middleware.Responder {
	return fn(ctx, params)
}

// WorkflowsGetGraphsByNameHandler interface for that can handle valid workflows get graphs by name params
type WorkflowsGetGraphsByNameHandler interface {
	Handle(context.Context, WorkflowsGetGraphsByNameParams) middleware.Responder
}

// NewWorkflowsGetGraphsByName creates a new http.Handler for the workflows get graphs by name operation
func NewWorkflowsGetGraphsByName(ctx *middleware.Context, handler WorkflowsGetGraphsByNameHandler) *WorkflowsGetGraphsByName {
	return &WorkflowsGetGraphsByName{Context: ctx, Handler: handler}
}

/*WorkflowsGetGraphsByName swagger:route GET /workflows/graphs/{injectableName} /api/2.0 workflowsGetGraphsByName

Get the specified workflow graph

Get the workflow graph with the specified value of the injectableName property.

*/
type WorkflowsGetGraphsByName struct {
	Context *middleware.Context
	Handler WorkflowsGetGraphsByNameHandler
}

func (o *WorkflowsGetGraphsByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWorkflowsGetGraphsByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// WorkflowsGetGraphsByNameOKBody workflows get graphs by name o k body
// swagger:model WorkflowsGetGraphsByNameOKBody
type WorkflowsGetGraphsByNameOKBody interface{}
