// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MfaAuthenticatorSelection v1 mfa authenticator selection
//
// swagger:model v1MfaAuthenticatorSelection
type V1MfaAuthenticatorSelection struct {

	// authenticator attachment
	AuthenticatorAttachment string `json:"authenticatorAttachment,omitempty"`

	// require resident key
	RequireResidentKey bool `json:"requireResidentKey,omitempty"`

	// resident key
	ResidentKey string `json:"residentKey,omitempty"`

	// user verification
	UserVerification string `json:"userVerification,omitempty"`
}

// Validate validates this v1 mfa authenticator selection
func (m *V1MfaAuthenticatorSelection) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MfaAuthenticatorSelection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MfaAuthenticatorSelection) UnmarshalBinary(b []byte) error {
	var res V1MfaAuthenticatorSelection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
