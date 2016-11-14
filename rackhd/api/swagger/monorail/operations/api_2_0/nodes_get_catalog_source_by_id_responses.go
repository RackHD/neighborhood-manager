package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/rackhd/api/swagger/models"
)

/*NodesGetCatalogSourceByIDOK Successfully retrieved specific source catalog of specified node

swagger:response nodesGetCatalogSourceByIdOK
*/
type NodesGetCatalogSourceByIDOK struct {

	// In: body
	Payload NodesGetCatalogSourceByIDOKBody `json:"body,omitempty"`
}

// NewNodesGetCatalogSourceByIDOK creates NodesGetCatalogSourceByIDOK with default headers values
func NewNodesGetCatalogSourceByIDOK() *NodesGetCatalogSourceByIDOK {
	return &NodesGetCatalogSourceByIDOK{}
}

// WithPayload adds the payload to the nodes get catalog source by Id o k response
func (o *NodesGetCatalogSourceByIDOK) WithPayload(payload NodesGetCatalogSourceByIDOKBody) *NodesGetCatalogSourceByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get catalog source by Id o k response
func (o *NodesGetCatalogSourceByIDOK) SetPayload(payload NodesGetCatalogSourceByIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetCatalogSourceByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*NodesGetCatalogSourceByIDNotFound The specified node was not found

swagger:response nodesGetCatalogSourceByIdNotFound
*/
type NodesGetCatalogSourceByIDNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewNodesGetCatalogSourceByIDNotFound creates NodesGetCatalogSourceByIDNotFound with default headers values
func NewNodesGetCatalogSourceByIDNotFound() *NodesGetCatalogSourceByIDNotFound {
	return &NodesGetCatalogSourceByIDNotFound{}
}

// WithPayload adds the payload to the nodes get catalog source by Id not found response
func (o *NodesGetCatalogSourceByIDNotFound) WithPayload(payload *models.Error) *NodesGetCatalogSourceByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get catalog source by Id not found response
func (o *NodesGetCatalogSourceByIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetCatalogSourceByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NodesGetCatalogSourceByIDDefault Unexpected error

swagger:response nodesGetCatalogSourceByIdDefault
*/
type NodesGetCatalogSourceByIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewNodesGetCatalogSourceByIDDefault creates NodesGetCatalogSourceByIDDefault with default headers values
func NewNodesGetCatalogSourceByIDDefault(code int) *NodesGetCatalogSourceByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &NodesGetCatalogSourceByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the nodes get catalog source by Id default response
func (o *NodesGetCatalogSourceByIDDefault) WithStatusCode(code int) *NodesGetCatalogSourceByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the nodes get catalog source by Id default response
func (o *NodesGetCatalogSourceByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the nodes get catalog source by Id default response
func (o *NodesGetCatalogSourceByIDDefault) WithPayload(payload *models.Error) *NodesGetCatalogSourceByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get catalog source by Id default response
func (o *NodesGetCatalogSourceByIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetCatalogSourceByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
