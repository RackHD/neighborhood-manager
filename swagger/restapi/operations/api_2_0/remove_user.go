package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RemoveUserHandlerFunc turns a function with the right signature into a remove user handler
type RemoveUserHandlerFunc func(RemoveUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveUserHandlerFunc) Handle(params RemoveUserParams) middleware.Responder {
	return fn(params)
}

// RemoveUserHandler interface for that can handle valid remove user params
type RemoveUserHandler interface {
	Handle(RemoveUserParams) middleware.Responder
}

// NewRemoveUser creates a new http.Handler for the remove user operation
func NewRemoveUser(ctx *middleware.Context, handler RemoveUserHandler) *RemoveUser {
	return &RemoveUser{Context: ctx, Handler: handler}
}

/*RemoveUser swagger:route DELETE /users/{name} /api/2.0 removeUser

Delete a user

Delete the specified user.

*/
type RemoveUser struct {
	Context *middleware.Context
	Handler RemoveUserHandler
}

func (o *RemoveUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewRemoveUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RemoveUserUnauthorizedBody remove user unauthorized body
// swagger:model RemoveUserUnauthorizedBody
type RemoveUserUnauthorizedBody interface{}

// RemoveUserForbiddenBody remove user forbidden body
// swagger:model RemoveUserForbiddenBody
type RemoveUserForbiddenBody interface{}