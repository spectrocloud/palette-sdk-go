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

// V1ClusterBackupStatus Cluster Backup Status
//
// swagger:model v1ClusterBackupStatus
type V1ClusterBackupStatus struct {

	// cluster backup statuses
	ClusterBackupStatuses []*V1ClusterBackupStatusMeta `json:"clusterBackupStatuses"`
}

// Validate validates this v1 cluster backup status
func (m *V1ClusterBackupStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterBackupStatuses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterBackupStatus) validateClusterBackupStatuses(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterBackupStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterBackupStatuses); i++ {
		if swag.IsZero(m.ClusterBackupStatuses[i]) { // not required
			continue
		}

		if m.ClusterBackupStatuses[i] != nil {
			if err := m.ClusterBackupStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterBackupStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterBackupStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterBackupStatus) UnmarshalBinary(b []byte) error {
	var res V1ClusterBackupStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
