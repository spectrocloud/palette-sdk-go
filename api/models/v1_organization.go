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

// V1Organization Describes user's organization details
//
// swagger:model v1Organization
type V1Organization struct {

	// Describes user's enabled authorization mode
	AuthType string `json:"authType,omitempty"`

	// Describes user's organization name
	Name string `json:"name,omitempty"`

	// Describes user's organization authentication url
	RedirectURL string `json:"redirectUrl,omitempty"`

	// Describes a list of allowed social logins for the organization
	SsoLogins V1SsoLogins `json:"ssoLogins,omitempty"`
}

// Validate validates this v1 organization
func (m *V1Organization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSsoLogins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Organization) validateSsoLogins(formats strfmt.Registry) error {
	if swag.IsZero(m.SsoLogins) { // not required
		return nil
	}

	if err := m.SsoLogins.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ssoLogins")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ssoLogins")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 organization based on the context it is used
func (m *V1Organization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSsoLogins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Organization) contextValidateSsoLogins(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SsoLogins.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ssoLogins")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ssoLogins")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Organization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Organization) UnmarshalBinary(b []byte) error {
	var res V1Organization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
