package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*LookupsPostCreated Successfully created new lookup

swagger:response lookupsPostCreated
*/
type LookupsPostCreated struct {

	// In: body
	Payload LookupsPostCreatedBody `json:"body,omitempty"`
}

// NewLookupsPostCreated creates LookupsPostCreated with default headers values
func NewLookupsPostCreated() *LookupsPostCreated {
	return &LookupsPostCreated{}
}

// WithPayload adds the payload to the lookups post created response
func (o *LookupsPostCreated) WithPayload(payload LookupsPostCreatedBody) *LookupsPostCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the lookups post created response
func (o *LookupsPostCreated) SetPayload(payload LookupsPostCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LookupsPostCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*LookupsPostDefault Unexpected error

swagger:response lookupsPostDefault
*/
type LookupsPostDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewLookupsPostDefault creates LookupsPostDefault with default headers values
func NewLookupsPostDefault(code int) *LookupsPostDefault {
	if code <= 0 {
		code = 500
	}

	return &LookupsPostDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the lookups post default response
func (o *LookupsPostDefault) WithStatusCode(code int) *LookupsPostDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the lookups post default response
func (o *LookupsPostDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the lookups post default response
func (o *LookupsPostDefault) WithPayload(payload *models.Error) *LookupsPostDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the lookups post default response
func (o *LookupsPostDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LookupsPostDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
