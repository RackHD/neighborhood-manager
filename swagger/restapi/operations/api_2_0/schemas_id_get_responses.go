package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*SchemasIDGetOK Successfully retrieved the schema

swagger:response schemasIdGetOK
*/
type SchemasIDGetOK struct {

	// In: body
	Payload SchemasIDGetOKBody `json:"body,omitempty"`
}

// NewSchemasIDGetOK creates SchemasIDGetOK with default headers values
func NewSchemasIDGetOK() *SchemasIDGetOK {
	return &SchemasIDGetOK{}
}

// WithPayload adds the payload to the schemas Id get o k response
func (o *SchemasIDGetOK) WithPayload(payload SchemasIDGetOKBody) *SchemasIDGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schemas Id get o k response
func (o *SchemasIDGetOK) SetPayload(payload SchemasIDGetOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemasIDGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*SchemasIDGetDefault Error

swagger:response schemasIdGetDefault
*/
type SchemasIDGetDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemasIDGetDefault creates SchemasIDGetDefault with default headers values
func NewSchemasIDGetDefault(code int) *SchemasIDGetDefault {
	if code <= 0 {
		code = 500
	}

	return &SchemasIDGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the schemas Id get default response
func (o *SchemasIDGetDefault) WithStatusCode(code int) *SchemasIDGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the schemas Id get default response
func (o *SchemasIDGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the schemas Id get default response
func (o *SchemasIDGetDefault) WithPayload(payload *models.ErrorResponse) *SchemasIDGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schemas Id get default response
func (o *SchemasIDGetDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemasIDGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
