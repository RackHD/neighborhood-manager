package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/rackhd/api/swagger/models"
)

/*ModifyRoleOK Successfully modified the role

swagger:response modifyRoleOK
*/
type ModifyRoleOK struct {

	// In: body
	Payload ModifyRoleOKBody `json:"body,omitempty"`
}

// NewModifyRoleOK creates ModifyRoleOK with default headers values
func NewModifyRoleOK() *ModifyRoleOK {
	return &ModifyRoleOK{}
}

// WithPayload adds the payload to the modify role o k response
func (o *ModifyRoleOK) WithPayload(payload ModifyRoleOKBody) *ModifyRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the modify role o k response
func (o *ModifyRoleOK) SetPayload(payload ModifyRoleOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ModifyRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ModifyRoleUnauthorized Unauthorized

swagger:response modifyRoleUnauthorized
*/
type ModifyRoleUnauthorized struct {

	// In: body
	Payload ModifyRoleUnauthorizedBody `json:"body,omitempty"`
}

// NewModifyRoleUnauthorized creates ModifyRoleUnauthorized with default headers values
func NewModifyRoleUnauthorized() *ModifyRoleUnauthorized {
	return &ModifyRoleUnauthorized{}
}

// WithPayload adds the payload to the modify role unauthorized response
func (o *ModifyRoleUnauthorized) WithPayload(payload ModifyRoleUnauthorizedBody) *ModifyRoleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the modify role unauthorized response
func (o *ModifyRoleUnauthorized) SetPayload(payload ModifyRoleUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ModifyRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ModifyRoleForbidden Forbidden

swagger:response modifyRoleForbidden
*/
type ModifyRoleForbidden struct {

	// In: body
	Payload ModifyRoleForbiddenBody `json:"body,omitempty"`
}

// NewModifyRoleForbidden creates ModifyRoleForbidden with default headers values
func NewModifyRoleForbidden() *ModifyRoleForbidden {
	return &ModifyRoleForbidden{}
}

// WithPayload adds the payload to the modify role forbidden response
func (o *ModifyRoleForbidden) WithPayload(payload ModifyRoleForbiddenBody) *ModifyRoleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the modify role forbidden response
func (o *ModifyRoleForbidden) SetPayload(payload ModifyRoleForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ModifyRoleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ModifyRoleDefault Unexpected error

swagger:response modifyRoleDefault
*/
type ModifyRoleDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewModifyRoleDefault creates ModifyRoleDefault with default headers values
func NewModifyRoleDefault(code int) *ModifyRoleDefault {
	if code <= 0 {
		code = 500
	}

	return &ModifyRoleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the modify role default response
func (o *ModifyRoleDefault) WithStatusCode(code int) *ModifyRoleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the modify role default response
func (o *ModifyRoleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the modify role default response
func (o *ModifyRoleDefault) WithPayload(payload *models.Error) *ModifyRoleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the modify role default response
func (o *ModifyRoleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ModifyRoleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
