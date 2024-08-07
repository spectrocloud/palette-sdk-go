// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GenericInstanceType v1 generic instance type
//
// swagger:model v1GenericInstanceType
type V1GenericInstanceType struct {

	// DiskGiB is the size of a virtual machine's disk, in GiB
	DiskGiB int32 `json:"diskGiB,omitempty"`

	// MemoryMiB is the size of a virtual machine's memory, in MiB
	MemoryMiB int64 `json:"memoryMiB,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// NumCPUs is the number of virtual processors in a virtual machine
	NumCPUs int32 `json:"numCPUs,omitempty"`
}

// Validate validates this v1 generic instance type
func (m *V1GenericInstanceType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GenericInstanceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GenericInstanceType) UnmarshalBinary(b []byte) error {
	var res V1GenericInstanceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
