package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/rackhd/api/swagger/models"
)

/*WorkflowsGetTasksByNameOK Successfully retrieved the workflow task with the specified injectable name

swagger:response workflowsGetTasksByNameOK
*/
type WorkflowsGetTasksByNameOK struct {

	// In: body
	Payload WorkflowsGetTasksByNameOKBody `json:"body,omitempty"`
}

// NewWorkflowsGetTasksByNameOK creates WorkflowsGetTasksByNameOK with default headers values
func NewWorkflowsGetTasksByNameOK() *WorkflowsGetTasksByNameOK {
	return &WorkflowsGetTasksByNameOK{}
}

// WithPayload adds the payload to the workflows get tasks by name o k response
func (o *WorkflowsGetTasksByNameOK) WithPayload(payload WorkflowsGetTasksByNameOKBody) *WorkflowsGetTasksByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows get tasks by name o k response
func (o *WorkflowsGetTasksByNameOK) SetPayload(payload WorkflowsGetTasksByNameOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsGetTasksByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*WorkflowsGetTasksByNameDefault Unexpected error

swagger:response workflowsGetTasksByNameDefault
*/
type WorkflowsGetTasksByNameDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewWorkflowsGetTasksByNameDefault creates WorkflowsGetTasksByNameDefault with default headers values
func NewWorkflowsGetTasksByNameDefault(code int) *WorkflowsGetTasksByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &WorkflowsGetTasksByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the workflows get tasks by name default response
func (o *WorkflowsGetTasksByNameDefault) WithStatusCode(code int) *WorkflowsGetTasksByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the workflows get tasks by name default response
func (o *WorkflowsGetTasksByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the workflows get tasks by name default response
func (o *WorkflowsGetTasksByNameDefault) WithPayload(payload *models.Error) *WorkflowsGetTasksByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows get tasks by name default response
func (o *WorkflowsGetTasksByNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsGetTasksByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
