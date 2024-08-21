// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OverlordOpenStackAccountEntity v1 overlord open stack account entity
//
// swagger:model v1OverlordOpenStackAccountEntity
type V1OverlordOpenStackAccountEntity struct {

	// account
	Account *V1OpenStackCloudAccount `json:"account,omitempty"`

	// share with projects
	ShareWithProjects bool `json:"shareWithProjects"`
}

// Validate validates this v1 overlord open stack account entity
func (m *V1OverlordOpenStackAccountEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OverlordOpenStackAccountEntity) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OverlordOpenStackAccountEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OverlordOpenStackAccountEntity) UnmarshalBinary(b []byte) error {
	var res V1OverlordOpenStackAccountEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}