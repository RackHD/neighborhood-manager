package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

// NewLookupsPatchByIDParams creates a new LookupsPatchByIDParams object
// with the default values initialized.
func NewLookupsPatchByIDParams() LookupsPatchByIDParams {
	var ()
	return LookupsPatchByIDParams{}
}

// LookupsPatchByIDParams contains all the bound params for the lookups patch by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters lookupsPatchById
type LookupsPatchByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The lookup properties to be modified
	  Required: true
	  In: body
	*/
	Body *models.Lookups20LookupBase
	/*The id of the lookup to patch
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *LookupsPatchByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Lookups20LookupBase
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body"))
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LookupsPatchByIDParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}