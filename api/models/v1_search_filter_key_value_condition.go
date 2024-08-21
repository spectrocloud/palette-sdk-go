// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SearchFilterKeyValueCondition v1 search filter key value condition
//
// swagger:model v1SearchFilterKeyValueCondition
type V1SearchFilterKeyValueCondition struct {

	// ignore case
	IgnoreCase bool `json:"ignoreCase,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// match
	Match *V1SearchFilterKeyValueConditionMatch `json:"match,omitempty"`

	// negation
	Negation bool `json:"negation,omitempty"`

	// operator
	Operator V1SearchFilterStringOperator `json:"operator,omitempty"`
}

// Validate validates this v1 search filter key value condition
func (m *V1SearchFilterKeyValueCondition) Validate(formats strfmt.Registry) error {
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

func (m *V1SearchFilterKeyValueCondition) validateMatch(formats strfmt.Registry) error {

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

func (m *V1SearchFilterKeyValueCondition) validateOperator(formats strfmt.Registry) error {

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
func (m *V1SearchFilterKeyValueCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchFilterKeyValueCondition) UnmarshalBinary(b []byte) error {
	var res V1SearchFilterKeyValueCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}