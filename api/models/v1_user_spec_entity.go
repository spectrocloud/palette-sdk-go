// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1UserSpecEntity User Entity input
//
// swagger:model v1UserSpecEntity
type V1UserSpecEntity struct {

	// email Id
	EmailID string `json:"emailId,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// login mode
	LoginMode string `json:"loginMode,omitempty"`

	// roles
	// Unique: true
	Roles []string `json:"roles"`

	// teams
	// Unique: true
	Teams []string `json:"teams"`
}

// Validate validates this v1 user spec entity
func (m *V1UserSpecEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UserSpecEntity) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	if err := validate.UniqueItems("roles", "body", m.Roles); err != nil {
		return err
	}

	return nil
}

func (m *V1UserSpecEntity) validateTeams(formats strfmt.Registry) error {
	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	if err := validate.UniqueItems("teams", "body", m.Teams); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 user spec entity based on context it is used
func (m *V1UserSpecEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UserSpecEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UserSpecEntity) UnmarshalBinary(b []byte) error {
	var res V1UserSpecEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
