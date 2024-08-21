// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SystemUserMeStatus User status with permissions
//
// swagger:model v1SystemUserMeStatus
type V1SystemUserMeStatus struct {

	// is email set
	IsEmailSet bool `json:"isEmailSet"`

	// is email verified
	IsEmailVerified bool `json:"isEmailVerified"`

	// is mfa enabled
	IsMfaEnabled bool `json:"isMfaEnabled"`

	// is password reset
	IsPasswordReset bool `json:"isPasswordReset"`

	// last email update time
	// Format: date-time
	LastEmailUpdateTime V1Time `json:"lastEmailUpdateTime,omitempty"`

	// last email verified time
	// Format: date-time
	LastEmailVerifiedTime V1Time `json:"lastEmailVerifiedTime,omitempty"`

	// last login time
	// Format: date-time
	LastLoginTime V1Time `json:"lastLoginTime,omitempty"`

	// last password update time
	// Format: date-time
	LastPasswordUpdateTime V1Time `json:"lastPasswordUpdateTime,omitempty"`
}

// Validate validates this v1 system user me status
func (m *V1SystemUserMeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastEmailUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastEmailVerifiedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastLoginTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastPasswordUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SystemUserMeStatus) validateLastEmailUpdateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastEmailUpdateTime) { // not required
		return nil
	}

	if err := m.LastEmailUpdateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastEmailUpdateTime")
		}
		return err
	}

	return nil
}

func (m *V1SystemUserMeStatus) validateLastEmailVerifiedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastEmailVerifiedTime) { // not required
		return nil
	}

	if err := m.LastEmailVerifiedTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastEmailVerifiedTime")
		}
		return err
	}

	return nil
}

func (m *V1SystemUserMeStatus) validateLastLoginTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastLoginTime) { // not required
		return nil
	}

	if err := m.LastLoginTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastLoginTime")
		}
		return err
	}

	return nil
}

func (m *V1SystemUserMeStatus) validateLastPasswordUpdateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastPasswordUpdateTime) { // not required
		return nil
	}

	if err := m.LastPasswordUpdateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastPasswordUpdateTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SystemUserMeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SystemUserMeStatus) UnmarshalBinary(b []byte) error {
	var res V1SystemUserMeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}