// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ProjectTeamsEntity v1 project teams entity
//
// swagger:model v1ProjectTeamsEntity
type V1ProjectTeamsEntity struct {

	// teams
	// Unique: true
	Teams []*V1TeamRoleMap `json:"teams"`
}

// Validate validates this v1 project teams entity
func (m *V1ProjectTeamsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectTeamsEntity) validateTeams(formats strfmt.Registry) error {
	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	if err := validate.UniqueItems("teams", "body", m.Teams); err != nil {
		return err
	}

	for i := 0; i < len(m.Teams); i++ {
		if swag.IsZero(m.Teams[i]) { // not required
			continue
		}

		if m.Teams[i] != nil {
			if err := m.Teams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 project teams entity based on the context it is used
func (m *V1ProjectTeamsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTeams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectTeamsEntity) contextValidateTeams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Teams); i++ {

		if m.Teams[i] != nil {

			if swag.IsZero(m.Teams[i]) { // not required
				return nil
			}

			if err := m.Teams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectTeamsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectTeamsEntity) UnmarshalBinary(b []byte) error {
	var res V1ProjectTeamsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
