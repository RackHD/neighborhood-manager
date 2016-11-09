package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*NodesGetTagsByIDOK Successfully retrieved node tags

swagger:response nodesGetTagsByIdOK
*/
type NodesGetTagsByIDOK struct {

	// In: body
	Payload []interface{} `json:"body,omitempty"`
}

// NewNodesGetTagsByIDOK creates NodesGetTagsByIDOK with default headers values
func NewNodesGetTagsByIDOK() *NodesGetTagsByIDOK {
	return &NodesGetTagsByIDOK{}
}

// WithPayload adds the payload to the nodes get tags by Id o k response
func (o *NodesGetTagsByIDOK) WithPayload(payload []interface{}) *NodesGetTagsByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get tags by Id o k response
func (o *NodesGetTagsByIDOK) SetPayload(payload []interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetTagsByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*NodesGetTagsByIDNotFound The specified node was not found

swagger:response nodesGetTagsByIdNotFound
*/
type NodesGetTagsByIDNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewNodesGetTagsByIDNotFound creates NodesGetTagsByIDNotFound with default headers values
func NewNodesGetTagsByIDNotFound() *NodesGetTagsByIDNotFound {
	return &NodesGetTagsByIDNotFound{}
}

// WithPayload adds the payload to the nodes get tags by Id not found response
func (o *NodesGetTagsByIDNotFound) WithPayload(payload *models.Error) *NodesGetTagsByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get tags by Id not found response
func (o *NodesGetTagsByIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetTagsByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NodesGetTagsByIDDefault Unexpected error

swagger:response nodesGetTagsByIdDefault
*/
type NodesGetTagsByIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewNodesGetTagsByIDDefault creates NodesGetTagsByIDDefault with default headers values
func NewNodesGetTagsByIDDefault(code int) *NodesGetTagsByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &NodesGetTagsByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the nodes get tags by Id default response
func (o *NodesGetTagsByIDDefault) WithStatusCode(code int) *NodesGetTagsByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the nodes get tags by Id default response
func (o *NodesGetTagsByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the nodes get tags by Id default response
func (o *NodesGetTagsByIDDefault) WithPayload(payload *models.Error) *NodesGetTagsByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get tags by Id default response
func (o *NodesGetTagsByIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetTagsByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
