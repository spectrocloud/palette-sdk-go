// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1HybridMachinePoolStatus Get the hybrid machine pool's cluster status
//
// swagger:model v1HybridMachinePoolStatus
type V1HybridMachinePoolStatus struct {

	// Health of the hybrid machine pool
	Health *V1HybridMachinePoolClusterHealth `json:"health,omitempty"`

	// State of the hybrid machine pool
	State string `json:"state,omitempty"`
}

// Validate validates this v1 hybrid machine pool status
func (m *V1HybridMachinePoolStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1HybridMachinePoolStatus) validateHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 hybrid machine pool status based on the context it is used
func (m *V1HybridMachinePoolStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1HybridMachinePoolStatus) contextValidateHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.Health != nil {

		if swag.IsZero(m.Health) { // not required
			return nil
		}

		if err := m.Health.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1HybridMachinePoolStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1HybridMachinePoolStatus) UnmarshalBinary(b []byte) error {
	var res V1HybridMachinePoolStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
