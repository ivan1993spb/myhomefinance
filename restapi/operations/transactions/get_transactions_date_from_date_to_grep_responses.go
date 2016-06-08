package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ivan1993spb/myhomefinance/models"
)

/*GetTransactionsDateFromDateToGrepOK list of transactions

swagger:response getTransactionsDateFromDateToGrepOK
*/
type GetTransactionsDateFromDateToGrepOK struct {

	// In: body
	Payload []*models.Transaction `json:"body,omitempty"`
}

// NewGetTransactionsDateFromDateToGrepOK creates GetTransactionsDateFromDateToGrepOK with default headers values
func NewGetTransactionsDateFromDateToGrepOK() *GetTransactionsDateFromDateToGrepOK {
	return &GetTransactionsDateFromDateToGrepOK{}
}

// WithPayload adds the payload to the get transactions date from date to grep o k response
func (o *GetTransactionsDateFromDateToGrepOK) WithPayload(payload []*models.Transaction) *GetTransactionsDateFromDateToGrepOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transactions date from date to grep o k response
func (o *GetTransactionsDateFromDateToGrepOK) SetPayload(payload []*models.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionsDateFromDateToGrepOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetTransactionsDateFromDateToGrepBadRequest invalid date range or name supplied

swagger:response getTransactionsDateFromDateToGrepBadRequest
*/
type GetTransactionsDateFromDateToGrepBadRequest struct {
}

// NewGetTransactionsDateFromDateToGrepBadRequest creates GetTransactionsDateFromDateToGrepBadRequest with default headers values
func NewGetTransactionsDateFromDateToGrepBadRequest() *GetTransactionsDateFromDateToGrepBadRequest {
	return &GetTransactionsDateFromDateToGrepBadRequest{}
}

// WriteResponse to the client
func (o *GetTransactionsDateFromDateToGrepBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

/*GetTransactionsDateFromDateToGrepServiceUnavailable server is currently unavailable

swagger:response getTransactionsDateFromDateToGrepServiceUnavailable
*/
type GetTransactionsDateFromDateToGrepServiceUnavailable struct {
}

// NewGetTransactionsDateFromDateToGrepServiceUnavailable creates GetTransactionsDateFromDateToGrepServiceUnavailable with default headers values
func NewGetTransactionsDateFromDateToGrepServiceUnavailable() *GetTransactionsDateFromDateToGrepServiceUnavailable {
	return &GetTransactionsDateFromDateToGrepServiceUnavailable{}
}

// WriteResponse to the client
func (o *GetTransactionsDateFromDateToGrepServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
}
