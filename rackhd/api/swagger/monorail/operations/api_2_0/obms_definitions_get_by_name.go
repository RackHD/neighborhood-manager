package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ObmsDefinitionsGetByNameHandlerFunc turns a function with the right signature into a obms definitions get by name handler
type ObmsDefinitionsGetByNameHandlerFunc func(context.Context, ObmsDefinitionsGetByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ObmsDefinitionsGetByNameHandlerFunc) Handle(ctx context.Context, params ObmsDefinitionsGetByNameParams) middleware.Responder {
	return fn(ctx, params)
}

// ObmsDefinitionsGetByNameHandler interface for that can handle valid obms definitions get by name params
type ObmsDefinitionsGetByNameHandler interface {
	Handle(context.Context, ObmsDefinitionsGetByNameParams) middleware.Responder
}

// NewObmsDefinitionsGetByName creates a new http.Handler for the obms definitions get by name operation
func NewObmsDefinitionsGetByName(ctx *middleware.Context, handler ObmsDefinitionsGetByNameHandler) *ObmsDefinitionsGetByName {
	return &ObmsDefinitionsGetByName{Context: ctx, Handler: handler}
}

/*ObmsDefinitionsGetByName swagger:route GET /obms/definitions/{name} /api/2.0 obmsDefinitionsGetByName

Get an OBM service definition

Get the contents of the specified OBM service schema.

*/
type ObmsDefinitionsGetByName struct {
	Context *middleware.Context
	Handler ObmsDefinitionsGetByNameHandler
}

func (o *ObmsDefinitionsGetByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewObmsDefinitionsGetByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ObmsDefinitionsGetByNameOKBody obms definitions get by name o k body
// swagger:model ObmsDefinitionsGetByNameOKBody
type ObmsDefinitionsGetByNameOKBody interface{}
