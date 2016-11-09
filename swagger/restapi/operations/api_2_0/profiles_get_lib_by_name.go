package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ProfilesGetLibByNameHandlerFunc turns a function with the right signature into a profiles get lib by name handler
type ProfilesGetLibByNameHandlerFunc func(ProfilesGetLibByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ProfilesGetLibByNameHandlerFunc) Handle(params ProfilesGetLibByNameParams) middleware.Responder {
	return fn(params)
}

// ProfilesGetLibByNameHandler interface for that can handle valid profiles get lib by name params
type ProfilesGetLibByNameHandler interface {
	Handle(ProfilesGetLibByNameParams) middleware.Responder
}

// NewProfilesGetLibByName creates a new http.Handler for the profiles get lib by name operation
func NewProfilesGetLibByName(ctx *middleware.Context, handler ProfilesGetLibByNameHandler) *ProfilesGetLibByName {
	return &ProfilesGetLibByName{Context: ctx, Handler: handler}
}

/*ProfilesGetLibByName swagger:route GET /profiles/library/{name} /api/2.0 profilesGetLibByName

Get contents of a profile

Get the contents of a profile specified by its name.

*/
type ProfilesGetLibByName struct {
	Context *middleware.Context
	Handler ProfilesGetLibByNameHandler
}

func (o *ProfilesGetLibByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewProfilesGetLibByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ProfilesGetLibByNameOKBody profiles get lib by name o k body
// swagger:model ProfilesGetLibByNameOKBody
type ProfilesGetLibByNameOKBody interface{}
