// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MaasZone Maas zone
//
// swagger:model v1MaasZone
type V1MaasZone struct {

	// Description of Maas domain
	Description string `json:"description,omitempty"`

	// Name of Maas zone
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 maas zone
func (m *V1MaasZone) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MaasZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MaasZone) UnmarshalBinary(b []byte) error {
	var res V1MaasZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}