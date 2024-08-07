// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterWorkloadPodContainer Cluster workload pod container
//
// swagger:model v1ClusterWorkloadPodContainer
type V1ClusterWorkloadPodContainer struct {

	// image
	Image string `json:"image,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// resources
	Resources *V1ClusterWorkloadPodContainerResources `json:"resources,omitempty"`
}

// Validate validates this v1 cluster workload pod container
func (m *V1ClusterWorkloadPodContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterWorkloadPodContainer) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterWorkloadPodContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterWorkloadPodContainer) UnmarshalBinary(b []byte) error {
	var res V1ClusterWorkloadPodContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
