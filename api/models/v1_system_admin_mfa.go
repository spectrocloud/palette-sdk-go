// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SystemAdminMfa System Administrator MFA configuration
//
// swagger:model v1SystemAdminMfa
type V1SystemAdminMfa struct {

	// devices
	Devices []string `json:"devices"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 system admin mfa
func (m *V1SystemAdminMfa) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SystemAdminMfa) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SystemAdminMfa) UnmarshalBinary(b []byte) error {
	var res V1SystemAdminMfa
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
