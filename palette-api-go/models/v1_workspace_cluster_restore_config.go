// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1WorkspaceClusterRestoreConfig Workspace cluster restore config
//
// swagger:model v1WorkspaceClusterRestoreConfig
type V1WorkspaceClusterRestoreConfig struct {

	// backup name
	BackupName string `json:"backupName,omitempty"`

	// cluster restore refs
	ClusterRestoreRefs []*V1WorkspaceClusterRestoreResponse `json:"clusterRestoreRefs"`

	// restore state
	RestoreState *V1WorkspaceRestoreState `json:"restoreState,omitempty"`

	// restore time
	// Format: date-time
	RestoreTime V1Time `json:"restoreTime,omitempty"`
}

// Validate validates this v1 workspace cluster restore config
func (m *V1WorkspaceClusterRestoreConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterRestoreRefs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoreState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoreTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceClusterRestoreConfig) validateClusterRestoreRefs(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterRestoreRefs) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterRestoreRefs); i++ {
		if swag.IsZero(m.ClusterRestoreRefs[i]) { // not required
			continue
		}

		if m.ClusterRestoreRefs[i] != nil {
			if err := m.ClusterRestoreRefs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterRestoreRefs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1WorkspaceClusterRestoreConfig) validateRestoreState(formats strfmt.Registry) error {

	if swag.IsZero(m.RestoreState) { // not required
		return nil
	}

	if m.RestoreState != nil {
		if err := m.RestoreState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoreState")
			}
			return err
		}
	}

	return nil
}

func (m *V1WorkspaceClusterRestoreConfig) validateRestoreTime(formats strfmt.Registry) error {

	if swag.IsZero(m.RestoreTime) { // not required
		return nil
	}

	if err := m.RestoreTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("restoreTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspaceClusterRestoreConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceClusterRestoreConfig) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceClusterRestoreConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
