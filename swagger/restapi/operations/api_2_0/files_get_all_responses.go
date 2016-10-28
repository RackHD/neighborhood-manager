package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*FilesGetAllOK Successfully retrieved the files

swagger:response filesGetAllOK
*/
type FilesGetAllOK struct {

	// In: body
	Payload FilesGetAllOKBody `json:"body,omitempty"`
}

// NewFilesGetAllOK creates FilesGetAllOK with default headers values
func NewFilesGetAllOK() *FilesGetAllOK {
	return &FilesGetAllOK{}
}

// WithPayload adds the payload to the files get all o k response
func (o *FilesGetAllOK) WithPayload(payload FilesGetAllOKBody) *FilesGetAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the files get all o k response
func (o *FilesGetAllOK) SetPayload(payload FilesGetAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FilesGetAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*FilesGetAllInternalServerError Failed to serve file request

swagger:response filesGetAllInternalServerError
*/
type FilesGetAllInternalServerError struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewFilesGetAllInternalServerError creates FilesGetAllInternalServerError with default headers values
func NewFilesGetAllInternalServerError() *FilesGetAllInternalServerError {
	return &FilesGetAllInternalServerError{}
}

// WithPayload adds the payload to the files get all internal server error response
func (o *FilesGetAllInternalServerError) WithPayload(payload *models.Error) *FilesGetAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the files get all internal server error response
func (o *FilesGetAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FilesGetAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FilesGetAllDefault Unexpected error

swagger:response filesGetAllDefault
*/
type FilesGetAllDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewFilesGetAllDefault creates FilesGetAllDefault with default headers values
func NewFilesGetAllDefault(code int) *FilesGetAllDefault {
	if code <= 0 {
		code = 500
	}

	return &FilesGetAllDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the files get all default response
func (o *FilesGetAllDefault) WithStatusCode(code int) *FilesGetAllDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the files get all default response
func (o *FilesGetAllDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the files get all default response
func (o *FilesGetAllDefault) WithPayload(payload *models.Error) *FilesGetAllDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the files get all default response
func (o *FilesGetAllDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FilesGetAllDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}