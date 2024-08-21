// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AzureManagedMachinePoolConfig v1 azure managed machine pool config
//
// swagger:model v1AzureManagedMachinePoolConfig
type V1AzureManagedMachinePoolConfig struct {

	// whether this pool is for system node Pool
	IsSystemNodePool bool `json:"isSystemNodePool"`

	// os type
	OsType V1OsType `json:"osType,omitempty"`
}

// Validate validates this v1 azure managed machine pool config
func (m *V1AzureManagedMachinePoolConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AzureManagedMachinePoolConfig) validateOsType(formats strfmt.Registry) error {

	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	if err := m.OsType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("osType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AzureManagedMachinePoolConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AzureManagedMachinePoolConfig) UnmarshalBinary(b []byte) error {
	var res V1AzureManagedMachinePoolConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}