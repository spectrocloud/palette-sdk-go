// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMEphemeralVolumeSource v1 Vm ephemeral volume source
//
// swagger:model v1VmEphemeralVolumeSource
type V1VMEphemeralVolumeSource struct {

	// persistent volume claim
	PersistentVolumeClaim *V1VMPersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`
}

// Validate validates this v1 Vm ephemeral volume source
func (m *V1VMEphemeralVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePersistentVolumeClaim(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMEphemeralVolumeSource) validatePersistentVolumeClaim(formats strfmt.Registry) error {

	if swag.IsZero(m.PersistentVolumeClaim) { // not required
		return nil
	}

	if m.PersistentVolumeClaim != nil {
		if err := m.PersistentVolumeClaim.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("persistentVolumeClaim")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMEphemeralVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMEphemeralVolumeSource) UnmarshalBinary(b []byte) error {
	var res V1VMEphemeralVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
