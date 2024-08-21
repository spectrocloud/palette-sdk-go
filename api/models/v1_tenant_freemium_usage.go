// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TenantFreemiumUsage v1 tenant freemium usage
//
// swagger:model v1TenantFreemiumUsage
type V1TenantFreemiumUsage struct {

	// is freemium
	IsFreemium bool `json:"isFreemium"`

	// is unlimited
	IsUnlimited bool `json:"isUnlimited"`

	// limit
	Limit *V1FreemiumUsageLimit `json:"limit,omitempty"`

	// usage
	Usage *V1FreemiumUsage `json:"usage,omitempty"`
}

// Validate validates this v1 tenant freemium usage
func (m *V1TenantFreemiumUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TenantFreemiumUsage) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if m.Limit != nil {
		if err := m.Limit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limit")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantFreemiumUsage) validateUsage(formats strfmt.Registry) error {

	if swag.IsZero(m.Usage) { // not required
		return nil
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TenantFreemiumUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TenantFreemiumUsage) UnmarshalBinary(b []byte) error {
	var res V1TenantFreemiumUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}