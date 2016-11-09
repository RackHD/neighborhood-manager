package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*ConfigGetOK Successfully retrieved the configuration

swagger:response configGetOK
*/
type ConfigGetOK struct {

	// In: body
	Payload ConfigGetOKBody `json:"body,omitempty"`
}

// NewConfigGetOK creates ConfigGetOK with default headers values
func NewConfigGetOK() *ConfigGetOK {
	return &ConfigGetOK{}
}

// WithPayload adds the payload to the config get o k response
func (o *ConfigGetOK) WithPayload(payload ConfigGetOKBody) *ConfigGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config get o k response
func (o *ConfigGetOK) SetPayload(payload ConfigGetOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ConfigGetDefault Unexpected error

swagger:response configGetDefault
*/
type ConfigGetDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewConfigGetDefault creates ConfigGetDefault with default headers values
func NewConfigGetDefault(code int) *ConfigGetDefault {
	if code <= 0 {
		code = 500
	}

	return &ConfigGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the config get default response
func (o *ConfigGetDefault) WithStatusCode(code int) *ConfigGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the config get default response
func (o *ConfigGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the config get default response
func (o *ConfigGetDefault) WithPayload(payload *models.Error) *ConfigGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config get default response
func (o *ConfigGetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
