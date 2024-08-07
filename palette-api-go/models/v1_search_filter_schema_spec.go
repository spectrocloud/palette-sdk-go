// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SearchFilterSchemaSpec v1 search filter schema spec
//
// swagger:model v1SearchFilterSchemaSpec
type V1SearchFilterSchemaSpec struct {

	// schema
	Schema *V1SearchFilterSchemaSpecProperties `json:"schema,omitempty"`
}

// Validate validates this v1 search filter schema spec
func (m *V1SearchFilterSchemaSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchFilterSchemaSpec) validateSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SearchFilterSchemaSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchFilterSchemaSpec) UnmarshalBinary(b []byte) error {
	var res V1SearchFilterSchemaSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
