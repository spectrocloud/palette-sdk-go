// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PairingCode Pairing code response
//
// swagger:model v1PairingCode
type V1PairingCode struct {

	// pairing code
	PairingCode string `json:"pairingCode,omitempty"`
}

// Validate validates this v1 pairing code
func (m *V1PairingCode) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PairingCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PairingCode) UnmarshalBinary(b []byte) error {
	var res V1PairingCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}