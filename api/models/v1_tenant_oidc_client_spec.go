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

// V1TenantOidcClientSpec Tenant
//
// swagger:model v1TenantOidcClientSpec
type V1TenantOidcClientSpec struct {

	// callback Url
	CallbackURL string `json:"callbackUrl"`

	// client Id
	ClientID string `json:"clientId"`

	// client secret
	ClientSecret string `json:"clientSecret"`

	// default teams
	DefaultTeams []string `json:"defaultTeams"`

	// is sso enabled
	IsSsoEnabled bool `json:"isSsoEnabled"`

	// issuer Tls
	IssuerTLS *V1OidcIssuerTLS `json:"issuerTls,omitempty"`

	// the issuer is the URL identifier for the service
	IssuerURL string `json:"issuerUrl"`

	// logout Url
	LogoutURL string `json:"logoutUrl"`

	// required claims
	RequiredClaims *V1TenantOidcClaims `json:"requiredClaims,omitempty"`

	// scopes
	Scopes []string `json:"scopes"`

	// scopes delimiter
	ScopesDelimiter string `json:"scopesDelimiter"`

	// When syncSsoTeams is set to true, all the teams from the OIDC configuration are pulled and saved in palette whereas when set to false, only the teams which are part of palette are pulled and saved
	SyncSsoTeams bool `json:"syncSsoTeams"`

	// user info
	UserInfo *V1OidcUserInfo `json:"userInfo,omitempty"`
}

// Validate validates this v1 tenant oidc client spec
func (m *V1TenantOidcClientSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIssuerTLS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredClaims(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TenantOidcClientSpec) validateIssuerTLS(formats strfmt.Registry) error {
	if swag.IsZero(m.IssuerTLS) { // not required
		return nil
	}

	if m.IssuerTLS != nil {
		if err := m.IssuerTLS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issuerTls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("issuerTls")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantOidcClientSpec) validateRequiredClaims(formats strfmt.Registry) error {
	if swag.IsZero(m.RequiredClaims) { // not required
		return nil
	}

	if m.RequiredClaims != nil {
		if err := m.RequiredClaims.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requiredClaims")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requiredClaims")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantOidcClientSpec) validateUserInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.UserInfo) { // not required
		return nil
	}

	if m.UserInfo != nil {
		if err := m.UserInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 tenant oidc client spec based on the context it is used
func (m *V1TenantOidcClientSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIssuerTLS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequiredClaims(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TenantOidcClientSpec) contextValidateIssuerTLS(ctx context.Context, formats strfmt.Registry) error {

	if m.IssuerTLS != nil {

		if swag.IsZero(m.IssuerTLS) { // not required
			return nil
		}

		if err := m.IssuerTLS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issuerTls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("issuerTls")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantOidcClientSpec) contextValidateRequiredClaims(ctx context.Context, formats strfmt.Registry) error {

	if m.RequiredClaims != nil {

		if swag.IsZero(m.RequiredClaims) { // not required
			return nil
		}

		if err := m.RequiredClaims.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requiredClaims")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requiredClaims")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantOidcClientSpec) contextValidateUserInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.UserInfo != nil {

		if swag.IsZero(m.UserInfo) { // not required
			return nil
		}

		if err := m.UserInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TenantOidcClientSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TenantOidcClientSpec) UnmarshalBinary(b []byte) error {
	var res V1TenantOidcClientSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
