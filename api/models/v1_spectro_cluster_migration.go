// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroClusterMigration Spectro cluster migration status
//
// swagger:model v1SpectroClusterMigration
type V1SpectroClusterMigration struct {

	// database
	Database *V1MgmtMigrationStatuses `json:"database,omitempty"`

	// state
	State string `json:"state"`

	// tenant
	Tenant *V1SpectroTenantMigration `json:"tenant,omitempty"`
}

// Validate validates this v1 spectro cluster migration
func (m *V1SpectroClusterMigration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterMigration) validateDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.Database) { // not required
		return nil
	}

	if m.Database != nil {
		if err := m.Database.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMigration) validateTenant(formats strfmt.Registry) error {

	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterMigration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterMigration) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterMigration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
