package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/rackhd/api/swagger/models"
)

/*ObmsDefinitionsGetByNameOK Successfully retrieved the specified OBM schema

swagger:response obmsDefinitionsGetByNameOK
*/
type ObmsDefinitionsGetByNameOK struct {

	// In: body
	Payload ObmsDefinitionsGetByNameOKBody `json:"body,omitempty"`
}

// NewObmsDefinitionsGetByNameOK creates ObmsDefinitionsGetByNameOK with default headers values
func NewObmsDefinitionsGetByNameOK() *ObmsDefinitionsGetByNameOK {
	return &ObmsDefinitionsGetByNameOK{}
}

// WithPayload adds the payload to the obms definitions get by name o k response
func (o *ObmsDefinitionsGetByNameOK) WithPayload(payload ObmsDefinitionsGetByNameOKBody) *ObmsDefinitionsGetByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the obms definitions get by name o k response
func (o *ObmsDefinitionsGetByNameOK) SetPayload(payload ObmsDefinitionsGetByNameOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObmsDefinitionsGetByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ObmsDefinitionsGetByNameDefault Unexpected error

swagger:response obmsDefinitionsGetByNameDefault
*/
type ObmsDefinitionsGetByNameDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewObmsDefinitionsGetByNameDefault creates ObmsDefinitionsGetByNameDefault with default headers values
func NewObmsDefinitionsGetByNameDefault(code int) *ObmsDefinitionsGetByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &ObmsDefinitionsGetByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the obms definitions get by name default response
func (o *ObmsDefinitionsGetByNameDefault) WithStatusCode(code int) *ObmsDefinitionsGetByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the obms definitions get by name default response
func (o *ObmsDefinitionsGetByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the obms definitions get by name default response
func (o *ObmsDefinitionsGetByNameDefault) WithPayload(payload *models.Error) *ObmsDefinitionsGetByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the obms definitions get by name default response
func (o *ObmsDefinitionsGetByNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObmsDefinitionsGetByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
