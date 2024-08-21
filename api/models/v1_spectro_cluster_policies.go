// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroClusterPolicies Cluster policies
//
// swagger:model v1SpectroClusterPolicies
type V1SpectroClusterPolicies struct {

	// backup policy
	BackupPolicy *V1ClusterBackupConfig `json:"backupPolicy,omitempty"`

	// scan policy
	ScanPolicy *V1ClusterComplianceScheduleConfig `json:"scanPolicy,omitempty"`
}

// Validate validates this v1 spectro cluster policies
func (m *V1SpectroClusterPolicies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterPolicies) validateBackupPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.BackupPolicy) { // not required
		return nil
	}

	if m.BackupPolicy != nil {
		if err := m.BackupPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterPolicies) validateScanPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ScanPolicy) { // not required
		return nil
	}

	if m.ScanPolicy != nil {
		if err := m.ScanPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterPolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterPolicies) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterPolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}