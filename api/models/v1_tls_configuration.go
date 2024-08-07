// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TLSConfiguration TLS configuration
//
// swagger:model v1TlsConfiguration
type V1TLSConfiguration struct {

	// ca
	Ca string `json:"ca,omitempty"`

	// certificate
	Certificate string `json:"certificate,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// insecure skip verify
	InsecureSkipVerify bool `json:"insecureSkipVerify"`

	// key
	Key string `json:"key,omitempty"`
}

// Validate validates this v1 Tls configuration
func (m *V1TLSConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1TLSConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TLSConfiguration) UnmarshalBinary(b []byte) error {
	var res V1TLSConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
