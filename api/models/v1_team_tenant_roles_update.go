// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TeamTenantRolesUpdate v1 team tenant roles update
//
// swagger:model v1TeamTenantRolesUpdate
type V1TeamTenantRolesUpdate struct {

	// roles
	Roles []string `json:"roles"`
}

// Validate validates this v1 team tenant roles update
func (m *V1TeamTenantRolesUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 team tenant roles update based on context it is used
func (m *V1TeamTenantRolesUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1TeamTenantRolesUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TeamTenantRolesUpdate) UnmarshalBinary(b []byte) error {
	var res V1TeamTenantRolesUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
