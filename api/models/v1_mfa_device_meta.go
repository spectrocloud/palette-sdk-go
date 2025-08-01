// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MfaDeviceMeta v1 mfa device meta
//
// swagger:model v1MfaDeviceMeta
type V1MfaDeviceMeta struct {

	// creation timestamp
	// Format: date-time
	CreationTimestamp V1Time `json:"creationTimestamp,omitempty"`

	// device name
	DeviceName string `json:"deviceName,omitempty"`
}

// Validate validates this v1 mfa device meta
func (m *V1MfaDeviceMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MfaDeviceMeta) validateCreationTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := m.CreationTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("creationTimestamp")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MfaDeviceMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MfaDeviceMeta) UnmarshalBinary(b []byte) error {
	var res V1MfaDeviceMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
