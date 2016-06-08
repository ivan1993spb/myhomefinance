package outflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ivan1993spb/myhomefinance/models"
)

/*GetOutflowOK outflow document found

swagger:response getOutflowOK
*/
type GetOutflowOK struct {

	// In: body
	Payload *models.Outflow `json:"body,omitempty"`
}

// NewGetOutflowOK creates GetOutflowOK with default headers values
func NewGetOutflowOK() *GetOutflowOK {
	return &GetOutflowOK{}
}

// WithPayload adds the payload to the get outflow o k response
func (o *GetOutflowOK) WithPayload(payload *models.Outflow) *GetOutflowOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get outflow o k response
func (o *GetOutflowOK) SetPayload(payload *models.Outflow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOutflowOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetOutflowNotFound outflow document not found

swagger:response getOutflowNotFound
*/
type GetOutflowNotFound struct {
}

// NewGetOutflowNotFound creates GetOutflowNotFound with default headers values
func NewGetOutflowNotFound() *GetOutflowNotFound {
	return &GetOutflowNotFound{}
}

// WriteResponse to the client
func (o *GetOutflowNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*GetOutflowServiceUnavailable server is currently unavailable

swagger:response getOutflowServiceUnavailable
*/
type GetOutflowServiceUnavailable struct {
}

// NewGetOutflowServiceUnavailable creates GetOutflowServiceUnavailable with default headers values
func NewGetOutflowServiceUnavailable() *GetOutflowServiceUnavailable {
	return &GetOutflowServiceUnavailable{}
}

// WriteResponse to the client
func (o *GetOutflowServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
}
