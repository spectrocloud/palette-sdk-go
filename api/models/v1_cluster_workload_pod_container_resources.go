// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterWorkloadPodContainerResources Cluster workload pod container resources
//
// swagger:model v1ClusterWorkloadPodContainerResources
type V1ClusterWorkloadPodContainerResources struct {

	// limits
	Limits *V1ClusterWorkloadPodContainerResource `json:"limits,omitempty"`

	// requests
	Requests *V1ClusterWorkloadPodContainerResource `json:"requests,omitempty"`
}

// Validate validates this v1 cluster workload pod container resources
func (m *V1ClusterWorkloadPodContainerResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterWorkloadPodContainerResources) validateLimits(formats strfmt.Registry) error {

	if swag.IsZero(m.Limits) { // not required
		return nil
	}

	if m.Limits != nil {
		if err := m.Limits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limits")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterWorkloadPodContainerResources) validateRequests(formats strfmt.Registry) error {

	if swag.IsZero(m.Requests) { // not required
		return nil
	}

	if m.Requests != nil {
		if err := m.Requests.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requests")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterWorkloadPodContainerResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterWorkloadPodContainerResources) UnmarshalBinary(b []byte) error {
	var res V1ClusterWorkloadPodContainerResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}