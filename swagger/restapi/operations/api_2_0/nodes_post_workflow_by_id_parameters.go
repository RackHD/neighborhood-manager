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

// NewNodesPostWorkflowByIDParams creates a new NodesPostWorkflowByIDParams object
// with the default values initialized.
func NewNodesPostWorkflowByIDParams() NodesPostWorkflowByIDParams {
	var ()
	return NodesPostWorkflowByIDParams{}
}

// NodesPostWorkflowByIDParams contains all the bound params for the nodes post workflow by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters nodesPostWorkflowById
type NodesPostWorkflowByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The name property set to the injectableName property of the workflow graph
	  In: body
	*/
	Body *models.PostNodeWorkflow
	/*The node identifier
	  Required: true
	  In: path
	*/
	Identifier string
	/*The optional name of the workflow graph to run
	  In: query
	*/
	Name *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NodesPostWorkflowByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	rIdentifier, rhkIdentifier, _ := route.Params.GetOK("identifier")
	if err := o.bindIdentifier(rIdentifier, rhkIdentifier, route.Formats); err != nil {
		res = append(res, err)
	}

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NodesPostWorkflowByIDParams) bindIdentifier(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Identifier = raw

	return nil
}

func (o *NodesPostWorkflowByIDParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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