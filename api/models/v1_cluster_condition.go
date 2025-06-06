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

// V1ClusterCondition v1 cluster condition
//
// swagger:model v1ClusterCondition
type V1ClusterCondition struct {

	// last probe time
	// Format: date-time
	LastProbeTime V1Time `json:"lastProbeTime,omitempty"`

	// last transition time
	// Format: date-time
	LastTransitionTime V1Time `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this v1 cluster condition
func (m *V1ClusterCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastProbeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterCondition) validateLastProbeTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastProbeTime) { // not required
		return nil
	}

	if err := m.LastProbeTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastProbeTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastProbeTime")
		}
		return err
	}

	return nil
}

func (m *V1ClusterCondition) validateLastTransitionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if err := m.LastTransitionTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTransitionTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastTransitionTime")
		}
		return err
	}

	return nil
}

func (m *V1ClusterCondition) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterCondition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 cluster condition based on the context it is used
func (m *V1ClusterCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastProbeTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastTransitionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterCondition) contextValidateLastProbeTime(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.LastProbeTime) { // not required
		return nil
	}

	if err := m.LastProbeTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastProbeTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastProbeTime")
		}
		return err
	}

	return nil
}

func (m *V1ClusterCondition) contextValidateLastTransitionTime(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if err := m.LastTransitionTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTransitionTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastTransitionTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterCondition) UnmarshalBinary(b []byte) error {
	var res V1ClusterCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
