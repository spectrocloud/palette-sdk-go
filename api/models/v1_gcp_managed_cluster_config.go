// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GcpManagedClusterConfig GCP managed cluster config
//
// swagger:model v1GcpManagedClusterConfig
type V1GcpManagedClusterConfig struct {

	// EnableAutopilot indicates whether to enable autopilot for this GKE cluster
	EnableAutoPilot bool `json:"enableAutoPilot,omitempty"`

	// Can be Region or Zone
	Location string `json:"location,omitempty"`
}

// Validate validates this v1 gcp managed cluster config
func (m *V1GcpManagedClusterConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 gcp managed cluster config based on context it is used
func (m *V1GcpManagedClusterConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GcpManagedClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GcpManagedClusterConfig) UnmarshalBinary(b []byte) error {
	var res V1GcpManagedClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
