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

// V1TenantResourceLimitEntity v1 tenant resource limit entity
//
// swagger:model v1TenantResourceLimitEntity
type V1TenantResourceLimitEntity struct {

	// kind
	Kind V1ResourceLimitType `json:"kind,omitempty"`

	// limit
	Limit int64 `json:"limit"`
}

// Validate validates this v1 tenant resource limit entity
func (m *V1TenantResourceLimitEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TenantResourceLimitEntity) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if err := m.Kind.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("kind")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("kind")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 tenant resource limit entity based on the context it is used
func (m *V1TenantResourceLimitEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TenantResourceLimitEntity) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if err := m.Kind.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("kind")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("kind")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TenantResourceLimitEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TenantResourceLimitEntity) UnmarshalBinary(b []byte) error {
	var res V1TenantResourceLimitEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
