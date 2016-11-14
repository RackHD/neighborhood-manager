package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/rackhd/api/swagger/models"
)

/*ViewsDeleteNoContent Successfully deleted the specified view

swagger:response viewsDeleteNoContent
*/
type ViewsDeleteNoContent struct {

	// In: body
	Payload ViewsDeleteNoContentBody `json:"body,omitempty"`
}

// NewViewsDeleteNoContent creates ViewsDeleteNoContent with default headers values
func NewViewsDeleteNoContent() *ViewsDeleteNoContent {
	return &ViewsDeleteNoContent{}
}

// WithPayload adds the payload to the views delete no content response
func (o *ViewsDeleteNoContent) WithPayload(payload ViewsDeleteNoContentBody) *ViewsDeleteNoContent {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the views delete no content response
func (o *ViewsDeleteNoContent) SetPayload(payload ViewsDeleteNoContentBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ViewsDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ViewsDeleteNotFound The view with specified name was not found

swagger:response viewsDeleteNotFound
*/
type ViewsDeleteNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewViewsDeleteNotFound creates ViewsDeleteNotFound with default headers values
func NewViewsDeleteNotFound() *ViewsDeleteNotFound {
	return &ViewsDeleteNotFound{}
}

// WithPayload adds the payload to the views delete not found response
func (o *ViewsDeleteNotFound) WithPayload(payload *models.Error) *ViewsDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the views delete not found response
func (o *ViewsDeleteNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ViewsDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ViewsDeleteDefault Unexpected error

swagger:response viewsDeleteDefault
*/
type ViewsDeleteDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewViewsDeleteDefault creates ViewsDeleteDefault with default headers values
func NewViewsDeleteDefault(code int) *ViewsDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &ViewsDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the views delete default response
func (o *ViewsDeleteDefault) WithStatusCode(code int) *ViewsDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the views delete default response
func (o *ViewsDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the views delete default response
func (o *ViewsDeleteDefault) WithPayload(payload *models.Error) *ViewsDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the views delete default response
func (o *ViewsDeleteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ViewsDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
