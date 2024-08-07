// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UserMeTenant v1 user me tenant
//
// swagger:model v1UserMeTenant
type V1UserMeTenant struct {

	// org name
	OrgName string `json:"orgName,omitempty"`

	// tenant Uid
	TenantUID string `json:"tenantUid,omitempty"`
}

// Validate validates this v1 user me tenant
func (m *V1UserMeTenant) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UserMeTenant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UserMeTenant) UnmarshalBinary(b []byte) error {
	var res V1UserMeTenant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
