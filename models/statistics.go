package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Statistics statistics

swagger:model Statistics
*/
type Statistics struct {

	/* balance at the end of date range

	Required: true
	*/
	BalanceEnd *float64 `json:"balance_end"`

	/* balance at the start of date range

	Required: true
	*/
	BalanceStart *float64 `json:"balance_start"`

	/* date from

	Required: true
	*/
	DateFrom *strfmt.Date `json:"date_from"`

	/* date to

	Required: true
	*/
	DateTo *strfmt.Date `json:"date_to"`

	/* inflow

	Required: true
	*/
	Inflow *float64 `json:"inflow"`

	/* mean satisfaction

	Required: true
	*/
	MeanSatisfaction *float32 `json:"mean_satisfaction"`

	/* outflow

	Required: true
	*/
	Outflow *float64 `json:"outflow"`
}

// Validate validates this statistics
func (m *Statistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalanceEnd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBalanceStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDateFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDateTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInflow(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeanSatisfaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutflow(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Statistics) validateBalanceEnd(formats strfmt.Registry) error {

	if err := validate.Required("balance_end", "body", m.BalanceEnd); err != nil {
		return err
	}

	return nil
}

func (m *Statistics) validateBalanceStart(formats strfmt.Registry) error {

	if err := validate.Required("balance_start", "body", m.BalanceStart); err != nil {
		return err
	}

	return nil
}

func (m *Statistics) validateDateFrom(formats strfmt.Registry) error {

	if err := validate.Required("date_from", "body", m.DateFrom); err != nil {
		return err
	}

	return nil
}

func (m *Statistics) validateDateTo(formats strfmt.Registry) error {

	if err := validate.Required("date_to", "body", m.DateTo); err != nil {
		return err
	}

	return nil
}

func (m *Statistics) validateInflow(formats strfmt.Registry) error {

	if err := validate.Required("inflow", "body", m.Inflow); err != nil {
		return err
	}

	return nil
}

func (m *Statistics) validateMeanSatisfaction(formats strfmt.Registry) error {

	if err := validate.Required("mean_satisfaction", "body", m.MeanSatisfaction); err != nil {
		return err
	}

	return nil
}

func (m *Statistics) validateOutflow(formats strfmt.Registry) error {

	if err := validate.Required("outflow", "body", m.Outflow); err != nil {
		return err
	}

	return nil
}