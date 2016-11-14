package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SkusIDDeletePackHandlerFunc turns a function with the right signature into a skus Id delete pack handler
type SkusIDDeletePackHandlerFunc func(context.Context, SkusIDDeletePackParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SkusIDDeletePackHandlerFunc) Handle(ctx context.Context, params SkusIDDeletePackParams) middleware.Responder {
	return fn(ctx, params)
}

// SkusIDDeletePackHandler interface for that can handle valid skus Id delete pack params
type SkusIDDeletePackHandler interface {
	Handle(context.Context, SkusIDDeletePackParams) middleware.Responder
}

// NewSkusIDDeletePack creates a new http.Handler for the skus Id delete pack operation
func NewSkusIDDeletePack(ctx *middleware.Context, handler SkusIDDeletePackHandler) *SkusIDDeletePack {
	return &SkusIDDeletePack{Context: ctx, Handler: handler}
}

/*SkusIDDeletePack swagger:route DELETE /skus/{identifier}/pack /api/2.0 skusIdDeletePack

Delete a SKU Pack

Delete a SKU Pack associated with the specified SKU.

*/
type SkusIDDeletePack struct {
	Context *middleware.Context
	Handler SkusIDDeletePackHandler
}

func (o *SkusIDDeletePack) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewSkusIDDeletePackParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SkusIDDeletePackNoContentBody skus ID delete pack no content body
// swagger:model SkusIDDeletePackNoContentBody
type SkusIDDeletePackNoContentBody interface{}
