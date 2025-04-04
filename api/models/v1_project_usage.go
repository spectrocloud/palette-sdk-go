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

// V1ProjectUsage Project usage object
//
// swagger:model v1ProjectUsage
type V1ProjectUsage struct {

	// alloy
	Alloy *V1ProjectUsageData `json:"alloy,omitempty"`

	// pure
	Pure *V1ProjectUsageData `json:"pure,omitempty"`
}

// Validate validates this v1 project usage
func (m *V1ProjectUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlloy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePure(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectUsage) validateAlloy(formats strfmt.Registry) error {
	if swag.IsZero(m.Alloy) { // not required
		return nil
	}

	if m.Alloy != nil {
		if err := m.Alloy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alloy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alloy")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectUsage) validatePure(formats strfmt.Registry) error {
	if swag.IsZero(m.Pure) { // not required
		return nil
	}

	if m.Pure != nil {
		if err := m.Pure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pure")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 project usage based on the context it is used
func (m *V1ProjectUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlloy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectUsage) contextValidateAlloy(ctx context.Context, formats strfmt.Registry) error {

	if m.Alloy != nil {

		if swag.IsZero(m.Alloy) { // not required
			return nil
		}

		if err := m.Alloy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alloy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alloy")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectUsage) contextValidatePure(ctx context.Context, formats strfmt.Registry) error {

	if m.Pure != nil {

		if swag.IsZero(m.Pure) { // not required
			return nil
		}

		if err := m.Pure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pure")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectUsage) UnmarshalBinary(b []byte) error {
	var res V1ProjectUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
