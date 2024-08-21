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

// V1WorkspaceClusterWorkloadPods Workspace cluster workload pods summary
//
// swagger:model v1WorkspaceClusterWorkloadPods
type V1WorkspaceClusterWorkloadPods struct {

	// metadata
	Metadata *V1RelatedObject `json:"metadata,omitempty"`

	// pods
	Pods []*V1ClusterWorkloadPod `json:"pods"`
}

// Validate validates this v1 workspace cluster workload pods
func (m *V1WorkspaceClusterWorkloadPods) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceClusterWorkloadPods) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1WorkspaceClusterWorkloadPods) validatePods(formats strfmt.Registry) error {

	if swag.IsZero(m.Pods) { // not required
		return nil
	}

	for i := 0; i < len(m.Pods); i++ {
		if swag.IsZero(m.Pods[i]) { // not required
			continue
		}

		if m.Pods[i] != nil {
			if err := m.Pods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspaceClusterWorkloadPods) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceClusterWorkloadPods) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceClusterWorkloadPods
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}