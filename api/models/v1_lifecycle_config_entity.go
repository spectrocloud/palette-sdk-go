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

// V1LifecycleConfigEntity v1 lifecycle config entity
//
// swagger:model v1LifecycleConfigEntity
type V1LifecycleConfigEntity struct {

	// lifecycle config
	LifecycleConfig *V1LifecycleConfig `json:"lifecycleConfig,omitempty"`
}

// Validate validates this v1 lifecycle config entity
func (m *V1LifecycleConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLifecycleConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1LifecycleConfigEntity) validateLifecycleConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.LifecycleConfig) { // not required
		return nil
	}

	if m.LifecycleConfig != nil {
		if err := m.LifecycleConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lifecycleConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lifecycleConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 lifecycle config entity based on the context it is used
func (m *V1LifecycleConfigEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLifecycleConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1LifecycleConfigEntity) contextValidateLifecycleConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.LifecycleConfig != nil {

		if swag.IsZero(m.LifecycleConfig) { // not required
			return nil
		}

		if err := m.LifecycleConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lifecycleConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lifecycleConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1LifecycleConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LifecycleConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1LifecycleConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
