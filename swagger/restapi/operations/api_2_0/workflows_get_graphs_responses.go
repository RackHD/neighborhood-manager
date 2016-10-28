package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*WorkflowsGetGraphsOK Successfully retrieved all workflow graphs

swagger:response workflowsGetGraphsOK
*/
type WorkflowsGetGraphsOK struct {

	// In: body
	Payload WorkflowsGetGraphsOKBody `json:"body,omitempty"`
}

// NewWorkflowsGetGraphsOK creates WorkflowsGetGraphsOK with default headers values
func NewWorkflowsGetGraphsOK() *WorkflowsGetGraphsOK {
	return &WorkflowsGetGraphsOK{}
}

// WithPayload adds the payload to the workflows get graphs o k response
func (o *WorkflowsGetGraphsOK) WithPayload(payload WorkflowsGetGraphsOKBody) *WorkflowsGetGraphsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows get graphs o k response
func (o *WorkflowsGetGraphsOK) SetPayload(payload WorkflowsGetGraphsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsGetGraphsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*WorkflowsGetGraphsDefault Unexpected error

swagger:response workflowsGetGraphsDefault
*/
type WorkflowsGetGraphsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewWorkflowsGetGraphsDefault creates WorkflowsGetGraphsDefault with default headers values
func NewWorkflowsGetGraphsDefault(code int) *WorkflowsGetGraphsDefault {
	if code <= 0 {
		code = 500
	}

	return &WorkflowsGetGraphsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the workflows get graphs default response
func (o *WorkflowsGetGraphsDefault) WithStatusCode(code int) *WorkflowsGetGraphsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the workflows get graphs default response
func (o *WorkflowsGetGraphsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the workflows get graphs default response
func (o *WorkflowsGetGraphsDefault) WithPayload(payload *models.Error) *WorkflowsGetGraphsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows get graphs default response
func (o *WorkflowsGetGraphsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsGetGraphsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}