package outflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOutflowDateFromDateToGrepParams creates a new GetOutflowDateFromDateToGrepParams object
// with the default values initialized.
func NewGetOutflowDateFromDateToGrepParams() GetOutflowDateFromDateToGrepParams {
	var ()
	return GetOutflowDateFromDateToGrepParams{}
}

// GetOutflowDateFromDateToGrepParams contains all the bound params for the get outflow date from date to grep operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetOutflowDateFromDateToGrep
type GetOutflowDateFromDateToGrepParams struct {

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
	/*
	  Max Length: 300
	  Min Length: 1
	  In: formData
	*/
	Destination *string
	/*
	  Max Length: 300
	  Min Length: 1
	  In: formData
	*/
	Name *string
	/*
	  Max Length: 300
	  Min Length: 1
	  In: formData
	*/
	Target *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetOutflowDateFromDateToGrepParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := runtime.Values(r.Form)

	rDateFrom, rhkDateFrom, _ := route.Params.GetOK("date_from")
	if err := o.bindDateFrom(rDateFrom, rhkDateFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	rDateTo, rhkDateTo, _ := route.Params.GetOK("date_to")
	if err := o.bindDateTo(rDateTo, rhkDateTo, route.Formats); err != nil {
		res = append(res, err)
	}

	fdDestination, fdhkDestination, _ := fds.GetOK("destination")
	if err := o.bindDestination(fdDestination, fdhkDestination, route.Formats); err != nil {
		res = append(res, err)
	}

	fdName, fdhkName, _ := fds.GetOK("name")
	if err := o.bindName(fdName, fdhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTarget, fdhkTarget, _ := fds.GetOK("target")
	if err := o.bindTarget(fdTarget, fdhkTarget, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOutflowDateFromDateToGrepParams) bindDateFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetOutflowDateFromDateToGrepParams) bindDateTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetOutflowDateFromDateToGrepParams) bindDestination(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Destination = &raw

	if err := o.validateDestination(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetOutflowDateFromDateToGrepParams) validateDestination(formats strfmt.Registry) error {

	if err := validate.MinLength("destination", "formData", string(*o.Destination), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("destination", "formData", string(*o.Destination), 300); err != nil {
		return err
	}

	return nil
}

func (o *GetOutflowDateFromDateToGrepParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Name = &raw

	if err := o.validateName(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetOutflowDateFromDateToGrepParams) validateName(formats strfmt.Registry) error {

	if err := validate.MinLength("name", "formData", string(*o.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "formData", string(*o.Name), 300); err != nil {
		return err
	}

	return nil
}

func (o *GetOutflowDateFromDateToGrepParams) bindTarget(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Target = &raw

	if err := o.validateTarget(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetOutflowDateFromDateToGrepParams) validateTarget(formats strfmt.Registry) error {

	if err := validate.MinLength("target", "formData", string(*o.Target), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("target", "formData", string(*o.Target), 300); err != nil {
		return err
	}

	return nil
}