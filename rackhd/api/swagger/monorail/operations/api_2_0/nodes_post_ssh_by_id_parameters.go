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

	"github.com/RackHD/neighborhood-manager/rackhd/api/swagger/models"
)

// NewNodesPostSSHByIDParams creates a new NodesPostSSHByIDParams object
// with the default values initialized.
func NewNodesPostSSHByIDParams() NodesPostSSHByIDParams {
	var ()
	return NodesPostSSHByIDParams{}
}

// NodesPostSSHByIDParams contains all the bound params for the nodes post Ssh by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters nodesPostSshById
type NodesPostSSHByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The ssh properties to create
	  Required: true
	  In: body
	*/
	Body *models.SSHSettings
	/*The node identifier
	  Required: true
	  In: path
	*/
	Identifier string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NodesPostSSHByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.SSHSettings
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

	rIdentifier, rhkIdentifier, _ := route.Params.GetOK("identifier")
	if err := o.bindIdentifier(rIdentifier, rhkIdentifier, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NodesPostSSHByIDParams) bindIdentifier(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Identifier = raw

	return nil
}
