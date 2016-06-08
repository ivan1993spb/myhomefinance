package notes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutNotesParams creates a new PutNotesParams object
// with the default values initialized.
func NewPutNotesParams() PutNotesParams {
	var ()
	return PutNotesParams{}
}

// PutNotesParams contains all the bound params for the put notes operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutNotes
type PutNotesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*timestamp
	  In: formData
	*/
	Datetime *strfmt.DateTime
	/*note id
	  Required: true
	  Minimum: 1
	  In: query
	*/
	ID int64
	/*note name
	  Max Length: 300
	  Min Length: 2
	  In: formData
	*/
	Name *string
	/*note text
	  In: formData
	*/
	Text *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PutNotesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := runtime.Values(r.Form)

	fdDatetime, fdhkDatetime, _ := fds.GetOK("datetime")
	if err := o.bindDatetime(fdDatetime, fdhkDatetime, route.Formats); err != nil {
		res = append(res, err)
	}

	qID, qhkID, _ := qs.GetOK("id")
	if err := o.bindID(qID, qhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdName, fdhkName, _ := fds.GetOK("name")
	if err := o.bindName(fdName, fdhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	fdText, fdhkText, _ := fds.GetOK("text")
	if err := o.bindText(fdText, fdhkText, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutNotesParams) bindDatetime(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("datetime", "formData", "strfmt.DateTime", raw)
	}
	o.Datetime = (value.(*strfmt.DateTime))

	return nil
}

func (o *PutNotesParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("id", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("id", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "query", "int64", raw)
	}
	o.ID = value

	if err := o.validateID(formats); err != nil {
		return err
	}

	return nil
}

func (o *PutNotesParams) validateID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("id", "query", int64(o.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (o *PutNotesParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *PutNotesParams) validateName(formats strfmt.Registry) error {

	if err := validate.MinLength("name", "formData", string(*o.Name), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "formData", string(*o.Name), 300); err != nil {
		return err
	}

	return nil
}

func (o *PutNotesParams) bindText(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Text = &raw

	return nil
}
