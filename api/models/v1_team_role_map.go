// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TeamRoleMap v1 team role map
//
// swagger:model v1TeamRoleMap
type V1TeamRoleMap struct {

	// roles
	Roles []string `json:"roles"`

	// team Id
	TeamID string `json:"teamId,omitempty"`
}

// Validate validates this v1 team role map
func (m *V1TeamRoleMap) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1TeamRoleMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TeamRoleMap) UnmarshalBinary(b []byte) error {
	var res V1TeamRoleMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}