// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterComplianceScanKubeHunterScheduleConfig Cluster compliance scan schedule config for kube hunter driver
//
// swagger:model v1ClusterComplianceScanKubeHunterScheduleConfig
type V1ClusterComplianceScanKubeHunterScheduleConfig struct {

	// schedule
	Schedule *V1ClusterFeatureSchedule `json:"schedule,omitempty"`
}

// Validate validates this v1 cluster compliance scan kube hunter schedule config
func (m *V1ClusterComplianceScanKubeHunterScheduleConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterComplianceScanKubeHunterScheduleConfig) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 cluster compliance scan kube hunter schedule config based on the context it is used
func (m *V1ClusterComplianceScanKubeHunterScheduleConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterComplianceScanKubeHunterScheduleConfig) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {

		if swag.IsZero(m.Schedule) { // not required
			return nil
		}

		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterComplianceScanKubeHunterScheduleConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterComplianceScanKubeHunterScheduleConfig) UnmarshalBinary(b []byte) error {
	var res V1ClusterComplianceScanKubeHunterScheduleConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
