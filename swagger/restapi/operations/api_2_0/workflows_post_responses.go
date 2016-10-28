package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*WorkflowsPostCreated Successfully posted the workflow

swagger:response workflowsPostCreated
*/
type WorkflowsPostCreated struct {

	// In: body
	Payload WorkflowsPostCreatedBody `json:"body,omitempty"`
}

// NewWorkflowsPostCreated creates WorkflowsPostCreated with default headers values
func NewWorkflowsPostCreated() *WorkflowsPostCreated {
	return &WorkflowsPostCreated{}
}

// WithPayload adds the payload to the workflows post created response
func (o *WorkflowsPostCreated) WithPayload(payload WorkflowsPostCreatedBody) *WorkflowsPostCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows post created response
func (o *WorkflowsPostCreated) SetPayload(payload WorkflowsPostCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsPostCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*WorkflowsPostInternalServerError Workflow was not run

swagger:response workflowsPostInternalServerError
*/
type WorkflowsPostInternalServerError struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewWorkflowsPostInternalServerError creates WorkflowsPostInternalServerError with default headers values
func NewWorkflowsPostInternalServerError() *WorkflowsPostInternalServerError {
	return &WorkflowsPostInternalServerError{}
}

// WithPayload adds the payload to the workflows post internal server error response
func (o *WorkflowsPostInternalServerError) WithPayload(payload *models.Error) *WorkflowsPostInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows post internal server error response
func (o *WorkflowsPostInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsPostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*WorkflowsPostDefault Upload failed

swagger:response workflowsPostDefault
*/
type WorkflowsPostDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewWorkflowsPostDefault creates WorkflowsPostDefault with default headers values
func NewWorkflowsPostDefault(code int) *WorkflowsPostDefault {
	if code <= 0 {
		code = 500
	}

	return &WorkflowsPostDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the workflows post default response
func (o *WorkflowsPostDefault) WithStatusCode(code int) *WorkflowsPostDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the workflows post default response
func (o *WorkflowsPostDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the workflows post default response
func (o *WorkflowsPostDefault) WithPayload(payload *models.Error) *WorkflowsPostDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the workflows post default response
func (o *WorkflowsPostDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WorkflowsPostDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}