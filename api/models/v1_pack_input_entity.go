// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PackInputEntity Pack request payload
//
// swagger:model v1PackInputEntity
type V1PackInputEntity struct {

	// pack
	Pack *V1PackManifestEntity `json:"pack,omitempty"`
}

// Validate validates this v1 pack input entity
func (m *V1PackInputEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackInputEntity) validatePack(formats strfmt.Registry) error {

	if swag.IsZero(m.Pack) { // not required
		return nil
	}

	if m.Pack != nil {
		if err := m.Pack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PackInputEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackInputEntity) UnmarshalBinary(b []byte) error {
	var res V1PackInputEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
