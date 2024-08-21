// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MachineManagementConfig v1 machine management config
//
// swagger:model v1MachineManagementConfig
type V1MachineManagementConfig struct {

	// OS patch config contains properties to patch node os with latest security packages.
	// If OsPatchConfig is not provided then node os will not be patched with latest security updates.
	// Note: For edge based cluster (like edge-native type) the osPatchConfig is NOT applicable, the values will be ignored.
	//
	OsPatchConfig *V1OsPatchConfig `json:"osPatchConfig,omitempty"`
}

// Validate validates this v1 machine management config
func (m *V1MachineManagementConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOsPatchConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineManagementConfig) validateOsPatchConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.OsPatchConfig) { // not required
		return nil
	}

	if m.OsPatchConfig != nil {
		if err := m.OsPatchConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osPatchConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineManagementConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineManagementConfig) UnmarshalBinary(b []byte) error {
	var res V1MachineManagementConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}