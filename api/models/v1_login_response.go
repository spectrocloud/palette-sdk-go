// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1LoginResponse Returns the allowed login method and information with the organization details
//
// swagger:model v1LoginResponse
type V1LoginResponse struct {

	// Describes the env type. Possible values [ saas, self-hosted, quick-start, enterprise, airgap]
	AppEnv string `json:"appEnv,omitempty"`

	// Describes the default mode of authentication. Possible values [password, sso]
	// Enum: ["password","sso"]
	AuthType string `json:"authType,omitempty"`

	// Organization name.
	OrgName string `json:"orgName,omitempty"`

	// Describes the default redirect Url for authentication. If authType is sso, it will have tenant configured saml/oidc idp url else it will be users organization url
	RedirectURL string `json:"redirectUrl"`

	// Describes the domain url on which the saas is available
	RootDomain string `json:"rootDomain,omitempty"`

	// Describes which security mode is enabled
	SecurityMode string `json:"securityMode,omitempty"`

	// Just Inside. Describes the allowed social logins
	SsoLogins V1SsoLogins `json:"ssoLogins,omitempty"`

	// Describes the total number of tenant present in the system
	TotalTenants int64 `json:"totalTenants,omitempty"`
}

// Validate validates this v1 login response
func (m *V1LoginResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsoLogins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1LoginResponseTypeAuthTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["password","sso"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1LoginResponseTypeAuthTypePropEnum = append(v1LoginResponseTypeAuthTypePropEnum, v)
	}
}

const (

	// V1LoginResponseAuthTypePassword captures enum value "password"
	V1LoginResponseAuthTypePassword string = "password"

	// V1LoginResponseAuthTypeSso captures enum value "sso"
	V1LoginResponseAuthTypeSso string = "sso"
)

// prop value enum
func (m *V1LoginResponse) validateAuthTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1LoginResponseTypeAuthTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1LoginResponse) validateAuthType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthTypeEnum("authType", "body", m.AuthType); err != nil {
		return err
	}

	return nil
}

func (m *V1LoginResponse) validateSsoLogins(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 login response based on the context it is used
func (m *V1LoginResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSsoLogins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1LoginResponse) contextValidateSsoLogins(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1LoginResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LoginResponse) UnmarshalBinary(b []byte) error {
	var res V1LoginResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
