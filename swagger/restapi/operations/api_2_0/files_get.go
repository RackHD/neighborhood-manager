package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FilesGetHandlerFunc turns a function with the right signature into a files get handler
type FilesGetHandlerFunc func(FilesGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FilesGetHandlerFunc) Handle(params FilesGetParams) middleware.Responder {
	return fn(params)
}

// FilesGetHandler interface for that can handle valid files get params
type FilesGetHandler interface {
	Handle(FilesGetParams) middleware.Responder
}

// NewFilesGet creates a new http.Handler for the files get operation
func NewFilesGet(ctx *middleware.Context, handler FilesGetHandler) *FilesGet {
	return &FilesGet{Context: ctx, Handler: handler}
}

/*FilesGet swagger:route GET /files/{fileidentifier} /api/2.0 filesGet

Get a file

Get file based on uuid or file name.

*/
type FilesGet struct {
	Context *middleware.Context
	Handler FilesGetHandler
}

func (o *FilesGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewFilesGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}