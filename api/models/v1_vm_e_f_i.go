// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMEFI If set, EFI will be used instead of BIOS.
//
// swagger:model v1VmEFI
type V1VMEFI struct {

	// If set, SecureBoot will be enabled and the OVMF roms will be swapped for SecureBoot-enabled ones. Requires SMM to be enabled. Defaults to true
	SecureBoot bool `json:"secureBoot,omitempty"`
}

// Validate validates this v1 Vm e f i
func (m *V1VMEFI) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMEFI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMEFI) UnmarshalBinary(b []byte) error {
	var res V1VMEFI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}