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

// V1SpectroClusterProfileEntity Cluster profile request payload
//
// swagger:model v1SpectroClusterProfileEntity
type V1SpectroClusterProfileEntity struct {

	// Cluster profile packs array
	// Unique: true
	PackValues []*V1PackValuesEntity `json:"packValues"`

	// Cluster profile uid to be replaced with new profile
	ReplaceWithProfile string `json:"replaceWithProfile,omitempty"`

	// Cluster profile uid
	UID string `json:"uid,omitempty"`

	// variables
	Variables []*V1SpectroClusterVariable `json:"variables"`
}

// Validate validates this v1 spectro cluster profile entity
func (m *V1SpectroClusterProfileEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterProfileEntity) validatePackValues(formats strfmt.Registry) error {

	if swag.IsZero(m.PackValues) { // not required
		return nil
	}

	if err := validate.UniqueItems("packValues", "body", m.PackValues); err != nil {
		return err
	}

	for i := 0; i < len(m.PackValues); i++ {
		if swag.IsZero(m.PackValues[i]) { // not required
			continue
		}

		if m.PackValues[i] != nil {
			if err := m.PackValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterProfileEntity) validateVariables(formats strfmt.Registry) error {

	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for i := 0; i < len(m.Variables); i++ {
		if swag.IsZero(m.Variables[i]) { // not required
			continue
		}

		if m.Variables[i] != nil {
			if err := m.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterProfileEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterProfileEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterProfileEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
