// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1InvoiceSpec Invoice specification
//
// swagger:model v1InvoiceSpec
type V1InvoiceSpec struct {

	// address
	Address *V1Address `json:"address,omitempty"`

	// billing period
	BillingPeriod *V1InvoiceBillingPeriod `json:"billingPeriod,omitempty"`

	// credits
	Credits *V1InvoiceCredits `json:"credits,omitempty"`

	// Environment type [Trial,MonthlyOnDemand,AnnualSubscription,OnPrem]
	EnvType string `json:"envType,omitempty"`

	// Month for which invoice is generated
	// Format: date-time
	Month V1Time `json:"month,omitempty"`

	// payment unit
	// Enum: [usd]
	PaymentUnit string `json:"paymentUnit,omitempty"`

	// plan
	Plan *V1InvoicePlan `json:"plan,omitempty"`
}

// Validate validates this v1 invoice spec
func (m *V1InvoiceSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1InvoiceSpec) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *V1InvoiceSpec) validateBillingPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingPeriod) { // not required
		return nil
	}

	if m.BillingPeriod != nil {
		if err := m.BillingPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingPeriod")
			}
			return err
		}
	}

	return nil
}

func (m *V1InvoiceSpec) validateCredits(formats strfmt.Registry) error {

	if swag.IsZero(m.Credits) { // not required
		return nil
	}

	if m.Credits != nil {
		if err := m.Credits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credits")
			}
			return err
		}
	}

	return nil
}

func (m *V1InvoiceSpec) validateMonth(formats strfmt.Registry) error {

	if swag.IsZero(m.Month) { // not required
		return nil
	}

	if err := m.Month.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("month")
		}
		return err
	}

	return nil
}

var v1InvoiceSpecTypePaymentUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["usd"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1InvoiceSpecTypePaymentUnitPropEnum = append(v1InvoiceSpecTypePaymentUnitPropEnum, v)
	}
}

const (

	// V1InvoiceSpecPaymentUnitUsd captures enum value "usd"
	V1InvoiceSpecPaymentUnitUsd string = "usd"
)

// prop value enum
func (m *V1InvoiceSpec) validatePaymentUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1InvoiceSpecTypePaymentUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1InvoiceSpec) validatePaymentUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentUnitEnum("paymentUnit", "body", m.PaymentUnit); err != nil {
		return err
	}

	return nil
}

func (m *V1InvoiceSpec) validatePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	if m.Plan != nil {
		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1InvoiceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1InvoiceSpec) UnmarshalBinary(b []byte) error {
	var res V1InvoiceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}