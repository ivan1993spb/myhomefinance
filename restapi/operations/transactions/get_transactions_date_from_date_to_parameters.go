package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTransactionsDateFromDateToParams creates a new GetTransactionsDateFromDateToParams object
// with the default values initialized.
func NewGetTransactionsDateFromDateToParams() GetTransactionsDateFromDateToParams {
	var ()
	return GetTransactionsDateFromDateToParams{}
}

// GetTransactionsDateFromDateToParams contains all the bound params for the get transactions date from date to operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetTransactionsDateFromDateTo
type GetTransactionsDateFromDateToParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: path
	*/
	DateFrom strfmt.Date
	/*
	  Required: true
	  In: path
	*/
	DateTo strfmt.Date
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetTransactionsDateFromDateToParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rDateFrom, rhkDateFrom, _ := route.Params.GetOK("date_from")
	if err := o.bindDateFrom(rDateFrom, rhkDateFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	rDateTo, rhkDateTo, _ := route.Params.GetOK("date_to")
	if err := o.bindDateTo(rDateTo, rhkDateTo, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionsDateFromDateToParams) bindDateFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := formats.Parse("date", raw)
	if err != nil {
		return errors.InvalidType("date_from", "path", "strfmt.Date", raw)
	}
	o.DateFrom = *(value.(*strfmt.Date))

	return nil
}

func (o *GetTransactionsDateFromDateToParams) bindDateTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := formats.Parse("date", raw)
	if err != nil {
		return errors.InvalidType("date_to", "path", "strfmt.Date", raw)
	}
	o.DateTo = *(value.(*strfmt.Date))

	return nil
}
