package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*GetTagOK Successfully retrieved information about the specified tag

swagger:response getTagOK
*/
type GetTagOK struct {

	// In: body
	Payload GetTagOKBody `json:"body,omitempty"`
}

// NewGetTagOK creates GetTagOK with default headers values
func NewGetTagOK() *GetTagOK {
	return &GetTagOK{}
}

// WithPayload adds the payload to the get tag o k response
func (o *GetTagOK) WithPayload(payload GetTagOKBody) *GetTagOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tag o k response
func (o *GetTagOK) SetPayload(payload GetTagOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetTagNotFound The tag name identifier was not found

swagger:response getTagNotFound
*/
type GetTagNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagNotFound creates GetTagNotFound with default headers values
func NewGetTagNotFound() *GetTagNotFound {
	return &GetTagNotFound{}
}

// WithPayload adds the payload to the get tag not found response
func (o *GetTagNotFound) WithPayload(payload *models.Error) *GetTagNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tag not found response
func (o *GetTagNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetTagDefault Unexpected error

swagger:response getTagDefault
*/
type GetTagDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTagDefault creates GetTagDefault with default headers values
func NewGetTagDefault(code int) *GetTagDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTagDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get tag default response
func (o *GetTagDefault) WithStatusCode(code int) *GetTagDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get tag default response
func (o *GetTagDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get tag default response
func (o *GetTagDefault) WithPayload(payload *models.Error) *GetTagDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tag default response
func (o *GetTagDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTagDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}