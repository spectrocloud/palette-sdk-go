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

// V1ClusterBackupConfig Cluster backup config
//
// swagger:model v1ClusterBackupConfig
type V1ClusterBackupConfig struct {

	// backup location name
	BackupLocationName string `json:"backupLocationName,omitempty"`

	// backup location Uid
	BackupLocationUID string `json:"backupLocationUid,omitempty"`

	// backup name
	BackupName string `json:"backupName,omitempty"`

	// backup prefix
	BackupPrefix string `json:"backupPrefix,omitempty"`

	// duration in hours
	DurationInHours int64 `json:"durationInHours,omitempty"`

	// include all disks
	IncludeAllDisks bool `json:"includeAllDisks,omitempty"`

	// include cluster resources
	IncludeClusterResources bool `json:"includeClusterResources,omitempty"`

	// location type
	LocationType string `json:"locationType,omitempty"`

	// namespaces
	// Unique: true
	Namespaces []string `json:"namespaces"`

	// schedule
	Schedule *V1ClusterFeatureSchedule `json:"schedule,omitempty"`
}

// Validate validates this v1 cluster backup config
func (m *V1ClusterBackupConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterBackupConfig) validateNamespaces(formats strfmt.Registry) error {

	if swag.IsZero(m.Namespaces) { // not required
		return nil
	}

	if err := validate.UniqueItems("namespaces", "body", m.Namespaces); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterBackupConfig) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterBackupConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterBackupConfig) UnmarshalBinary(b []byte) error {
	var res V1ClusterBackupConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}