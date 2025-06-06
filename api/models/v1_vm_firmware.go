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

// V1VMFirmware v1 Vm firmware
//
// swagger:model v1VmFirmware
type V1VMFirmware struct {

	// bootloader
	Bootloader *V1VMBootloader `json:"bootloader,omitempty"`

	// kernel boot
	KernelBoot *V1VMKernelBoot `json:"kernelBoot,omitempty"`

	// The system-serial-number in SMBIOS
	Serial string `json:"serial,omitempty"`

	// UUID reported by the vmi bios. Defaults to a random generated uid.
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this v1 Vm firmware
func (m *V1VMFirmware) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootloader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKernelBoot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMFirmware) validateBootloader(formats strfmt.Registry) error {
	if swag.IsZero(m.Bootloader) { // not required
		return nil
	}

	if m.Bootloader != nil {
		if err := m.Bootloader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootloader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bootloader")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFirmware) validateKernelBoot(formats strfmt.Registry) error {
	if swag.IsZero(m.KernelBoot) { // not required
		return nil
	}

	if m.KernelBoot != nil {
		if err := m.KernelBoot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kernelBoot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kernelBoot")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 Vm firmware based on the context it is used
func (m *V1VMFirmware) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBootloader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKernelBoot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMFirmware) contextValidateBootloader(ctx context.Context, formats strfmt.Registry) error {

	if m.Bootloader != nil {

		if swag.IsZero(m.Bootloader) { // not required
			return nil
		}

		if err := m.Bootloader.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootloader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bootloader")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFirmware) contextValidateKernelBoot(ctx context.Context, formats strfmt.Registry) error {

	if m.KernelBoot != nil {

		if swag.IsZero(m.KernelBoot) { // not required
			return nil
		}

		if err := m.KernelBoot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kernelBoot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kernelBoot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMFirmware) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMFirmware) UnmarshalBinary(b []byte) error {
	var res V1VMFirmware
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
