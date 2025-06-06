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

// V1Organizations Returns a list of user's organizations details and login methods
//
// swagger:model v1Organizations
type V1Organizations struct {

	// Describes a list of user's organization
	// Unique: true
	Organizations []*V1Organization `json:"organizations"`
}

// Validate validates this v1 organizations
func (m *V1Organizations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Organizations) validateOrganizations(formats strfmt.Registry) error {
	if swag.IsZero(m.Organizations) { // not required
		return nil
	}

	if err := validate.UniqueItems("organizations", "body", m.Organizations); err != nil {
		return err
	}

	for i := 0; i < len(m.Organizations); i++ {
		if swag.IsZero(m.Organizations[i]) { // not required
			continue
		}

		if m.Organizations[i] != nil {
			if err := m.Organizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 organizations based on the context it is used
func (m *V1Organizations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrganizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Organizations) contextValidateOrganizations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Organizations); i++ {

		if m.Organizations[i] != nil {

			if swag.IsZero(m.Organizations[i]) { // not required
				return nil
			}

			if err := m.Organizations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Organizations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Organizations) UnmarshalBinary(b []byte) error {
	var res V1Organizations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
