package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*ObmsPutCreated Successfully put the OBM service

swagger:response obmsPutCreated
*/
type ObmsPutCreated struct {

	// In: body
	Payload ObmsPutCreatedBody `json:"body,omitempty"`
}

// NewObmsPutCreated creates ObmsPutCreated with default headers values
func NewObmsPutCreated() *ObmsPutCreated {
	return &ObmsPutCreated{}
}

// WithPayload adds the payload to the obms put created response
func (o *ObmsPutCreated) WithPayload(payload ObmsPutCreatedBody) *ObmsPutCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the obms put created response
func (o *ObmsPutCreated) SetPayload(payload ObmsPutCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObmsPutCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ObmsPutInternalServerError OBM service creation failed

swagger:response obmsPutInternalServerError
*/
type ObmsPutInternalServerError struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewObmsPutInternalServerError creates ObmsPutInternalServerError with default headers values
func NewObmsPutInternalServerError() *ObmsPutInternalServerError {
	return &ObmsPutInternalServerError{}
}

// WithPayload adds the payload to the obms put internal server error response
func (o *ObmsPutInternalServerError) WithPayload(payload *models.Error) *ObmsPutInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the obms put internal server error response
func (o *ObmsPutInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObmsPutInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ObmsPutDefault Unexpected error

swagger:response obmsPutDefault
*/
type ObmsPutDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewObmsPutDefault creates ObmsPutDefault with default headers values
func NewObmsPutDefault(code int) *ObmsPutDefault {
	if code <= 0 {
		code = 500
	}

	return &ObmsPutDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the obms put default response
func (o *ObmsPutDefault) WithStatusCode(code int) *ObmsPutDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the obms put default response
func (o *ObmsPutDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the obms put default response
func (o *ObmsPutDefault) WithPayload(payload *models.Error) *ObmsPutDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the obms put default response
func (o *ObmsPutDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObmsPutDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
