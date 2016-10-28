package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/RackHD/neighborhood-manager/swagger/models"
)

/*SkuPackPostCreated Successfully created the SKU Pack

swagger:response skuPackPostCreated
*/
type SkuPackPostCreated struct {

	// In: body
	Payload SkuPackPostCreatedBody `json:"body,omitempty"`
}

// NewSkuPackPostCreated creates SkuPackPostCreated with default headers values
func NewSkuPackPostCreated() *SkuPackPostCreated {
	return &SkuPackPostCreated{}
}

// WithPayload adds the payload to the sku pack post created response
func (o *SkuPackPostCreated) WithPayload(payload SkuPackPostCreatedBody) *SkuPackPostCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sku pack post created response
func (o *SkuPackPostCreated) SetPayload(payload SkuPackPostCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SkuPackPostCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*SkuPackPostDefault Unexpected error

swagger:response skuPackPostDefault
*/
type SkuPackPostDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewSkuPackPostDefault creates SkuPackPostDefault with default headers values
func NewSkuPackPostDefault(code int) *SkuPackPostDefault {
	if code <= 0 {
		code = 500
	}

	return &SkuPackPostDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the sku pack post default response
func (o *SkuPackPostDefault) WithStatusCode(code int) *SkuPackPostDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the sku pack post default response
func (o *SkuPackPostDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the sku pack post default response
func (o *SkuPackPostDefault) WithPayload(payload *models.Error) *SkuPackPostDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sku pack post default response
func (o *SkuPackPostDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SkuPackPostDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}