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

// V1TagFilterItem v1 tag filter item
//
// swagger:model v1TagFilterItem
type V1TagFilterItem struct {

	// key
	Key string `json:"key,omitempty"`

	// negation
	Negation bool `json:"negation,omitempty"`

	// operator
	Operator V1SearchFilterKeyValueOperator `json:"operator,omitempty"`

	// values
	// Unique: true
	Values []string `json:"values"`
}

// Validate validates this v1 tag filter item
func (m *V1TagFilterItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TagFilterItem) validateOperator(formats strfmt.Registry) error {
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

func (m *V1TagFilterItem) validateValues(formats strfmt.Registry) error {
	if swag.IsZero(m.Values) { // not required
		return nil
	}

	if err := validate.UniqueItems("values", "body", m.Values); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 tag filter item based on the context it is used
func (m *V1TagFilterItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TagFilterItem) contextValidateOperator(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1TagFilterItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TagFilterItem) UnmarshalBinary(b []byte) error {
	var res V1TagFilterItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
