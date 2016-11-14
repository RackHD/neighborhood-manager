package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/rackhd/api/swagger/models"
)

/*TemplatesLibPutCreated Successfully created or updated the specified template.

swagger:response templatesLibPutCreated
*/
type TemplatesLibPutCreated struct {

	// In: body
	Payload TemplatesLibPutCreatedBody `json:"body,omitempty"`
}

// NewTemplatesLibPutCreated creates TemplatesLibPutCreated with default headers values
func NewTemplatesLibPutCreated() *TemplatesLibPutCreated {
	return &TemplatesLibPutCreated{}
}

// WithPayload adds the payload to the templates lib put created response
func (o *TemplatesLibPutCreated) WithPayload(payload TemplatesLibPutCreatedBody) *TemplatesLibPutCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the templates lib put created response
func (o *TemplatesLibPutCreated) SetPayload(payload TemplatesLibPutCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TemplatesLibPutCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*TemplatesLibPutNotFound The specified template was not found.

swagger:response templatesLibPutNotFound
*/
type TemplatesLibPutNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewTemplatesLibPutNotFound creates TemplatesLibPutNotFound with default headers values
func NewTemplatesLibPutNotFound() *TemplatesLibPutNotFound {
	return &TemplatesLibPutNotFound{}
}

// WithPayload adds the payload to the templates lib put not found response
func (o *TemplatesLibPutNotFound) WithPayload(payload *models.Error) *TemplatesLibPutNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the templates lib put not found response
func (o *TemplatesLibPutNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TemplatesLibPutNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TemplatesLibPutDefault Unexpected error

swagger:response templatesLibPutDefault
*/
type TemplatesLibPutDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewTemplatesLibPutDefault creates TemplatesLibPutDefault with default headers values
func NewTemplatesLibPutDefault(code int) *TemplatesLibPutDefault {
	if code <= 0 {
		code = 500
	}

	return &TemplatesLibPutDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the templates lib put default response
func (o *TemplatesLibPutDefault) WithStatusCode(code int) *TemplatesLibPutDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the templates lib put default response
func (o *TemplatesLibPutDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the templates lib put default response
func (o *TemplatesLibPutDefault) WithPayload(payload *models.Error) *TemplatesLibPutDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the templates lib put default response
func (o *TemplatesLibPutDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TemplatesLibPutDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
