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

// V1OverlordVsphereAccountCreate v1 overlord vsphere account create
//
// swagger:model v1OverlordVsphereAccountCreate
type V1OverlordVsphereAccountCreate struct {

	// account
	Account *V1VsphereCloudAccount `json:"account,omitempty"`

	// Name for the private gateway & cloud account
	Name string `json:"name,omitempty"`

	// share with projects
	ShareWithProjects bool `json:"shareWithProjects"`
}

// Validate validates this v1 overlord vsphere account create
func (m *V1OverlordVsphereAccountCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OverlordVsphereAccountCreate) validateAccount(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 overlord vsphere account create based on the context it is used
func (m *V1OverlordVsphereAccountCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OverlordVsphereAccountCreate) contextValidateAccount(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1OverlordVsphereAccountCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OverlordVsphereAccountCreate) UnmarshalBinary(b []byte) error {
	var res V1OverlordVsphereAccountCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
