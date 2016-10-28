package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*PollersDataGetOK Successfully retrieved poller data

swagger:response pollersDataGetOK
*/
type PollersDataGetOK struct {

	// In: body
	Payload PollersDataGetOKBody `json:"body,omitempty"`
}

// NewPollersDataGetOK creates PollersDataGetOK with default headers values
func NewPollersDataGetOK() *PollersDataGetOK {
	return &PollersDataGetOK{}
}

// WithPayload adds the payload to the pollers data get o k response
func (o *PollersDataGetOK) WithPayload(payload PollersDataGetOKBody) *PollersDataGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pollers data get o k response
func (o *PollersDataGetOK) SetPayload(payload PollersDataGetOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PollersDataGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*PollersDataGetNoContent Successfully processed the request and did not return any data

swagger:response pollersDataGetNoContent
*/
type PollersDataGetNoContent struct {

	// In: body
	Payload PollersDataGetNoContentBody `json:"body,omitempty"`
}

// NewPollersDataGetNoContent creates PollersDataGetNoContent with default headers values
func NewPollersDataGetNoContent() *PollersDataGetNoContent {
	return &PollersDataGetNoContent{}
}

// WithPayload adds the payload to the pollers data get no content response
func (o *PollersDataGetNoContent) WithPayload(payload PollersDataGetNoContentBody) *PollersDataGetNoContent {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pollers data get no content response
func (o *PollersDataGetNoContent) SetPayload(payload PollersDataGetNoContentBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PollersDataGetNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*PollersDataGetNotFound Poller with specified identifier was not found

swagger:response pollersDataGetNotFound
*/
type PollersDataGetNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPollersDataGetNotFound creates PollersDataGetNotFound with default headers values
func NewPollersDataGetNotFound() *PollersDataGetNotFound {
	return &PollersDataGetNotFound{}
}

// WithPayload adds the payload to the pollers data get not found response
func (o *PollersDataGetNotFound) WithPayload(payload *models.Error) *PollersDataGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pollers data get not found response
func (o *PollersDataGetNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PollersDataGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PollersDataGetDefault Unexpected error

swagger:response pollersDataGetDefault
*/
type PollersDataGetDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPollersDataGetDefault creates PollersDataGetDefault with default headers values
func NewPollersDataGetDefault(code int) *PollersDataGetDefault {
	if code <= 0 {
		code = 500
	}

	return &PollersDataGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the pollers data get default response
func (o *PollersDataGetDefault) WithStatusCode(code int) *PollersDataGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the pollers data get default response
func (o *PollersDataGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the pollers data get default response
func (o *PollersDataGetDefault) WithPayload(payload *models.Error) *PollersDataGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pollers data get default response
func (o *PollersDataGetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PollersDataGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}