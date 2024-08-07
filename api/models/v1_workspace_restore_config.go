// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1WorkspaceRestoreConfig Workspace cluster restore config
//
// swagger:model v1WorkspaceRestoreConfig
type V1WorkspaceRestoreConfig struct {

	// backup name
	// Required: true
	BackupName *string `json:"backupName"`

	// include cluster resources
	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	// include namespaces
	// Unique: true
	IncludeNamespaces []string `json:"includeNamespaces"`

	// preserve node ports
	PreserveNodePorts bool `json:"preserveNodePorts,omitempty"`

	// restore p vs
	RestorePVs bool `json:"restorePVs,omitempty"`

	// source cluster Uid
	// Required: true
	SourceClusterUID *string `json:"sourceClusterUid"`
}

// Validate validates this v1 workspace restore config
func (m *V1WorkspaceRestoreConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludeNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceClusterUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceRestoreConfig) validateBackupName(formats strfmt.Registry) error {

	if err := validate.Required("backupName", "body", m.BackupName); err != nil {
		return err
	}

	return nil
}

func (m *V1WorkspaceRestoreConfig) validateIncludeNamespaces(formats strfmt.Registry) error {

	if swag.IsZero(m.IncludeNamespaces) { // not required
		return nil
	}

	if err := validate.UniqueItems("includeNamespaces", "body", m.IncludeNamespaces); err != nil {
		return err
	}

	return nil
}

func (m *V1WorkspaceRestoreConfig) validateSourceClusterUID(formats strfmt.Registry) error {

	if err := validate.Required("sourceClusterUid", "body", m.SourceClusterUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspaceRestoreConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceRestoreConfig) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceRestoreConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
