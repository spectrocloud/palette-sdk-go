// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMSSHPublicKeyAccessCredentialPropagationMethod SSHPublicKeyAccessCredentialPropagationMethod represents the method used to inject a ssh public key into the vm guest. Only one of its members may be specified.
//
// swagger:model v1VmSshPublicKeyAccessCredentialPropagationMethod
type V1VMSSHPublicKeyAccessCredentialPropagationMethod struct {

	// config drive
	ConfigDrive V1VMConfigDriveSSHPublicKeyAccessCredentialPropagation `json:"configDrive,omitempty"`

	// qemu guest agent
	QemuGuestAgent *V1VMQemuGuestAgentSSHPublicKeyAccessCredentialPropagation `json:"qemuGuestAgent,omitempty"`
}

// Validate validates this v1 Vm Ssh public key access credential propagation method
func (m *V1VMSSHPublicKeyAccessCredentialPropagationMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQemuGuestAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMSSHPublicKeyAccessCredentialPropagationMethod) validateQemuGuestAgent(formats strfmt.Registry) error {

	if swag.IsZero(m.QemuGuestAgent) { // not required
		return nil
	}

	if m.QemuGuestAgent != nil {
		if err := m.QemuGuestAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qemuGuestAgent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMSSHPublicKeyAccessCredentialPropagationMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMSSHPublicKeyAccessCredentialPropagationMethod) UnmarshalBinary(b []byte) error {
	var res V1VMSSHPublicKeyAccessCredentialPropagationMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
