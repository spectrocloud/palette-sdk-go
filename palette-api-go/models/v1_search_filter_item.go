// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SearchFilterItem v1 search filter item
//
// swagger:model v1SearchFilterItem
type V1SearchFilterItem struct {

	// condition
	Condition *V1SearchFilterCondition `json:"condition,omitempty"`

	// property
	Property string `json:"property,omitempty"`

	// type
	Type V1SearchFilterPropertyType `json:"type,omitempty"`
}

// Validate validates this v1 search filter item
func (m *V1SearchFilterItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchFilterItem) validateCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	if m.Condition != nil {
		if err := m.Condition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("condition")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterItem) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SearchFilterItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchFilterItem) UnmarshalBinary(b []byte) error {
	var res V1SearchFilterItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
