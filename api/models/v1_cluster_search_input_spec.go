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

// V1ClusterSearchInputSpec v1 cluster search input spec
//
// swagger:model v1ClusterSearchInputSpec
type V1ClusterSearchInputSpec struct {

	// inputs
	Inputs map[string]V1ClusterSearchInputSpecProperty `json:"inputs,omitempty"`
}

// Validate validates this v1 cluster search input spec
func (m *V1ClusterSearchInputSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterSearchInputSpec) validateInputs(formats strfmt.Registry) error {

	if swag.IsZero(m.Inputs) { // not required
		return nil
	}

	for k := range m.Inputs {

		if err := validate.Required("inputs"+"."+k, "body", m.Inputs[k]); err != nil {
			return err
		}
		if val, ok := m.Inputs[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterSearchInputSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterSearchInputSpec) UnmarshalBinary(b []byte) error {
	var res V1ClusterSearchInputSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}