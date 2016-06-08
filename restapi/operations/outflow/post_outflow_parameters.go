package outflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ivan1993spb/myhomefinance/models"
)

// NewPostOutflowParams creates a new PostOutflowParams object
// with the default values initialized.
func NewPostOutflowParams() PostOutflowParams {
	var ()
	return PostOutflowParams{}
}

// PostOutflowParams contains all the bound params for the post outflow operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostOutflow
type PostOutflowParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*object contains outflow data

	  Required: true
	  In: body
	*/
	OutflowRawData *models.OutflowRaw
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PostOutflowParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.OutflowRaw
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("outflowRawData", "body"))
			} else {
				res = append(res, errors.NewParseError("outflowRawData", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.OutflowRawData = &body
			}
		}

	} else {
		res = append(res, errors.Required("outflowRawData", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}