// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MaasPool Maas pool
//
// swagger:model v1MaasPool
type V1MaasPool struct {

	// Description of Maas domain
	Description string `json:"description,omitempty"`

	// Name of Maas pool
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 maas pool
func (m *V1MaasPool) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MaasPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MaasPool) UnmarshalBinary(b []byte) error {
	var res V1MaasPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}