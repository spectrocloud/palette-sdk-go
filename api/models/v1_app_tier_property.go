// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AppTierProperty Application tier property object
//
// swagger:model v1AppTierProperty
type V1AppTierProperty struct {

	// Application tier property format
	Format string `json:"format,omitempty"`

	// Application tier property name
	Name string `json:"name,omitempty"`

	// Application tier property data type
	Type string `json:"type,omitempty"`

	// Application tier property value
	Value string `json:"value,omitempty"`
}

// Validate validates this v1 app tier property
func (m *V1AppTierProperty) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 app tier property based on context it is used
func (m *V1AppTierProperty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AppTierProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppTierProperty) UnmarshalBinary(b []byte) error {
	var res V1AppTierProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
