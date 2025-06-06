// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterWorkloadPodContainerResource Cluster workload pod container resource
//
// swagger:model v1ClusterWorkloadPodContainerResource
type V1ClusterWorkloadPodContainerResource struct {

	// cpu
	CPU int32 `json:"cpu"`

	// cpu unit
	CPUUnit string `json:"cpuUnit,omitempty"`

	// memory
	Memory int64 `json:"memory"`

	// memory unit
	MemoryUnit string `json:"memoryUnit,omitempty"`
}

// Validate validates this v1 cluster workload pod container resource
func (m *V1ClusterWorkloadPodContainerResource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 cluster workload pod container resource based on context it is used
func (m *V1ClusterWorkloadPodContainerResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterWorkloadPodContainerResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterWorkloadPodContainerResource) UnmarshalBinary(b []byte) error {
	var res V1ClusterWorkloadPodContainerResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
