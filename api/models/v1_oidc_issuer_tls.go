// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OidcIssuerTLS v1 oidc issuer Tls
//
// swagger:model v1OidcIssuerTls
type V1OidcIssuerTLS struct {

	// ca certificate base64
	CaCertificateBase64 string `json:"caCertificateBase64"`

	// insecure skip verify
	InsecureSkipVerify *bool `json:"insecureSkipVerify"`
}

// Validate validates this v1 oidc issuer Tls
func (m *V1OidcIssuerTLS) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1OidcIssuerTLS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OidcIssuerTLS) UnmarshalBinary(b []byte) error {
	var res V1OidcIssuerTLS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}