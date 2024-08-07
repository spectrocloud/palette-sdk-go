// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Macros v1 macros
//
// swagger:model v1Macros
type V1Macros struct {

	// macros
	// Unique: true
	Macros []*V1Macro `json:"macros"`
}

// Validate validates this v1 macros
func (m *V1Macros) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMacros(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Macros) validateMacros(formats strfmt.Registry) error {

	if swag.IsZero(m.Macros) { // not required
		return nil
	}

	if err := validate.UniqueItems("macros", "body", m.Macros); err != nil {
		return err
	}

	for i := 0; i < len(m.Macros); i++ {
		if swag.IsZero(m.Macros[i]) { // not required
			continue
		}

		if m.Macros[i] != nil {
			if err := m.Macros[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("macros" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Macros) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Macros) UnmarshalBinary(b []byte) error {
	var res V1Macros
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
