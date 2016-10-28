package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRoleHandlerFunc turns a function with the right signature into a get role handler
type GetRoleHandlerFunc func(GetRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRoleHandlerFunc) Handle(params GetRoleParams) middleware.Responder {
	return fn(params)
}

// GetRoleHandler interface for that can handle valid get role params
type GetRoleHandler interface {
	Handle(GetRoleParams) middleware.Responder
}

// NewGetRole creates a new http.Handler for the get role operation
func NewGetRole(ctx *middleware.Context, handler GetRoleHandler) *GetRole {
	return &GetRole{Context: ctx, Handler: handler}
}

/*GetRole swagger:route GET /roles/{name} /api/2.0 getRole

Get a role

Get information about the specified role.

*/
type GetRole struct {
	Context *middleware.Context
	Handler GetRoleHandler
}

func (o *GetRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetRoleUnauthorizedBody get role unauthorized body
// swagger:model GetRoleUnauthorizedBody
type GetRoleUnauthorizedBody interface{}

// GetRoleForbiddenBody get role forbidden body
// swagger:model GetRoleForbiddenBody
type GetRoleForbiddenBody interface{}