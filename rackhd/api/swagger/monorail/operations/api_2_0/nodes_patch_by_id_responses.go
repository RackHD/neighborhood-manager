package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/rackhd/api/swagger/models"
)

/*NodesPatchByIDOK Successfully modified the specified node

swagger:response nodesPatchByIdOK
*/
type NodesPatchByIDOK struct {

	// In: body
	Payload NodesPatchByIDOKBody `json:"body,omitempty"`
}

// NewNodesPatchByIDOK creates NodesPatchByIDOK with default headers values
func NewNodesPatchByIDOK() *NodesPatchByIDOK {
	return &NodesPatchByIDOK{}
}

// WithPayload adds the payload to the nodes patch by Id o k response
func (o *NodesPatchByIDOK) WithPayload(payload NodesPatchByIDOKBody) *NodesPatchByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes patch by Id o k response
func (o *NodesPatchByIDOK) SetPayload(payload NodesPatchByIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesPatchByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*NodesPatchByIDNotFound The specified node was not found

swagger:response nodesPatchByIdNotFound
*/
type NodesPatchByIDNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewNodesPatchByIDNotFound creates NodesPatchByIDNotFound with default headers values
func NewNodesPatchByIDNotFound() *NodesPatchByIDNotFound {
	return &NodesPatchByIDNotFound{}
}

// WithPayload adds the payload to the nodes patch by Id not found response
func (o *NodesPatchByIDNotFound) WithPayload(payload *models.Error) *NodesPatchByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes patch by Id not found response
func (o *NodesPatchByIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesPatchByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NodesPatchByIDDefault Unexpected error

swagger:response nodesPatchByIdDefault
*/
type NodesPatchByIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewNodesPatchByIDDefault creates NodesPatchByIDDefault with default headers values
func NewNodesPatchByIDDefault(code int) *NodesPatchByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &NodesPatchByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the nodes patch by Id default response
func (o *NodesPatchByIDDefault) WithStatusCode(code int) *NodesPatchByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the nodes patch by Id default response
func (o *NodesPatchByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the nodes patch by Id default response
func (o *NodesPatchByIDDefault) WithPayload(payload *models.Error) *NodesPatchByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes patch by Id default response
func (o *NodesPatchByIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesPatchByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
