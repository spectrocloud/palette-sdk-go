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

// V1VariableNames v1 variable names
//
// swagger:model v1VariableNames
type V1VariableNames struct {

	// Array of variable names
	// Required: true
	// Unique: true
	Variables []string `json:"variables"`
}

// Validate validates this v1 variable names
func (m *V1VariableNames) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VariableNames) validateVariables(formats strfmt.Registry) error {

	if err := validate.Required("variables", "body", m.Variables); err != nil {
		return err
	}

	if err := validate.UniqueItems("variables", "body", m.Variables); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VariableNames) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VariableNames) UnmarshalBinary(b []byte) error {
	var res V1VariableNames
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
