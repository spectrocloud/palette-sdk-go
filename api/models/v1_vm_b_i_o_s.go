// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMBIOS If set (default), BIOS will be used.
//
// swagger:model v1VmBIOS
type V1VMBIOS struct {

	// If set, the BIOS output will be transmitted over serial
	UseSerial bool `json:"useSerial,omitempty"`
}

// Validate validates this v1 Vm b i o s
func (m *V1VMBIOS) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 Vm b i o s based on context it is used
func (m *V1VMBIOS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMBIOS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMBIOS) UnmarshalBinary(b []byte) error {
	var res V1VMBIOS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
