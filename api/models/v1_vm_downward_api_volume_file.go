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

// V1VMDownwardAPIVolumeFile DownwardAPIVolumeFile represents information to create the file containing the pod field
//
// swagger:model v1VmDownwardApiVolumeFile
type V1VMDownwardAPIVolumeFile struct {

	// field ref
	FieldRef *V1VMObjectFieldSelector `json:"fieldRef,omitempty"`

	// Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Mode int32 `json:"mode,omitempty"`

	// Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	// Required: true
	Path *string `json:"path"`

	// resource field ref
	ResourceFieldRef *V1VMResourceFieldSelector `json:"resourceFieldRef,omitempty"`
}

// Validate validates this v1 Vm downward Api volume file
func (m *V1VMDownwardAPIVolumeFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceFieldRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMDownwardAPIVolumeFile) validateFieldRef(formats strfmt.Registry) error {

	if swag.IsZero(m.FieldRef) { // not required
		return nil
	}

	if m.FieldRef != nil {
		if err := m.FieldRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fieldRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMDownwardAPIVolumeFile) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *V1VMDownwardAPIVolumeFile) validateResourceFieldRef(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceFieldRef) { // not required
		return nil
	}

	if m.ResourceFieldRef != nil {
		if err := m.ResourceFieldRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceFieldRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMDownwardAPIVolumeFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMDownwardAPIVolumeFile) UnmarshalBinary(b []byte) error {
	var res V1VMDownwardAPIVolumeFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}