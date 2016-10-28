package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostTaskByIDHandlerFunc turns a function with the right signature into a post task by Id handler
type PostTaskByIDHandlerFunc func(PostTaskByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTaskByIDHandlerFunc) Handle(params PostTaskByIDParams) middleware.Responder {
	return fn(params)
}

// PostTaskByIDHandler interface for that can handle valid post task by Id params
type PostTaskByIDHandler interface {
	Handle(PostTaskByIDParams) middleware.Responder
}

// NewPostTaskByID creates a new http.Handler for the post task by Id operation
func NewPostTaskByID(ctx *middleware.Context, handler PostTaskByIDHandler) *PostTaskByID {
	return &PostTaskByID{Context: ctx, Handler: handler}
}

/*PostTaskByID swagger:route POST /tasks/{identifier} /api/2.0 postTaskById

Post a task

Start the specified task

*/
type PostTaskByID struct {
	Context *middleware.Context
	Handler PostTaskByIDHandler
}

func (o *PostTaskByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostTaskByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostTaskByIDCreatedBody post task by ID created body
// swagger:model PostTaskByIDCreatedBody
type PostTaskByIDCreatedBody interface{}