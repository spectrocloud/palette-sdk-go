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

// V1ClusterRestoreConfig Cluster restore config
//
// swagger:model v1ClusterRestoreConfig
type V1ClusterRestoreConfig struct {

	// backup name
	// Required: true
	BackupName *string `json:"backupName"`

	// backup request Uid
	// Required: true
	BackupRequestUID *string `json:"backupRequestUid"`

	// destination cluster Uid
	// Required: true
	DestinationClusterUID *string `json:"destinationClusterUid"`

	// include cluster resource mode
	IncludeClusterResourceMode V1IncludeClusterResourceMode `json:"includeClusterResourceMode,omitempty"`

	// Deprecated. Use includeClusterResourceMode
	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	// include namespaces
	// Unique: true
	IncludeNamespaces []string `json:"includeNamespaces"`

	// preserve node ports
	PreserveNodePorts bool `json:"preserveNodePorts,omitempty"`

	// restore p vs
	RestorePVs bool `json:"restorePVs,omitempty"`
}

// Validate validates this v1 cluster restore config
func (m *V1ClusterRestoreConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupRequestUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationClusterUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludeClusterResourceMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludeNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterRestoreConfig) validateBackupName(formats strfmt.Registry) error {

	if err := validate.Required("backupName", "body", m.BackupName); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterRestoreConfig) validateBackupRequestUID(formats strfmt.Registry) error {

	if err := validate.Required("backupRequestUid", "body", m.BackupRequestUID); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterRestoreConfig) validateDestinationClusterUID(formats strfmt.Registry) error {

	if err := validate.Required("destinationClusterUid", "body", m.DestinationClusterUID); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterRestoreConfig) validateIncludeClusterResourceMode(formats strfmt.Registry) error {

	if swag.IsZero(m.IncludeClusterResourceMode) { // not required
		return nil
	}

	if err := m.IncludeClusterResourceMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("includeClusterResourceMode")
		}
		return err
	}

	return nil
}

func (m *V1ClusterRestoreConfig) validateIncludeNamespaces(formats strfmt.Registry) error {

	if swag.IsZero(m.IncludeNamespaces) { // not required
		return nil
	}

	if err := validate.UniqueItems("includeNamespaces", "body", m.IncludeNamespaces); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterRestoreConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterRestoreConfig) UnmarshalBinary(b []byte) error {
	var res V1ClusterRestoreConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
