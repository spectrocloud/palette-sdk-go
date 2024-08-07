// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1WorkspaceClusterBackupResponse Workspace cluster backup response
//
// swagger:model v1WorkspaceClusterBackupResponse
type V1WorkspaceClusterBackupResponse struct {

	// backup status meta
	BackupStatusMeta *V1BackupStatusMeta `json:"backupStatusMeta,omitempty"`

	// backup Uid
	BackupUID string `json:"backupUid,omitempty"`

	// cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// cluster Uid
	ClusterUID string `json:"clusterUid,omitempty"`
}

// Validate validates this v1 workspace cluster backup response
func (m *V1WorkspaceClusterBackupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupStatusMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceClusterBackupResponse) validateBackupStatusMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.BackupStatusMeta) { // not required
		return nil
	}

	if m.BackupStatusMeta != nil {
		if err := m.BackupStatusMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupStatusMeta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspaceClusterBackupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceClusterBackupResponse) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceClusterBackupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
