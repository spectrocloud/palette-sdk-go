// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EdgeHostIdentity v1 edge host identity
//
// swagger:model v1EdgeHostIdentity
type V1EdgeHostIdentity struct {

	// CACert is the client CA certificate
	CaCert string `json:"caCert,omitempty"`

	// Mode indicates a system or session connection to the host
	Mode string `json:"mode,omitempty"`

	// SocketPath is an optional path to the socket on the host, if not using defaults
	SocketPath string `json:"socketPath,omitempty"`

	// SSHSecret to the secret containing ssh-username
	SSHSecret *V1EdgeHostSSHSecret `json:"sshSecret,omitempty"`
}

// Validate validates this v1 edge host identity
func (m *V1EdgeHostIdentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSSHSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeHostIdentity) validateSSHSecret(formats strfmt.Registry) error {

	if swag.IsZero(m.SSHSecret) { // not required
		return nil
	}

	if m.SSHSecret != nil {
		if err := m.SSHSecret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshSecret")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeHostIdentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeHostIdentity) UnmarshalBinary(b []byte) error {
	var res V1EdgeHostIdentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}