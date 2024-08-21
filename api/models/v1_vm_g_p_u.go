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

// V1VMGPU v1 Vm g p u
//
// swagger:model v1VmGPU
type V1VMGPU struct {

	// device name
	// Required: true
	DeviceName *string `json:"deviceName"`

	// Name of the GPU device as exposed by a device plugin
	// Required: true
	Name *string `json:"name"`

	// If specified, the virtual network interface address and its tag will be provided to the guest via config drive
	Tag string `json:"tag,omitempty"`

	// virtual g p u options
	VirtualGPUOptions *V1VMVGPUOptions `json:"virtualGPUOptions,omitempty"`
}

// Validate validates this v1 Vm g p u
func (m *V1VMGPU) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualGPUOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMGPU) validateDeviceName(formats strfmt.Registry) error {

	if err := validate.Required("deviceName", "body", m.DeviceName); err != nil {
		return err
	}

	return nil
}

func (m *V1VMGPU) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1VMGPU) validateVirtualGPUOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualGPUOptions) { // not required
		return nil
	}

	if m.VirtualGPUOptions != nil {
		if err := m.VirtualGPUOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualGPUOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMGPU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMGPU) UnmarshalBinary(b []byte) error {
	var res V1VMGPU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}