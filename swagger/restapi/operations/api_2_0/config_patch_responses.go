package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*ConfigPatchOK Successfully modified the configuration

swagger:response configPatchOK
*/
type ConfigPatchOK struct {

	// In: body
	Payload ConfigPatchOKBody `json:"body,omitempty"`
}

// NewConfigPatchOK creates ConfigPatchOK with default headers values
func NewConfigPatchOK() *ConfigPatchOK {
	return &ConfigPatchOK{}
}

// WithPayload adds the payload to the config patch o k response
func (o *ConfigPatchOK) WithPayload(payload ConfigPatchOKBody) *ConfigPatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config patch o k response
func (o *ConfigPatchOK) SetPayload(payload ConfigPatchOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigPatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ConfigPatchDefault Unexpected error

swagger:response configPatchDefault
*/
type ConfigPatchDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewConfigPatchDefault creates ConfigPatchDefault with default headers values
func NewConfigPatchDefault(code int) *ConfigPatchDefault {
	if code <= 0 {
		code = 500
	}

	return &ConfigPatchDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the config patch default response
func (o *ConfigPatchDefault) WithStatusCode(code int) *ConfigPatchDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the config patch default response
func (o *ConfigPatchDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the config patch default response
func (o *ConfigPatchDefault) WithPayload(payload *models.Error) *ConfigPatchDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config patch default response
func (o *ConfigPatchDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigPatchDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
