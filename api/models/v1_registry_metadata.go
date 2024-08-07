// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1RegistryMetadata Registry meta
//
// swagger:model v1RegistryMetadata
type V1RegistryMetadata struct {

	// is default
	IsDefault bool `json:"isDefault"`

	// is private
	IsPrivate bool `json:"isPrivate"`

	// kind
	Kind string `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 registry metadata
func (m *V1RegistryMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1RegistryMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RegistryMetadata) UnmarshalBinary(b []byte) error {
	var res V1RegistryMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
