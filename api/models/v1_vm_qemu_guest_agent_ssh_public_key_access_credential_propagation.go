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

// V1VMQemuGuestAgentSSHPublicKeyAccessCredentialPropagation v1 Vm qemu guest agent Ssh public key access credential propagation
//
// swagger:model v1VmQemuGuestAgentSshPublicKeyAccessCredentialPropagation
type V1VMQemuGuestAgentSSHPublicKeyAccessCredentialPropagation struct {

	// Users represents a list of guest users that should have the ssh public keys added to their authorized_keys file.
	// Required: true
	Users []string `json:"users"`
}

// Validate validates this v1 Vm qemu guest agent Ssh public key access credential propagation
func (m *V1VMQemuGuestAgentSSHPublicKeyAccessCredentialPropagation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMQemuGuestAgentSSHPublicKeyAccessCredentialPropagation) validateUsers(formats strfmt.Registry) error {

	if err := validate.Required("users", "body", m.Users); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 Vm qemu guest agent Ssh public key access credential propagation based on context it is used
func (m *V1VMQemuGuestAgentSSHPublicKeyAccessCredentialPropagation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMQemuGuestAgentSSHPublicKeyAccessCredentialPropagation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMQemuGuestAgentSSHPublicKeyAccessCredentialPropagation) UnmarshalBinary(b []byte) error {
	var res V1VMQemuGuestAgentSSHPublicKeyAccessCredentialPropagation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
