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

// V1SearchFilterDateCondition v1 search filter date condition
//
// swagger:model v1SearchFilterDateCondition
type V1SearchFilterDateCondition struct {

	// match
	Match *V1SearchFilterDateConditionMatch `json:"match,omitempty"`

	// negation
	Negation bool `json:"negation,omitempty"`

	// operator
	Operator V1SearchFilterDateOperator `json:"operator,omitempty"`
}

// Validate validates this v1 search filter date condition
func (m *V1SearchFilterDateCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchFilterDateCondition) validateMatch(formats strfmt.Registry) error {
	if swag.IsZero(m.Match) { // not required
		return nil
	}

	if m.Match != nil {
		if err := m.Match.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("match")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("match")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterDateCondition) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	if err := m.Operator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operator")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operator")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 search filter date condition based on the context it is used
func (m *V1SearchFilterDateCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchFilterDateCondition) contextValidateMatch(ctx context.Context, formats strfmt.Registry) error {

	if m.Match != nil {

		if swag.IsZero(m.Match) { // not required
			return nil
		}

		if err := m.Match.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("match")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("match")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterDateCondition) contextValidateOperator(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	if err := m.Operator.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operator")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operator")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SearchFilterDateCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchFilterDateCondition) UnmarshalBinary(b []byte) error {
	var res V1SearchFilterDateCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
