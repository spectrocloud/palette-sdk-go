// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1RegistryConfiguration Registry configuration
//
// swagger:model v1RegistryConfiguration
type V1RegistryConfiguration struct {

	// auth
	Auth *V1RegistryAuth `json:"auth,omitempty"`

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 registry configuration
func (m *V1RegistryConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RegistryConfiguration) validateAuth(formats strfmt.Registry) error {

	if swag.IsZero(m.Auth) { // not required
		return nil
	}

	if m.Auth != nil {
		if err := m.Auth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RegistryConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RegistryConfiguration) UnmarshalBinary(b []byte) error {
	var res V1RegistryConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
