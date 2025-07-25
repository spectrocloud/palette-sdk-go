// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ThemeMetadata Theme metadata
//
// swagger:model v1ThemeMetadata
type V1ThemeMetadata struct {

	// Tells if the theme is active or not
	Active bool `json:"active"`

	// Name of the theme
	Name string `json:"name,omitempty"`

	// Uid of the theme
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 theme metadata
func (m *V1ThemeMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ThemeMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ThemeMetadata) UnmarshalBinary(b []byte) error {
	var res V1ThemeMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
