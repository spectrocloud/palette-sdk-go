// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroTenantClusterMigration Spectro tenant cluster migration status
//
// swagger:model v1SpectroTenantClusterMigration
type V1SpectroTenantClusterMigration struct {

	// message
	Message string `json:"message,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 spectro tenant cluster migration
func (m *V1SpectroTenantClusterMigration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroTenantClusterMigration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroTenantClusterMigration) UnmarshalBinary(b []byte) error {
	var res V1SpectroTenantClusterMigration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
