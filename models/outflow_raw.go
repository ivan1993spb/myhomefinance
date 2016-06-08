package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*OutflowRaw OutflowRaw is raw data for building outflow documents and represents checks from stores


swagger:model OutflowRaw
*/
type OutflowRaw struct {

	/* default: current date and time
	 */
	Datetime strfmt.DateTime `json:"datetime,omitempty"`

	/* destination

	Required: true
	Max Length: 300
	Min Length: 2
	*/
	Destination *string `json:"destination"`

	/* list
	 */
	List []*OutflowRawItem `json:"list,omitempty"`
}

// Validate validates this outflow raw
func (m *OutflowRaw) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutflowRaw) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	if err := validate.MinLength("destination", "body", string(*m.Destination), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("destination", "body", string(*m.Destination), 300); err != nil {
		return err
	}

	return nil
}

func (m *OutflowRaw) validateList(formats strfmt.Registry) error {

	if swag.IsZero(m.List) { // not required
		return nil
	}

	for i := 0; i < len(m.List); i++ {

		if swag.IsZero(m.List[i]) { // not required
			continue
		}

		if m.List[i] != nil {

			if err := m.List[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}