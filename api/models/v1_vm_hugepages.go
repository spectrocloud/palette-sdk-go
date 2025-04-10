// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMHugepages Hugepages allow to use hugepages for the VirtualMachineInstance instead of regular memory.
//
// swagger:model v1VmHugepages
type V1VMHugepages struct {

	// PageSize specifies the hugepage size, for x86_64 architecture valid values are 1Gi and 2Mi.
	PageSize string `json:"pageSize,omitempty"`
}

// Validate validates this v1 Vm hugepages
func (m *V1VMHugepages) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 Vm hugepages based on context it is used
func (m *V1VMHugepages) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMHugepages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMHugepages) UnmarshalBinary(b []byte) error {
	var res V1VMHugepages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
