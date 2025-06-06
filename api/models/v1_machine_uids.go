// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MachineUids v1 machine uids
//
// swagger:model v1MachineUids
type V1MachineUids struct {

	// machine uids
	MachineUids []string `json:"machineUids"`
}

// Validate validates this v1 machine uids
func (m *V1MachineUids) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 machine uids based on context it is used
func (m *V1MachineUids) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineUids) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineUids) UnmarshalBinary(b []byte) error {
	var res V1MachineUids
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
