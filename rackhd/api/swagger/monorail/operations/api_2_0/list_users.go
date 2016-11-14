package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListUsersHandlerFunc turns a function with the right signature into a list users handler
type ListUsersHandlerFunc func(context.Context, ListUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListUsersHandlerFunc) Handle(ctx context.Context, params ListUsersParams) middleware.Responder {
	return fn(ctx, params)
}

// ListUsersHandler interface for that can handle valid list users params
type ListUsersHandler interface {
	Handle(context.Context, ListUsersParams) middleware.Responder
}

// NewListUsers creates a new http.Handler for the list users operation
func NewListUsers(ctx *middleware.Context, handler ListUsersHandler) *ListUsers {
	return &ListUsers{Context: ctx, Handler: handler}
}

/*ListUsers swagger:route GET /users /api/2.0 listUsers

Get the list of users

Get the list of users currently stored in the system.

*/
type ListUsers struct {
	Context *middleware.Context
	Handler ListUsersHandler
}

func (o *ListUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewListUsersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ListUsersUnauthorizedBody list users unauthorized body
// swagger:model ListUsersUnauthorizedBody
type ListUsersUnauthorizedBody interface{}

// ListUsersForbiddenBody list users forbidden body
// swagger:model ListUsersForbiddenBody
type ListUsersForbiddenBody interface{}
