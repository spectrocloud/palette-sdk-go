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

// V1MonthlyUsage Monthly usage object
//
// swagger:model v1MonthlyUsage
type V1MonthlyUsage struct {

	// Month of usage
	// Format: date-time
	Month V1Time `json:"month,omitempty"`

	// List of tenants usage
	// Unique: true
	TenantUsages []*V1TenantUsage `json:"tenantUsages"`

	// Credits used by imported clusters
	UsedAlloyCredits float64 `json:"usedAlloyCredits,omitempty"`

	// Credits used by managed clusters
	UsedPureCredits float64 `json:"usedPureCredits,omitempty"`
}

// Validate validates this v1 monthly usage
func (m *V1MonthlyUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantUsages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MonthlyUsage) validateMonth(formats strfmt.Registry) error {

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

func (m *V1MonthlyUsage) validateTenantUsages(formats strfmt.Registry) error {

	if swag.IsZero(m.TenantUsages) { // not required
		return nil
	}

	if err := validate.UniqueItems("tenantUsages", "body", m.TenantUsages); err != nil {
		return err
	}

	for i := 0; i < len(m.TenantUsages); i++ {
		if swag.IsZero(m.TenantUsages[i]) { // not required
			continue
		}

		if m.TenantUsages[i] != nil {
			if err := m.TenantUsages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenantUsages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MonthlyUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MonthlyUsage) UnmarshalBinary(b []byte) error {
	var res V1MonthlyUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
