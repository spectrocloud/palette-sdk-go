// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1VMLabelSelectorRequirement A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
//
// swagger:model v1VmLabelSelectorRequirement
type V1VMLabelSelectorRequirement struct {

	// key is the label key that the selector applies to.
	// Required: true
	Key *string `json:"key"`

	// operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
	// Required: true
	Operator *string `json:"operator"`

	// values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Values []string `json:"values"`
}

// Validate validates this v1 Vm label selector requirement
func (m *V1VMLabelSelectorRequirement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
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

func (m *V1VMLabelSelectorRequirement) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *V1VMLabelSelectorRequirement) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMLabelSelectorRequirement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMLabelSelectorRequirement) UnmarshalBinary(b []byte) error {
	var res V1VMLabelSelectorRequirement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}