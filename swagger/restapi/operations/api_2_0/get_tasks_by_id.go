package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTasksByIDHandlerFunc turns a function with the right signature into a get tasks by Id handler
type GetTasksByIDHandlerFunc func(GetTasksByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTasksByIDHandlerFunc) Handle(params GetTasksByIDParams) middleware.Responder {
	return fn(params)
}

// GetTasksByIDHandler interface for that can handle valid get tasks by Id params
type GetTasksByIDHandler interface {
	Handle(GetTasksByIDParams) middleware.Responder
}

// NewGetTasksByID creates a new http.Handler for the get tasks by Id operation
func NewGetTasksByID(ctx *middleware.Context, handler GetTasksByIDHandler) *GetTasksByID {
	return &GetTasksByID{Context: ctx, Handler: handler}
}

/*GetTasksByID swagger:route GET /tasks/{identifier} /api/2.0 getTasksById

Get a task

Get the specified task.

*/
type GetTasksByID struct {
	Context *middleware.Context
	Handler GetTasksByIDHandler
}

func (o *GetTasksByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetTasksByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetTasksByIDOKBody get tasks by ID o k body
// swagger:model GetTasksByIDOKBody
type GetTasksByIDOKBody interface{}