// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OsPatchEntity v1 os patch entity
//
// swagger:model v1OsPatchEntity
type V1OsPatchEntity struct {

	// os patch config
	OsPatchConfig *V1OsPatchConfig `json:"osPatchConfig,omitempty"`
}

// Validate validates this v1 os patch entity
func (m *V1OsPatchEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOsPatchConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OsPatchEntity) validateOsPatchConfig(formats strfmt.Registry) error {

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
func (m *V1OsPatchEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OsPatchEntity) UnmarshalBinary(b []byte) error {
	var res V1OsPatchEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}