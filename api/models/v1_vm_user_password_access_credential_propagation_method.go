// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMUserPasswordAccessCredentialPropagationMethod UserPasswordAccessCredentialPropagationMethod represents the method used to inject a user passwords into the vm guest. Only one of its members may be specified.
//
// swagger:model v1VmUserPasswordAccessCredentialPropagationMethod
type V1VMUserPasswordAccessCredentialPropagationMethod struct {

	// qemu guest agent
	QemuGuestAgent V1VMQemuGuestAgentUserPasswordAccessCredentialPropagation `json:"qemuGuestAgent,omitempty"`
}

// Validate validates this v1 Vm user password access credential propagation method
func (m *V1VMUserPasswordAccessCredentialPropagationMethod) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 Vm user password access credential propagation method based on context it is used
func (m *V1VMUserPasswordAccessCredentialPropagationMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMUserPasswordAccessCredentialPropagationMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMUserPasswordAccessCredentialPropagationMethod) UnmarshalBinary(b []byte) error {
	var res V1VMUserPasswordAccessCredentialPropagationMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
