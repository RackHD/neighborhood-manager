package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*FilesPutCreated Successfully stored file

swagger:response filesPutCreated
*/
type FilesPutCreated struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewFilesPutCreated creates FilesPutCreated with default headers values
func NewFilesPutCreated() *FilesPutCreated {
	return &FilesPutCreated{}
}

// WithPayload adds the payload to the files put created response
func (o *FilesPutCreated) WithPayload(payload string) *FilesPutCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the files put created response
func (o *FilesPutCreated) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FilesPutCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*FilesPutInternalServerError Failure serving file request

swagger:response filesPutInternalServerError
*/
type FilesPutInternalServerError struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewFilesPutInternalServerError creates FilesPutInternalServerError with default headers values
func NewFilesPutInternalServerError() *FilesPutInternalServerError {
	return &FilesPutInternalServerError{}
}

// WithPayload adds the payload to the files put internal server error response
func (o *FilesPutInternalServerError) WithPayload(payload *models.Error) *FilesPutInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the files put internal server error response
func (o *FilesPutInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FilesPutInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FilesPutDefault Unexpected error

swagger:response filesPutDefault
*/
type FilesPutDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewFilesPutDefault creates FilesPutDefault with default headers values
func NewFilesPutDefault(code int) *FilesPutDefault {
	if code <= 0 {
		code = 500
	}

	return &FilesPutDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the files put default response
func (o *FilesPutDefault) WithStatusCode(code int) *FilesPutDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the files put default response
func (o *FilesPutDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the files put default response
func (o *FilesPutDefault) WithPayload(payload *models.Error) *FilesPutDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the files put default response
func (o *FilesPutDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FilesPutDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}