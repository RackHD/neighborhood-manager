package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*GetAllTagsOK Successfully retrieved all tags

swagger:response getAllTagsOK
*/
type GetAllTagsOK struct {

	// In: body
	Payload []interface{} `json:"body,omitempty"`
}

// NewGetAllTagsOK creates GetAllTagsOK with default headers values
func NewGetAllTagsOK() *GetAllTagsOK {
	return &GetAllTagsOK{}
}

// WithPayload adds the payload to the get all tags o k response
func (o *GetAllTagsOK) WithPayload(payload []interface{}) *GetAllTagsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all tags o k response
func (o *GetAllTagsOK) SetPayload(payload []interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllTagsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetAllTagsNotFound Not found

swagger:response getAllTagsNotFound
*/
type GetAllTagsNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllTagsNotFound creates GetAllTagsNotFound with default headers values
func NewGetAllTagsNotFound() *GetAllTagsNotFound {
	return &GetAllTagsNotFound{}
}

// WithPayload adds the payload to the get all tags not found response
func (o *GetAllTagsNotFound) WithPayload(payload *models.Error) *GetAllTagsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all tags not found response
func (o *GetAllTagsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllTagsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAllTagsDefault Unexpected error

swagger:response getAllTagsDefault
*/
type GetAllTagsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllTagsDefault creates GetAllTagsDefault with default headers values
func NewGetAllTagsDefault(code int) *GetAllTagsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAllTagsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get all tags default response
func (o *GetAllTagsDefault) WithStatusCode(code int) *GetAllTagsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get all tags default response
func (o *GetAllTagsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get all tags default response
func (o *GetAllTagsDefault) WithPayload(payload *models.Error) *GetAllTagsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all tags default response
func (o *GetAllTagsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllTagsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
