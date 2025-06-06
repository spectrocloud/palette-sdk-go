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

// V1ClusterWorkloadMetadata Cluster workload metadata
//
// swagger:model v1ClusterWorkloadMetadata
type V1ClusterWorkloadMetadata struct {

	// creation timestamp
	// Format: date-time
	CreationTimestamp V1Time `json:"creationTimestamp,omitempty"`

	// entity
	Entity *V1ClusterWorkloadRef `json:"entity,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this v1 cluster workload metadata
func (m *V1ClusterWorkloadMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterWorkloadMetadata) validateCreationTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := m.CreationTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("creationTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("creationTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1ClusterWorkloadMetadata) validateEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	if m.Entity != nil {
		if err := m.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 cluster workload metadata based on the context it is used
func (m *V1ClusterWorkloadMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreationTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterWorkloadMetadata) contextValidateCreationTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := m.CreationTimestamp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("creationTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("creationTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1ClusterWorkloadMetadata) contextValidateEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.Entity != nil {

		if swag.IsZero(m.Entity) { // not required
			return nil
		}

		if err := m.Entity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterWorkloadMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterWorkloadMetadata) UnmarshalBinary(b []byte) error {
	var res V1ClusterWorkloadMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
