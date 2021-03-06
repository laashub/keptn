// Code generated by go-swagger; DO NOT EDIT.

package service_default_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/keptn/keptn/configuration-service/models"
)

// PutProjectProjectNameServiceServiceNameResourceCreatedCode is the HTTP code returned for type PutProjectProjectNameServiceServiceNameResourceCreated
const PutProjectProjectNameServiceServiceNameResourceCreatedCode int = 201

/*PutProjectProjectNameServiceServiceNameResourceCreated Success. Service default resources have been updated. The version of the new configuration is returned.

swagger:response putProjectProjectNameServiceServiceNameResourceCreated
*/
type PutProjectProjectNameServiceServiceNameResourceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Version `json:"body,omitempty"`
}

// NewPutProjectProjectNameServiceServiceNameResourceCreated creates PutProjectProjectNameServiceServiceNameResourceCreated with default headers values
func NewPutProjectProjectNameServiceServiceNameResourceCreated() *PutProjectProjectNameServiceServiceNameResourceCreated {

	return &PutProjectProjectNameServiceServiceNameResourceCreated{}
}

// WithPayload adds the payload to the put project project name service service name resource created response
func (o *PutProjectProjectNameServiceServiceNameResourceCreated) WithPayload(payload *models.Version) *PutProjectProjectNameServiceServiceNameResourceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put project project name service service name resource created response
func (o *PutProjectProjectNameServiceServiceNameResourceCreated) SetPayload(payload *models.Version) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutProjectProjectNameServiceServiceNameResourceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutProjectProjectNameServiceServiceNameResourceBadRequestCode is the HTTP code returned for type PutProjectProjectNameServiceServiceNameResourceBadRequest
const PutProjectProjectNameServiceServiceNameResourceBadRequestCode int = 400

/*PutProjectProjectNameServiceServiceNameResourceBadRequest Failed. Service default resources could not be updated.

swagger:response putProjectProjectNameServiceServiceNameResourceBadRequest
*/
type PutProjectProjectNameServiceServiceNameResourceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutProjectProjectNameServiceServiceNameResourceBadRequest creates PutProjectProjectNameServiceServiceNameResourceBadRequest with default headers values
func NewPutProjectProjectNameServiceServiceNameResourceBadRequest() *PutProjectProjectNameServiceServiceNameResourceBadRequest {

	return &PutProjectProjectNameServiceServiceNameResourceBadRequest{}
}

// WithPayload adds the payload to the put project project name service service name resource bad request response
func (o *PutProjectProjectNameServiceServiceNameResourceBadRequest) WithPayload(payload *models.Error) *PutProjectProjectNameServiceServiceNameResourceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put project project name service service name resource bad request response
func (o *PutProjectProjectNameServiceServiceNameResourceBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutProjectProjectNameServiceServiceNameResourceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutProjectProjectNameServiceServiceNameResourceDefault Error

swagger:response putProjectProjectNameServiceServiceNameResourceDefault
*/
type PutProjectProjectNameServiceServiceNameResourceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutProjectProjectNameServiceServiceNameResourceDefault creates PutProjectProjectNameServiceServiceNameResourceDefault with default headers values
func NewPutProjectProjectNameServiceServiceNameResourceDefault(code int) *PutProjectProjectNameServiceServiceNameResourceDefault {
	if code <= 0 {
		code = 500
	}

	return &PutProjectProjectNameServiceServiceNameResourceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put project project name service service name resource default response
func (o *PutProjectProjectNameServiceServiceNameResourceDefault) WithStatusCode(code int) *PutProjectProjectNameServiceServiceNameResourceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put project project name service service name resource default response
func (o *PutProjectProjectNameServiceServiceNameResourceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put project project name service service name resource default response
func (o *PutProjectProjectNameServiceServiceNameResourceDefault) WithPayload(payload *models.Error) *PutProjectProjectNameServiceServiceNameResourceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put project project name service service name resource default response
func (o *PutProjectProjectNameServiceServiceNameResourceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutProjectProjectNameServiceServiceNameResourceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
