// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1UserAssetsLocationAzureSpec Azure location specification
//
// swagger:model v1UserAssetsLocationAzureSpec
type V1UserAssetsLocationAzureSpec struct {

	// config
	// Required: true
	Config *V1AzureStorageConfig `json:"config"`

	// Set to 'true', if location has to be set as default
	IsDefault bool `json:"isDefault,omitempty"`

	// Azure location type [azure]
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 user assets location azure spec
func (m *V1UserAssetsLocationAzureSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UserAssetsLocationAzureSpec) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1UserAssetsLocationAzureSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UserAssetsLocationAzureSpec) UnmarshalBinary(b []byte) error {
	var res V1UserAssetsLocationAzureSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
