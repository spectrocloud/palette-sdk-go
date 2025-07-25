// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1YearlyUsage Yearly usage object
//
// swagger:model v1YearlyUsage
type V1YearlyUsage struct {

	// billing period
	BillingPeriod *V1InvoiceBillingPeriod `json:"billingPeriod,omitempty"`

	// List of monthly usages
	// Unique: true
	MonthlyUsages []*V1MonthlyUsage `json:"monthlyUsages"`

	// product usages
	ProductUsages *V1ProductUsage `json:"productUsages,omitempty"`
}

// Validate validates this v1 yearly usage
func (m *V1YearlyUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonthlyUsages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductUsages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1YearlyUsage) validateBillingPeriod(formats strfmt.Registry) error {

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

func (m *V1YearlyUsage) validateMonthlyUsages(formats strfmt.Registry) error {

	if swag.IsZero(m.MonthlyUsages) { // not required
		return nil
	}

	if err := validate.UniqueItems("monthlyUsages", "body", m.MonthlyUsages); err != nil {
		return err
	}

	for i := 0; i < len(m.MonthlyUsages); i++ {
		if swag.IsZero(m.MonthlyUsages[i]) { // not required
			continue
		}

		if m.MonthlyUsages[i] != nil {
			if err := m.MonthlyUsages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("monthlyUsages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1YearlyUsage) validateProductUsages(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductUsages) { // not required
		return nil
	}

	if m.ProductUsages != nil {
		if err := m.ProductUsages.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productUsages")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1YearlyUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1YearlyUsage) UnmarshalBinary(b []byte) error {
	var res V1YearlyUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
