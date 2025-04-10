// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMRealtime Realtime holds the tuning knobs specific for realtime workloads.
//
// swagger:model v1VmRealtime
type V1VMRealtime struct {

	// Mask defines the vcpu mask expression that defines which vcpus are used for realtime. Format matches libvirt's expressions. Example: "0-3,^1","0,2,3","2-3"
	Mask string `json:"mask,omitempty"`
}

// Validate validates this v1 Vm realtime
func (m *V1VMRealtime) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 Vm realtime based on context it is used
func (m *V1VMRealtime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMRealtime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMRealtime) UnmarshalBinary(b []byte) error {
	var res V1VMRealtime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
