// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PackResolvedValues Pack resolved values
//
// swagger:model v1PackResolvedValues
type V1PackResolvedValues struct {

	// Pack resolved values map
	Resolved map[string]string `json:"resolved,omitempty"`
}

// Validate validates this v1 pack resolved values
func (m *V1PackResolvedValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PackResolvedValues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackResolvedValues) UnmarshalBinary(b []byte) error {
	var res V1PackResolvedValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}