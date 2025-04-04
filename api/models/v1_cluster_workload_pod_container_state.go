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

// V1ClusterWorkloadPodContainerState Cluster workload pod container state
//
// swagger:model v1ClusterWorkloadPodContainerState
type V1ClusterWorkloadPodContainerState struct {

	// exit code
	ExitCode int32 `json:"exitCode"`

	// finished at
	// Format: date-time
	FinishedAt V1Time `json:"finishedAt,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// started at
	// Format: date-time
	StartedAt V1Time `json:"startedAt,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 cluster workload pod container state
func (m *V1ClusterWorkloadPodContainerState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterWorkloadPodContainerState) validateFinishedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := m.FinishedAt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("finishedAt")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("finishedAt")
		}
		return err
	}

	return nil
}

func (m *V1ClusterWorkloadPodContainerState) validateStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := m.StartedAt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("startedAt")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("startedAt")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 cluster workload pod container state based on the context it is used
func (m *V1ClusterWorkloadPodContainerState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFinishedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterWorkloadPodContainerState) contextValidateFinishedAt(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := m.FinishedAt.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("finishedAt")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("finishedAt")
		}
		return err
	}

	return nil
}

func (m *V1ClusterWorkloadPodContainerState) contextValidateStartedAt(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := m.StartedAt.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("startedAt")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("startedAt")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterWorkloadPodContainerState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterWorkloadPodContainerState) UnmarshalBinary(b []byte) error {
	var res V1ClusterWorkloadPodContainerState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
