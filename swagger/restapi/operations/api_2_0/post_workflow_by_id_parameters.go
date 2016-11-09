package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

// NewPostWorkflowByIDParams creates a new PostWorkflowByIDParams object
// with the default values initialized.
func NewPostWorkflowByIDParams() PostWorkflowByIDParams {
	var ()
	return PostWorkflowByIDParams{}
}

// PostWorkflowByIDParams contains all the bound params for the post workflow by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters postWorkflowById
type PostWorkflowByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The workflow options to post
	  In: body
	*/
	Body *models.PostNodeWorkflow
	/*Query string specifying the optional workflow injectable name
	  In: query
	*/
	Name *string
	/*The tag identifier
	  Required: true
	  In: path
	*/
	TagName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PostWorkflowByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PostNodeWorkflow
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	}

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	rTagName, rhkTagName, _ := route.Params.GetOK("tagName")
	if err := o.bindTagName(rTagName, rhkTagName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWorkflowByIDParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Name = &raw

	return nil
}

func (o *PostWorkflowByIDParams) bindTagName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.TagName = raw

	return nil
}
