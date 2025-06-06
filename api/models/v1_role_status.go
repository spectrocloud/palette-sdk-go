// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1RoleStatus Role status
//
// swagger:model v1RoleStatus
type V1RoleStatus struct {

	// Specifies if role account is enabled/disabled
	IsEnabled bool `json:"isEnabled"`
}

// Validate validates this v1 role status
func (m *V1RoleStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 role status based on context it is used
func (m *V1RoleStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1RoleStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RoleStatus) UnmarshalBinary(b []byte) error {
	var res V1RoleStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
