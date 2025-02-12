// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CloudAccountStatus Status of the account
//
// swagger:model v1CloudAccountStatus
type V1CloudAccountStatus struct {

	// Cloud account status
	State string `json:"state,omitempty"`

	// Token expiry time
	// Format: date-time
	TokenExpiry V1Time `json:"tokenExpiry,omitempty"`

	// Token generation time
	// Format: date-time
	TokenGenerationTime V1Time `json:"tokenGenerationTime,omitempty"`
}

// Validate validates this v1 cloud account status
func (m *V1CloudAccountStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTokenExpiry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenGenerationTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CloudAccountStatus) validateTokenExpiry(formats strfmt.Registry) error {

	if swag.IsZero(m.TokenExpiry) { // not required
		return nil
	}

	if err := m.TokenExpiry.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tokenExpiry")
		}
		return err
	}

	return nil
}

func (m *V1CloudAccountStatus) validateTokenGenerationTime(formats strfmt.Registry) error {

	if swag.IsZero(m.TokenGenerationTime) { // not required
		return nil
	}

	if err := m.TokenGenerationTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tokenGenerationTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CloudAccountStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CloudAccountStatus) UnmarshalBinary(b []byte) error {
	var res V1CloudAccountStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
