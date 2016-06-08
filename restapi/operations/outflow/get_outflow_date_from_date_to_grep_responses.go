package outflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ivan1993spb/myhomefinance/models"
)

/*GetOutflowDateFromDateToGrepOK list of outflow documents

swagger:response getOutflowDateFromDateToGrepOK
*/
type GetOutflowDateFromDateToGrepOK struct {

	// In: body
	Payload []*models.Outflow `json:"body,omitempty"`
}

// NewGetOutflowDateFromDateToGrepOK creates GetOutflowDateFromDateToGrepOK with default headers values
func NewGetOutflowDateFromDateToGrepOK() *GetOutflowDateFromDateToGrepOK {
	return &GetOutflowDateFromDateToGrepOK{}
}

// WithPayload adds the payload to the get outflow date from date to grep o k response
func (o *GetOutflowDateFromDateToGrepOK) WithPayload(payload []*models.Outflow) *GetOutflowDateFromDateToGrepOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get outflow date from date to grep o k response
func (o *GetOutflowDateFromDateToGrepOK) SetPayload(payload []*models.Outflow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOutflowDateFromDateToGrepOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetOutflowDateFromDateToGrepBadRequest invalid date range or other parameter(s) supplied

swagger:response getOutflowDateFromDateToGrepBadRequest
*/
type GetOutflowDateFromDateToGrepBadRequest struct {
}

// NewGetOutflowDateFromDateToGrepBadRequest creates GetOutflowDateFromDateToGrepBadRequest with default headers values
func NewGetOutflowDateFromDateToGrepBadRequest() *GetOutflowDateFromDateToGrepBadRequest {
	return &GetOutflowDateFromDateToGrepBadRequest{}
}

// WriteResponse to the client
func (o *GetOutflowDateFromDateToGrepBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

/*GetOutflowDateFromDateToGrepServiceUnavailable server is currently unavailable

swagger:response getOutflowDateFromDateToGrepServiceUnavailable
*/
type GetOutflowDateFromDateToGrepServiceUnavailable struct {
}

// NewGetOutflowDateFromDateToGrepServiceUnavailable creates GetOutflowDateFromDateToGrepServiceUnavailable with default headers values
func NewGetOutflowDateFromDateToGrepServiceUnavailable() *GetOutflowDateFromDateToGrepServiceUnavailable {
	return &GetOutflowDateFromDateToGrepServiceUnavailable{}
}

// WriteResponse to the client
func (o *GetOutflowDateFromDateToGrepServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
}
