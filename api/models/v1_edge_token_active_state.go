// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EdgeTokenActiveState Edge token active state
//
// swagger:model v1EdgeTokenActiveState
type V1EdgeTokenActiveState struct {

	// Set to 'true', if the token is active
	IsActive bool `json:"isActive,omitempty"`
}

// Validate validates this v1 edge token active state
func (m *V1EdgeTokenActiveState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 edge token active state based on context it is used
func (m *V1EdgeTokenActiveState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeTokenActiveState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeTokenActiveState) UnmarshalBinary(b []byte) error {
	var res V1EdgeTokenActiveState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
