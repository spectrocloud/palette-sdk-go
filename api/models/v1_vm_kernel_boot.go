// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMKernelBoot Represents the firmware blob used to assist in the kernel boot process. Used for setting the kernel, initrd and command line arguments
//
// swagger:model v1VmKernelBoot
type V1VMKernelBoot struct {

	// container
	Container *V1VMKernelBootContainer `json:"container,omitempty"`

	// Arguments to be passed to the kernel at boot time
	KernelArgs string `json:"kernelArgs,omitempty"`
}

// Validate validates this v1 Vm kernel boot
func (m *V1VMKernelBoot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMKernelBoot) validateContainer(formats strfmt.Registry) error {
	if swag.IsZero(m.Container) { // not required
		return nil
	}

	if m.Container != nil {
		if err := m.Container.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("container")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 Vm kernel boot based on the context it is used
func (m *V1VMKernelBoot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMKernelBoot) contextValidateContainer(ctx context.Context, formats strfmt.Registry) error {

	if m.Container != nil {

		if swag.IsZero(m.Container) { // not required
			return nil
		}

		if err := m.Container.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("container")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMKernelBoot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMKernelBoot) UnmarshalBinary(b []byte) error {
	var res V1VMKernelBoot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
