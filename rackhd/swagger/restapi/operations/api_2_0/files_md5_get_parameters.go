package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFilesMd5GetParams creates a new FilesMd5GetParams object
// with the default values initialized.
func NewFilesMd5GetParams() FilesMd5GetParams {
	var ()
	return FilesMd5GetParams{}
}

// FilesMd5GetParams contains all the bound params for the files md5 get operation
// typically these are obtained from a http.Request
//
// swagger:parameters filesMd5Get
type FilesMd5GetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*File name of a file as provided when you originally stored it
	  Required: true
	  In: path
	*/
	Filename string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FilesMd5GetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rFilename, rhkFilename, _ := route.Params.GetOK("filename")
	if err := o.bindFilename(rFilename, rhkFilename, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FilesMd5GetParams) bindFilename(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Filename = raw

	return nil
}
