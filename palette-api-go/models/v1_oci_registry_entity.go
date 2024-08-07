// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OciRegistryEntity Oci registry credentials
//
// swagger:model v1OciRegistryEntity
type V1OciRegistryEntity struct {

	// auth
	Auth *V1RegistryAuth `json:"auth,omitempty"`

	// default region
	DefaultRegion string `json:"defaultRegion,omitempty"`

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// provider type
	ProviderType string `json:"providerType,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 oci registry entity
func (m *V1OciRegistryEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OciRegistryEntity) validateAuth(formats strfmt.Registry) error {

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
func (m *V1OciRegistryEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OciRegistryEntity) UnmarshalBinary(b []byte) error {
	var res V1OciRegistryEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
