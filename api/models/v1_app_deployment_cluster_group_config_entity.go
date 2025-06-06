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

// V1AppDeploymentClusterGroupConfigEntity Application deployment cluster group config
//
// swagger:model v1AppDeploymentClusterGroupConfigEntity
type V1AppDeploymentClusterGroupConfigEntity struct {

	// target spec
	TargetSpec *V1AppDeploymentClusterGroupTargetSpec `json:"targetSpec,omitempty"`
}

// Validate validates this v1 app deployment cluster group config entity
func (m *V1AppDeploymentClusterGroupConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppDeploymentClusterGroupConfigEntity) validateTargetSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetSpec) { // not required
		return nil
	}

	if m.TargetSpec != nil {
		if err := m.TargetSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetSpec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 app deployment cluster group config entity based on the context it is used
func (m *V1AppDeploymentClusterGroupConfigEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTargetSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppDeploymentClusterGroupConfigEntity) contextValidateTargetSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetSpec != nil {

		if swag.IsZero(m.TargetSpec) { // not required
			return nil
		}

		if err := m.TargetSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppDeploymentClusterGroupConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppDeploymentClusterGroupConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1AppDeploymentClusterGroupConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
