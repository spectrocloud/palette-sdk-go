// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UserRoleMap v1 user role map
//
// swagger:model v1UserRoleMap
type V1UserRoleMap struct {

	// roles
	Roles []string `json:"roles"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this v1 user role map
func (m *V1UserRoleMap) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UserRoleMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UserRoleMap) UnmarshalBinary(b []byte) error {
	var res V1UserRoleMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
