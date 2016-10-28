package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*LookupsPatchByIDOK Successfully modified the lookup

swagger:response lookupsPatchByIdOK
*/
type LookupsPatchByIDOK struct {

	// In: body
	Payload LookupsPatchByIDOKBody `json:"body,omitempty"`
}

// NewLookupsPatchByIDOK creates LookupsPatchByIDOK with default headers values
func NewLookupsPatchByIDOK() *LookupsPatchByIDOK {
	return &LookupsPatchByIDOK{}
}

// WithPayload adds the payload to the lookups patch by Id o k response
func (o *LookupsPatchByIDOK) WithPayload(payload LookupsPatchByIDOKBody) *LookupsPatchByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the lookups patch by Id o k response
func (o *LookupsPatchByIDOK) SetPayload(payload LookupsPatchByIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LookupsPatchByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*LookupsPatchByIDNotFound The specified lookup was not found

swagger:response lookupsPatchByIdNotFound
*/
type LookupsPatchByIDNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewLookupsPatchByIDNotFound creates LookupsPatchByIDNotFound with default headers values
func NewLookupsPatchByIDNotFound() *LookupsPatchByIDNotFound {
	return &LookupsPatchByIDNotFound{}
}

// WithPayload adds the payload to the lookups patch by Id not found response
func (o *LookupsPatchByIDNotFound) WithPayload(payload *models.Error) *LookupsPatchByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the lookups patch by Id not found response
func (o *LookupsPatchByIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LookupsPatchByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*LookupsPatchByIDDefault Unexpected error

swagger:response lookupsPatchByIdDefault
*/
type LookupsPatchByIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewLookupsPatchByIDDefault creates LookupsPatchByIDDefault with default headers values
func NewLookupsPatchByIDDefault(code int) *LookupsPatchByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &LookupsPatchByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the lookups patch by Id default response
func (o *LookupsPatchByIDDefault) WithStatusCode(code int) *LookupsPatchByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the lookups patch by Id default response
func (o *LookupsPatchByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the lookups patch by Id default response
func (o *LookupsPatchByIDDefault) WithPayload(payload *models.Error) *LookupsPatchByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the lookups patch by Id default response
func (o *LookupsPatchByIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LookupsPatchByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}