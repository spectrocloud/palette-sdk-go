// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TunnelEndpoint v1 tunnel endpoint
//
// swagger:model v1TunnelEndpoint
type V1TunnelEndpoint struct {

	// Describes the URL where the client has to connect to the tunnel server
	Endpoint string `json:"endpoint,omitempty"`

	// tls
	TLS *V1TunnelEndpointTLS `json:"tls,omitempty"`
}

// Validate validates this v1 tunnel endpoint
func (m *V1TunnelEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTLS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TunnelEndpoint) validateTLS(formats strfmt.Registry) error {

	if swag.IsZero(m.TLS) { // not required
		return nil
	}

	if m.TLS != nil {
		if err := m.TLS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TunnelEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TunnelEndpoint) UnmarshalBinary(b []byte) error {
	var res V1TunnelEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1TunnelEndpointTLS Describes the Tunnel tls config which client will use to make a request to the tunnel server
//
// swagger:model V1TunnelEndpointTLS
type V1TunnelEndpointTLS struct {

	// ca cert
	CaCert string `json:"caCert,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// insecure skip verify
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
}

// Validate validates this v1 tunnel endpoint TLS
func (m *V1TunnelEndpointTLS) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1TunnelEndpointTLS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TunnelEndpointTLS) UnmarshalBinary(b []byte) error {
	var res V1TunnelEndpointTLS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
