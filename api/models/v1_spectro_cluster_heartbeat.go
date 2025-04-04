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

// V1SpectroClusterHeartbeat Cluster heartbeat message
//
// swagger:model v1SpectroClusterHeartbeat
type V1SpectroClusterHeartbeat struct {

	// Version of the agent
	// Required: true
	AgentVersion *string `json:"agentVersion"`

	// Heartbeat message
	Message string `json:"message,omitempty"`
}

// Validate validates this v1 spectro cluster heartbeat
func (m *V1SpectroClusterHeartbeat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterHeartbeat) validateAgentVersion(formats strfmt.Registry) error {

	if err := validate.Required("agentVersion", "body", m.AgentVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 spectro cluster heartbeat based on context it is used
func (m *V1SpectroClusterHeartbeat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterHeartbeat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterHeartbeat) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterHeartbeat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
