// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1WorkspaceBackupConfigEntity Cluster backup config
//
// swagger:model v1WorkspaceBackupConfigEntity
type V1WorkspaceBackupConfigEntity struct {

	// backup config
	BackupConfig *V1ClusterBackupConfig `json:"backupConfig,omitempty"`

	// cluster uids
	// Unique: true
	ClusterUids []string `json:"clusterUids"`

	// include all clusters
	IncludeAllClusters bool `json:"includeAllClusters,omitempty"`
}

// Validate validates this v1 workspace backup config entity
func (m *V1WorkspaceBackupConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterUids(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceBackupConfigEntity) validateBackupConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupConfig) { // not required
		return nil
	}

	if m.BackupConfig != nil {
		if err := m.BackupConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1WorkspaceBackupConfigEntity) validateClusterUids(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterUids) { // not required
		return nil
	}

	if err := validate.UniqueItems("clusterUids", "body", m.ClusterUids); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 workspace backup config entity based on the context it is used
func (m *V1WorkspaceBackupConfigEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceBackupConfigEntity) contextValidateBackupConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupConfig != nil {

		if swag.IsZero(m.BackupConfig) { // not required
			return nil
		}

		if err := m.BackupConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspaceBackupConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceBackupConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceBackupConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
