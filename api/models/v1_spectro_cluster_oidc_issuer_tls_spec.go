// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroClusterOidcIssuerTLSSpec v1 spectro cluster oidc issuer Tls spec
//
// swagger:model v1SpectroClusterOidcIssuerTlsSpec
type V1SpectroClusterOidcIssuerTLSSpec struct {

	// ca certificate base64
	CaCertificateBase64 string `json:"caCertificateBase64"`

	// insecure skip verify
	InsecureSkipVerify *bool `json:"insecureSkipVerify"`
}

// Validate validates this v1 spectro cluster oidc issuer Tls spec
func (m *V1SpectroClusterOidcIssuerTLSSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 spectro cluster oidc issuer Tls spec based on context it is used
func (m *V1SpectroClusterOidcIssuerTLSSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterOidcIssuerTLSSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterOidcIssuerTLSSpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterOidcIssuerTLSSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
