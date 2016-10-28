package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*ObmsGetByIDOK Successfully retrieved the specified OBM service

swagger:response obmsGetByIdOK
*/
type ObmsGetByIDOK struct {

	// In: body
	Payload ObmsGetByIDOKBody `json:"body,omitempty"`
}

// NewObmsGetByIDOK creates ObmsGetByIDOK with default headers values
func NewObmsGetByIDOK() *ObmsGetByIDOK {
	return &ObmsGetByIDOK{}
}

// WithPayload adds the payload to the obms get by Id o k response
func (o *ObmsGetByIDOK) WithPayload(payload ObmsGetByIDOKBody) *ObmsGetByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the obms get by Id o k response
func (o *ObmsGetByIDOK) SetPayload(payload ObmsGetByIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObmsGetByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ObmsGetByIDDefault Unexpected error

swagger:response obmsGetByIdDefault
*/
type ObmsGetByIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewObmsGetByIDDefault creates ObmsGetByIDDefault with default headers values
func NewObmsGetByIDDefault(code int) *ObmsGetByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &ObmsGetByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the obms get by Id default response
func (o *ObmsGetByIDDefault) WithStatusCode(code int) *ObmsGetByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the obms get by Id default response
func (o *ObmsGetByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the obms get by Id default response
func (o *ObmsGetByIDDefault) WithPayload(payload *models.Error) *ObmsGetByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the obms get by Id default response
func (o *ObmsGetByIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObmsGetByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}