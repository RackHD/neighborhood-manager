package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteTagHandlerFunc turns a function with the right signature into a delete tag handler
type DeleteTagHandlerFunc func(DeleteTagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTagHandlerFunc) Handle(params DeleteTagParams) middleware.Responder {
	return fn(params)
}

// DeleteTagHandler interface for that can handle valid delete tag params
type DeleteTagHandler interface {
	Handle(DeleteTagParams) middleware.Responder
}

// NewDeleteTag creates a new http.Handler for the delete tag operation
func NewDeleteTag(ctx *middleware.Context, handler DeleteTagHandler) *DeleteTag {
	return &DeleteTag{Context: ctx, Handler: handler}
}

/*DeleteTag swagger:route DELETE /tags/{tagName} /api/2.0 deleteTag

Delete a tag

Delete the specified tag.

*/
type DeleteTag struct {
	Context *middleware.Context
	Handler DeleteTagHandler
}

func (o *DeleteTag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteTagParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteTagNoContentBody delete tag no content body
// swagger:model DeleteTagNoContentBody
type DeleteTagNoContentBody interface{}
