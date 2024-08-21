// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PackDependency Pack template dependency
//
// swagger:model v1PackDependency
type V1PackDependency struct {

	// Pack template dependency pack layer
	Layer string `json:"layer,omitempty"`

	// Pack template dependency pack name
	Name string `json:"name,omitempty"`

	// If true then dependency pack values can't be overridden
	ReadOnly bool `json:"readOnly,omitempty"`
}

// Validate validates this v1 pack dependency
func (m *V1PackDependency) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PackDependency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackDependency) UnmarshalBinary(b []byte) error {
	var res V1PackDependency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}