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

// NewDeleteNotesParams creates a new DeleteNotesParams object
// with the default values initialized.
func NewDeleteNotesParams() DeleteNotesParams {
	var ()
	return DeleteNotesParams{}
}

// DeleteNotesParams contains all the bound params for the delete notes operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteNotes
type DeleteNotesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*note id
	  Required: true
	  Minimum: 1
	  In: query
	*/
	ID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteNotesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qID, qhkID, _ := qs.GetOK("id")
	if err := o.bindID(qID, qhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteNotesParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *DeleteNotesParams) validateID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("id", "query", int64(o.ID), 1, false); err != nil {
		return err
	}

	return nil
}
