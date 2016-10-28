package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewWorkflowsGetGraphsParams creates a new WorkflowsGetGraphsParams object
// with the default values initialized.
func NewWorkflowsGetGraphsParams() WorkflowsGetGraphsParams {
	var ()
	return WorkflowsGetGraphsParams{}
}

// WorkflowsGetGraphsParams contains all the bound params for the workflows get graphs operation
// typically these are obtained from a http.Request
//
// swagger:parameters workflowsGetGraphs
type WorkflowsGetGraphsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *WorkflowsGetGraphsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}