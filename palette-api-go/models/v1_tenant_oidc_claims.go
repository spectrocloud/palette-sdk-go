// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TenantOidcClaims v1 tenant oidc claims
//
// swagger:model v1TenantOidcClaims
type V1TenantOidcClaims struct {

	// email
	Email string `json:"Email"`

	// first name
	FirstName string `json:"FirstName"`

	// last name
	LastName string `json:"LastName"`

	// spectro team
	SpectroTeam string `json:"SpectroTeam"`
}

// Validate validates this v1 tenant oidc claims
func (m *V1TenantOidcClaims) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1TenantOidcClaims) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TenantOidcClaims) UnmarshalBinary(b []byte) error {
	var res V1TenantOidcClaims
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
