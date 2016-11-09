package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NodesGetCatalogByIDHandlerFunc turns a function with the right signature into a nodes get catalog by Id handler
type NodesGetCatalogByIDHandlerFunc func(NodesGetCatalogByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NodesGetCatalogByIDHandlerFunc) Handle(params NodesGetCatalogByIDParams) middleware.Responder {
	return fn(params)
}

// NodesGetCatalogByIDHandler interface for that can handle valid nodes get catalog by Id params
type NodesGetCatalogByIDHandler interface {
	Handle(NodesGetCatalogByIDParams) middleware.Responder
}

// NewNodesGetCatalogByID creates a new http.Handler for the nodes get catalog by Id operation
func NewNodesGetCatalogByID(ctx *middleware.Context, handler NodesGetCatalogByIDHandler) *NodesGetCatalogByID {
	return &NodesGetCatalogByID{Context: ctx, Handler: handler}
}

/*NodesGetCatalogByID swagger:route GET /nodes/{identifier}/catalogs /api/2.0 nodesGetCatalogById

Get the catalogs for a node

Get a list of all of the catalogs for the specified node.

*/
type NodesGetCatalogByID struct {
	Context *middleware.Context
	Handler NodesGetCatalogByIDHandler
}

func (o *NodesGetCatalogByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewNodesGetCatalogByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// NodesGetCatalogByIDOKBody nodes get catalog by ID o k body
// swagger:model NodesGetCatalogByIDOKBody
type NodesGetCatalogByIDOKBody interface{}
