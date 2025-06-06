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

// V1TenantSsoAuthProvidersEntity v1 tenant sso auth providers entity
//
// swagger:model v1TenantSsoAuthProvidersEntity
type V1TenantSsoAuthProvidersEntity struct {

	// is enabled
	IsEnabled bool `json:"isEnabled"`

	// sso logins
	// Unique: true
	SsoLogins []string `json:"ssoLogins"`
}

// Validate validates this v1 tenant sso auth providers entity
func (m *V1TenantSsoAuthProvidersEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSsoLogins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TenantSsoAuthProvidersEntity) validateSsoLogins(formats strfmt.Registry) error {
	if swag.IsZero(m.SsoLogins) { // not required
		return nil
	}

	if err := validate.UniqueItems("ssoLogins", "body", m.SsoLogins); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 tenant sso auth providers entity based on context it is used
func (m *V1TenantSsoAuthProvidersEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1TenantSsoAuthProvidersEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TenantSsoAuthProvidersEntity) UnmarshalBinary(b []byte) error {
	var res V1TenantSsoAuthProvidersEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
