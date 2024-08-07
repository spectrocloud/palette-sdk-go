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

// V1EdgeHostsMetadataFilterSpec Edge hosts metadata filter spec
//
// swagger:model v1EdgeHostsMetadataFilterSpec
type V1EdgeHostsMetadataFilterSpec struct {

	// name
	Name *V1FilterString `json:"name,omitempty"`

	// states
	// Unique: true
	States []V1EdgeHostState `json:"states"`
}

// Validate validates this v1 edge hosts metadata filter spec
func (m *V1EdgeHostsMetadataFilterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeHostsMetadataFilterSpec) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

func (m *V1EdgeHostsMetadataFilterSpec) validateStates(formats strfmt.Registry) error {

	if swag.IsZero(m.States) { // not required
		return nil
	}

	if err := validate.UniqueItems("states", "body", m.States); err != nil {
		return err
	}

	for i := 0; i < len(m.States); i++ {

		if err := m.States[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("states" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeHostsMetadataFilterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeHostsMetadataFilterSpec) UnmarshalBinary(b []byte) error {
	var res V1EdgeHostsMetadataFilterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
