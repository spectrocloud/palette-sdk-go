// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1VMMemoryDumpVolumeSource v1 Vm memory dump volume source
//
// swagger:model v1VmMemoryDumpVolumeSource
type V1VMMemoryDumpVolumeSource struct {

	// ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// Required: true
	ClaimName *string `json:"claimName"`

	// Hotpluggable indicates whether the volume can be hotplugged and hotunplugged.
	Hotpluggable bool `json:"hotpluggable,omitempty"`

	// Will force the ReadOnly setting in VolumeMounts. Default false.
	ReadOnly bool `json:"readOnly,omitempty"`
}

// Validate validates this v1 Vm memory dump volume source
func (m *V1VMMemoryDumpVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClaimName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMMemoryDumpVolumeSource) validateClaimName(formats strfmt.Registry) error {

	if err := validate.Required("claimName", "body", m.ClaimName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 Vm memory dump volume source based on context it is used
func (m *V1VMMemoryDumpVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMMemoryDumpVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMMemoryDumpVolumeSource) UnmarshalBinary(b []byte) error {
	var res V1VMMemoryDumpVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
