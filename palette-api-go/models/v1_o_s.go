// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OS v1 o s
//
// swagger:model v1OS
type V1OS struct {

	// family
	Family string `json:"family,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 o s
func (m *V1OS) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1OS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OS) UnmarshalBinary(b []byte) error {
	var res V1OS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
