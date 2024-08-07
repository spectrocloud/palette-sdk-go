// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VirtualInstanceType v1 virtual instance type
//
// swagger:model v1VirtualInstanceType
type V1VirtualInstanceType struct {

	// Maximum CPU cores
	MaxCPU int32 `json:"maxCPU,omitempty"`

	// Maximum memory in MiB
	MaxMemInMiB int32 `json:"maxMemInMiB,omitempty"`

	// Maximum storage in GiB
	MaxStorageGiB int32 `json:"maxStorageGiB,omitempty"`

	// Minimum CPU cores
	MinCPU int32 `json:"minCPU,omitempty"`

	// Minimum memory in MiB
	MinMemInMiB int32 `json:"minMemInMiB,omitempty"`

	// Minimum storage in GiB
	MinStorageGiB int32 `json:"minStorageGiB,omitempty"`
}

// Validate validates this v1 virtual instance type
func (m *V1VirtualInstanceType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VirtualInstanceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VirtualInstanceType) UnmarshalBinary(b []byte) error {
	var res V1VirtualInstanceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
