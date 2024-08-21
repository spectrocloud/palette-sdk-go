// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SearchFilterFloatCondition v1 search filter float condition
//
// swagger:model v1SearchFilterFloatCondition
type V1SearchFilterFloatCondition struct {

	// match
	Match *V1SearchFilterFloatConditionMatch `json:"match,omitempty"`

	// negation
	Negation bool `json:"negation,omitempty"`

	// operator
	Operator V1SearchFilterIntegerOperator `json:"operator,omitempty"`
}

// Validate validates this v1 search filter float condition
func (m *V1SearchFilterFloatCondition) Validate(formats strfmt.Registry) error {
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

func (m *V1SearchFilterFloatCondition) validateMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Match) { // not required
		return nil
	}

	if m.Match != nil {
		if err := m.Match.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("match")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterFloatCondition) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	if err := m.Operator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operator")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SearchFilterFloatCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchFilterFloatCondition) UnmarshalBinary(b []byte) error {
	var res V1SearchFilterFloatCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}