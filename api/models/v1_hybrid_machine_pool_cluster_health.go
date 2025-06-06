// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1HybridMachinePoolClusterHealth Machine pool cluster meta health information
//
// swagger:model v1HybridMachinePoolClusterHealth
type V1HybridMachinePoolClusterHealth struct {

	// is heart beat failed
	IsHeartBeatFailed bool `json:"isHeartBeatFailed"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 hybrid machine pool cluster health
func (m *V1HybridMachinePoolClusterHealth) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 hybrid machine pool cluster health based on context it is used
func (m *V1HybridMachinePoolClusterHealth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1HybridMachinePoolClusterHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1HybridMachinePoolClusterHealth) UnmarshalBinary(b []byte) error {
	var res V1HybridMachinePoolClusterHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
