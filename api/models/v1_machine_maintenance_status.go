// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MachineMaintenanceStatus Machine maintenance status
//
// swagger:model v1MachineMaintenanceStatus
type V1MachineMaintenanceStatus struct {

	// action
	Action string `json:"action,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 machine maintenance status
func (m *V1MachineMaintenanceStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 machine maintenance status based on context it is used
func (m *V1MachineMaintenanceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineMaintenanceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineMaintenanceStatus) UnmarshalBinary(b []byte) error {
	var res V1MachineMaintenanceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
