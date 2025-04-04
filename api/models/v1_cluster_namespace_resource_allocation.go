// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ClusterNamespaceResourceAllocation Cluster namespace resource allocation
//
// swagger:model v1ClusterNamespaceResourceAllocation
type V1ClusterNamespaceResourceAllocation struct {

	// cpu cores
	// Minimum: > 0
	CPUCores float64 `json:"cpuCores,omitempty"`

	// memory mi b
	// Minimum: > 0
	MemoryMiB float64 `json:"memoryMiB,omitempty"`
}

// Validate validates this v1 cluster namespace resource allocation
func (m *V1ClusterNamespaceResourceAllocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryMiB(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterNamespaceResourceAllocation) validateCPUCores(formats strfmt.Registry) error {
	if swag.IsZero(m.CPUCores) { // not required
		return nil
	}

	if err := validate.Minimum("cpuCores", "body", m.CPUCores, 0, true); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterNamespaceResourceAllocation) validateMemoryMiB(formats strfmt.Registry) error {
	if swag.IsZero(m.MemoryMiB) { // not required
		return nil
	}

	if err := validate.Minimum("memoryMiB", "body", m.MemoryMiB, 0, true); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 cluster namespace resource allocation based on context it is used
func (m *V1ClusterNamespaceResourceAllocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterNamespaceResourceAllocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterNamespaceResourceAllocation) UnmarshalBinary(b []byte) error {
	var res V1ClusterNamespaceResourceAllocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
