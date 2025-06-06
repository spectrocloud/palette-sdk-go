// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OverlordOpenStackAccountCreate v1 overlord open stack account create
//
// swagger:model v1OverlordOpenStackAccountCreate
type V1OverlordOpenStackAccountCreate struct {

	// account
	Account *V1OpenStackCloudAccount `json:"account,omitempty"`

	// Name for the private gateway & cloud account
	Name string `json:"name,omitempty"`

	// share with projects
	ShareWithProjects bool `json:"shareWithProjects"`
}

// Validate validates this v1 overlord open stack account create
func (m *V1OverlordOpenStackAccountCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OverlordOpenStackAccountCreate) validateAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 overlord open stack account create based on the context it is used
func (m *V1OverlordOpenStackAccountCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OverlordOpenStackAccountCreate) contextValidateAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.Account != nil {

		if swag.IsZero(m.Account) { // not required
			return nil
		}

		if err := m.Account.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OverlordOpenStackAccountCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OverlordOpenStackAccountCreate) UnmarshalBinary(b []byte) error {
	var res V1OverlordOpenStackAccountCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
