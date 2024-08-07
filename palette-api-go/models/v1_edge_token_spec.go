// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EdgeTokenSpec Edge token specification
//
// swagger:model v1EdgeTokenSpec
type V1EdgeTokenSpec struct {

	// Default project where the edgehost will be placed on the token authorization
	DefaultProject *V1EdgeTokenProject `json:"defaultProject,omitempty"`

	// Edge token expiry date
	// Format: date-time
	Expiry V1Time `json:"expiry,omitempty"`

	// Edge token
	Token string `json:"token,omitempty"`
}

// Validate validates this v1 edge token spec
func (m *V1EdgeTokenSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeTokenSpec) validateDefaultProject(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultProject) { // not required
		return nil
	}

	if m.DefaultProject != nil {
		if err := m.DefaultProject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultProject")
			}
			return err
		}
	}

	return nil
}

func (m *V1EdgeTokenSpec) validateExpiry(formats strfmt.Registry) error {

	if swag.IsZero(m.Expiry) { // not required
		return nil
	}

	if err := m.Expiry.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("expiry")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeTokenSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeTokenSpec) UnmarshalBinary(b []byte) error {
	var res V1EdgeTokenSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
