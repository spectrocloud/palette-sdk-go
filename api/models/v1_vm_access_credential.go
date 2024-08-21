// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMAccessCredential AccessCredential represents a credential source that can be used to authorize remote access to the vm guest Only one of its members may be specified.
//
// swagger:model v1VmAccessCredential
type V1VMAccessCredential struct {

	// ssh public key
	SSHPublicKey *V1VMSSHPublicKeyAccessCredential `json:"sshPublicKey,omitempty"`

	// user password
	UserPassword *V1VMUserPasswordAccessCredential `json:"userPassword,omitempty"`
}

// Validate validates this v1 Vm access credential
func (m *V1VMAccessCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSSHPublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserPassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMAccessCredential) validateSSHPublicKey(formats strfmt.Registry) error {

	if swag.IsZero(m.SSHPublicKey) { // not required
		return nil
	}

	if m.SSHPublicKey != nil {
		if err := m.SSHPublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshPublicKey")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMAccessCredential) validateUserPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.UserPassword) { // not required
		return nil
	}

	if m.UserPassword != nil {
		if err := m.UserPassword.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userPassword")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMAccessCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMAccessCredential) UnmarshalBinary(b []byte) error {
	var res V1VMAccessCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}