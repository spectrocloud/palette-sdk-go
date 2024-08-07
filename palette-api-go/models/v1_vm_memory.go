// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMMemory Memory allows specifying the VirtualMachineInstance memory features.
//
// swagger:model v1VmMemory
type V1VMMemory struct {

	// guest
	Guest V1VMQuantity `json:"guest,omitempty"`

	// hugepages
	Hugepages *V1VMHugepages `json:"hugepages,omitempty"`
}

// Validate validates this v1 Vm memory
func (m *V1VMMemory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGuest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHugepages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMMemory) validateGuest(formats strfmt.Registry) error {

	if swag.IsZero(m.Guest) { // not required
		return nil
	}

	if err := m.Guest.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("guest")
		}
		return err
	}

	return nil
}

func (m *V1VMMemory) validateHugepages(formats strfmt.Registry) error {

	if swag.IsZero(m.Hugepages) { // not required
		return nil
	}

	if m.Hugepages != nil {
		if err := m.Hugepages.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hugepages")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMMemory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMMemory) UnmarshalBinary(b []byte) error {
	var res V1VMMemory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
