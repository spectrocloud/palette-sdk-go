// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterGroupLimitConfig Cluster group limit config
//
// swagger:model v1ClusterGroupLimitConfig
type V1ClusterGroupLimitConfig struct {

	// Deprecated. Use field cpuMilliCore
	CPU int32 `json:"cpu,omitempty"`

	// CPU in milli cores
	CPUMilliCore int32 `json:"cpuMilliCore,omitempty"`

	// Deprecated. Use field memoryMiB
	Memory int32 `json:"memory,omitempty"`

	// Memory in MiB
	MemoryMiB int32 `json:"memoryMiB,omitempty"`

	// Over subscription percentage
	OverSubscription int32 `json:"overSubscription,omitempty"`

	// Storage in GiB
	StorageGiB int32 `json:"storageGiB,omitempty"`
}

// Validate validates this v1 cluster group limit config
func (m *V1ClusterGroupLimitConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterGroupLimitConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterGroupLimitConfig) UnmarshalBinary(b []byte) error {
	var res V1ClusterGroupLimitConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
