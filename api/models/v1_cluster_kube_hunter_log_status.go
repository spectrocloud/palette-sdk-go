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

// V1ClusterKubeHunterLogStatus Cluster compliance scan KubeHunter Log Status
//
// swagger:model v1ClusterKubeHunterLogStatus
type V1ClusterKubeHunterLogStatus struct {

	// actor
	Actor *V1ClusterFeatureActor `json:"actor,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// reports
	Reports map[string]V1KubeHunterReport `json:"reports,omitempty"`

	// request Uid
	RequestUID string `json:"requestUid,omitempty"`

	// scan time
	ScanTime *V1ClusterScanTime `json:"scanTime,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 cluster kube hunter log status
func (m *V1ClusterKubeHunterLogStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterKubeHunterLogStatus) validateActor(formats strfmt.Registry) error {

	if swag.IsZero(m.Actor) { // not required
		return nil
	}

	if m.Actor != nil {
		if err := m.Actor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actor")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterKubeHunterLogStatus) validateReports(formats strfmt.Registry) error {

	if swag.IsZero(m.Reports) { // not required
		return nil
	}

	for k := range m.Reports {

		if err := validate.Required("reports"+"."+k, "body", m.Reports[k]); err != nil {
			return err
		}
		if val, ok := m.Reports[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterKubeHunterLogStatus) validateScanTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ScanTime) { // not required
		return nil
	}

	if m.ScanTime != nil {
		if err := m.ScanTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanTime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterKubeHunterLogStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterKubeHunterLogStatus) UnmarshalBinary(b []byte) error {
	var res V1ClusterKubeHunterLogStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
