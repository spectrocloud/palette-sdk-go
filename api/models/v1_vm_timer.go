// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMTimer Represents all available timers in a vmi.
//
// swagger:model v1VmTimer
type V1VMTimer struct {

	// hpet
	Hpet *V1VMHPETTimer `json:"hpet,omitempty"`

	// hyperv
	Hyperv *V1VMHypervTimer `json:"hyperv,omitempty"`

	// kvm
	Kvm *V1VMKVMTimer `json:"kvm,omitempty"`

	// pit
	Pit *V1VMPITTimer `json:"pit,omitempty"`

	// rtc
	Rtc *V1VMRTCTimer `json:"rtc,omitempty"`
}

// Validate validates this v1 Vm timer
func (m *V1VMTimer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHpet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHyperv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRtc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMTimer) validateHpet(formats strfmt.Registry) error {

	if swag.IsZero(m.Hpet) { // not required
		return nil
	}

	if m.Hpet != nil {
		if err := m.Hpet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hpet")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMTimer) validateHyperv(formats strfmt.Registry) error {

	if swag.IsZero(m.Hyperv) { // not required
		return nil
	}

	if m.Hyperv != nil {
		if err := m.Hyperv.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hyperv")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMTimer) validateKvm(formats strfmt.Registry) error {

	if swag.IsZero(m.Kvm) { // not required
		return nil
	}

	if m.Kvm != nil {
		if err := m.Kvm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kvm")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMTimer) validatePit(formats strfmt.Registry) error {

	if swag.IsZero(m.Pit) { // not required
		return nil
	}

	if m.Pit != nil {
		if err := m.Pit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pit")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMTimer) validateRtc(formats strfmt.Registry) error {

	if swag.IsZero(m.Rtc) { // not required
		return nil
	}

	if m.Rtc != nil {
		if err := m.Rtc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rtc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMTimer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMTimer) UnmarshalBinary(b []byte) error {
	var res V1VMTimer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}