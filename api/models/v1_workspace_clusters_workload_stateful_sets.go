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

// V1WorkspaceClustersWorkloadStatefulSets Workspace clusters workload statefulsets summary
//
// swagger:model v1WorkspaceClustersWorkloadStatefulSets
type V1WorkspaceClustersWorkloadStatefulSets struct {

	// clusters
	Clusters []*V1WorkspaceClusterWorkloadStatefulSets `json:"clusters"`

	// metadata
	Metadata *V1ObjectMetaInputEntity `json:"metadata,omitempty"`
}

// Validate validates this v1 workspace clusters workload stateful sets
func (m *V1WorkspaceClustersWorkloadStatefulSets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceClustersWorkloadStatefulSets) validateClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1WorkspaceClustersWorkloadStatefulSets) validateMetadata(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V1WorkspaceClustersWorkloadStatefulSets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceClustersWorkloadStatefulSets) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceClustersWorkloadStatefulSets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
