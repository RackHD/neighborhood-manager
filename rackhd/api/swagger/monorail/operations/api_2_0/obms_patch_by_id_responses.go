package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/rackhd/api/swagger/models"
)

/*ObmsPatchByIDOK Successfully patched the specified OBM settings

swagger:response obmsPatchByIdOK
*/
type ObmsPatchByIDOK struct {

	// In: body
	Payload ObmsPatchByIDOKBody `json:"body,omitempty"`
}

// NewObmsPatchByIDOK creates ObmsPatchByIDOK with default headers values
func NewObmsPatchByIDOK() *ObmsPatchByIDOK {
	return &ObmsPatchByIDOK{}
}

// WithPayload adds the payload to the obms patch by Id o k response
func (o *ObmsPatchByIDOK) WithPayload(payload ObmsPatchByIDOKBody) *ObmsPatchByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the obms patch by Id o k response
func (o *ObmsPatchByIDOK) SetPayload(payload ObmsPatchByIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObmsPatchByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ObmsPatchByIDInternalServerError OBM patch failed

swagger:response obmsPatchByIdInternalServerError
*/
type ObmsPatchByIDInternalServerError struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewObmsPatchByIDInternalServerError creates ObmsPatchByIDInternalServerError with default headers values
func NewObmsPatchByIDInternalServerError() *ObmsPatchByIDInternalServerError {
	return &ObmsPatchByIDInternalServerError{}
}

// WithPayload adds the payload to the obms patch by Id internal server error response
func (o *ObmsPatchByIDInternalServerError) WithPayload(payload *models.Error) *ObmsPatchByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the obms patch by Id internal server error response
func (o *ObmsPatchByIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObmsPatchByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ObmsPatchByIDDefault Unexpected error

swagger:response obmsPatchByIdDefault
*/
type ObmsPatchByIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewObmsPatchByIDDefault creates ObmsPatchByIDDefault with default headers values
func NewObmsPatchByIDDefault(code int) *ObmsPatchByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &ObmsPatchByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the obms patch by Id default response
func (o *ObmsPatchByIDDefault) WithStatusCode(code int) *ObmsPatchByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the obms patch by Id default response
func (o *ObmsPatchByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the obms patch by Id default response
func (o *ObmsPatchByIDDefault) WithPayload(payload *models.Error) *ObmsPatchByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the obms patch by Id default response
func (o *ObmsPatchByIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObmsPatchByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
