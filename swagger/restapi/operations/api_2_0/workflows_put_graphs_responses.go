package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*WorkflowsPutGraphsCreated Successfully updated workflow graph

swagger:response workflowsPutGraphsCreated
*/
type WorkflowsPutGraphsCreated struct {

	// In: body
	Payload WorkflowsPutGraphsCreatedBody `json:"body,omitempty"`
}

// NewWorkflowsPutGraphsCreated creates WorkflowsPutGraphsCreated with default headers values
func NewWorkflowsPutGraphsCreated() *WorkflowsPutGraphsCreated {
	return &WorkflowsPutGraphsCreated{}
}

// WithPayload adds the payload to the workflows put graphs created response
func (o *WorkflowsPutGraphsCreated) WithPayload(payload WorkflowsPutGraphsCreatedBody) *WorkflowsPutGraphsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows put graphs created response
func (o *WorkflowsPutGraphsCreated) SetPayload(payload WorkflowsPutGraphsCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsPutGraphsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*WorkflowsPutGraphsInternalServerError Workflow graph was not updated

swagger:response workflowsPutGraphsInternalServerError
*/
type WorkflowsPutGraphsInternalServerError struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewWorkflowsPutGraphsInternalServerError creates WorkflowsPutGraphsInternalServerError with default headers values
func NewWorkflowsPutGraphsInternalServerError() *WorkflowsPutGraphsInternalServerError {
	return &WorkflowsPutGraphsInternalServerError{}
}

// WithPayload adds the payload to the workflows put graphs internal server error response
func (o *WorkflowsPutGraphsInternalServerError) WithPayload(payload *models.Error) *WorkflowsPutGraphsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows put graphs internal server error response
func (o *WorkflowsPutGraphsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsPutGraphsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*WorkflowsPutGraphsDefault Upload failed

swagger:response workflowsPutGraphsDefault
*/
type WorkflowsPutGraphsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewWorkflowsPutGraphsDefault creates WorkflowsPutGraphsDefault with default headers values
func NewWorkflowsPutGraphsDefault(code int) *WorkflowsPutGraphsDefault {
	if code <= 0 {
		code = 500
	}

	return &WorkflowsPutGraphsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the workflows put graphs default response
func (o *WorkflowsPutGraphsDefault) WithStatusCode(code int) *WorkflowsPutGraphsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the workflows put graphs default response
func (o *WorkflowsPutGraphsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the workflows put graphs default response
func (o *WorkflowsPutGraphsDefault) WithPayload(payload *models.Error) *WorkflowsPutGraphsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows put graphs default response
func (o *WorkflowsPutGraphsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsPutGraphsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
