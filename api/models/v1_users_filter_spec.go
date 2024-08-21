// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UsersFilterSpec Users filter spec
//
// swagger:model v1UsersFilterSpec
type V1UsersFilterSpec struct {

	// email Id
	EmailID *V1FilterString `json:"emailId,omitempty"`

	// name
	Name *V1FilterString `json:"name,omitempty"`
}

// Validate validates this v1 users filter spec
func (m *V1UsersFilterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UsersFilterSpec) validateEmailID(formats strfmt.Registry) error {

	if swag.IsZero(m.EmailID) { // not required
		return nil
	}

	if m.EmailID != nil {
		if err := m.EmailID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailId")
			}
			return err
		}
	}

	return nil
}

func (m *V1UsersFilterSpec) validateName(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V1UsersFilterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UsersFilterSpec) UnmarshalBinary(b []byte) error {
	var res V1UsersFilterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}