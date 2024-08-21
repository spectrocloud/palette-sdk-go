// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMResourceRequirements v1 Vm resource requirements
//
// swagger:model v1VmResourceRequirements
type V1VMResourceRequirements struct {

	// Limits describes the maximum amount of compute resources allowed. Valid resource keys are "memory" and "cpu".
	Limits interface{} `json:"limits,omitempty"`

	// Don't ask the scheduler to take the guest-management overhead into account. Instead put the overhead only into the container's memory limit. This can lead to crashes if all memory is in use on a node. Defaults to false.
	OvercommitGuestOverhead bool `json:"overcommitGuestOverhead,omitempty"`

	// Requests is a description of the initial vmi resources. Valid resource keys are "memory" and "cpu".
	Requests interface{} `json:"requests,omitempty"`
}

// Validate validates this v1 Vm resource requirements
func (m *V1VMResourceRequirements) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMResourceRequirements) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMResourceRequirements) UnmarshalBinary(b []byte) error {
	var res V1VMResourceRequirements
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}