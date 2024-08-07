// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterComplianceScheduleConfig Cluster compliance scan schedule configuration
//
// swagger:model v1ClusterComplianceScheduleConfig
type V1ClusterComplianceScheduleConfig struct {

	// kube bench
	KubeBench *V1ClusterComplianceScanKubeBenchScheduleConfig `json:"kubeBench,omitempty"`

	// kube hunter
	KubeHunter *V1ClusterComplianceScanKubeHunterScheduleConfig `json:"kubeHunter,omitempty"`

	// sonobuoy
	Sonobuoy *V1ClusterComplianceScanSonobuoyScheduleConfig `json:"sonobuoy,omitempty"`
}

// Validate validates this v1 cluster compliance schedule config
func (m *V1ClusterComplianceScheduleConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKubeBench(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubeHunter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSonobuoy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterComplianceScheduleConfig) validateKubeBench(formats strfmt.Registry) error {

	if swag.IsZero(m.KubeBench) { // not required
		return nil
	}

	if m.KubeBench != nil {
		if err := m.KubeBench.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeBench")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterComplianceScheduleConfig) validateKubeHunter(formats strfmt.Registry) error {

	if swag.IsZero(m.KubeHunter) { // not required
		return nil
	}

	if m.KubeHunter != nil {
		if err := m.KubeHunter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeHunter")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterComplianceScheduleConfig) validateSonobuoy(formats strfmt.Registry) error {

	if swag.IsZero(m.Sonobuoy) { // not required
		return nil
	}

	if m.Sonobuoy != nil {
		if err := m.Sonobuoy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sonobuoy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterComplianceScheduleConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterComplianceScheduleConfig) UnmarshalBinary(b []byte) error {
	var res V1ClusterComplianceScheduleConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
