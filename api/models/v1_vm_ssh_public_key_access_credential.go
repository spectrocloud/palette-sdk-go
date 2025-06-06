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

// V1VMSSHPublicKeyAccessCredential SSHPublicKeyAccessCredential represents a source and propagation method for injecting ssh public keys into a vm guest
//
// swagger:model v1VmSshPublicKeyAccessCredential
type V1VMSSHPublicKeyAccessCredential struct {

	// propagation method
	// Required: true
	PropagationMethod *V1VMSSHPublicKeyAccessCredentialPropagationMethod `json:"propagationMethod"`

	// source
	// Required: true
	Source *V1VMSSHPublicKeyAccessCredentialSource `json:"source"`
}

// Validate validates this v1 Vm Ssh public key access credential
func (m *V1VMSSHPublicKeyAccessCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePropagationMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMSSHPublicKeyAccessCredential) validatePropagationMethod(formats strfmt.Registry) error {

	if err := validate.Required("propagationMethod", "body", m.PropagationMethod); err != nil {
		return err
	}

	if m.PropagationMethod != nil {
		if err := m.PropagationMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("propagationMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("propagationMethod")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMSSHPublicKeyAccessCredential) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 Vm Ssh public key access credential based on the context it is used
func (m *V1VMSSHPublicKeyAccessCredential) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePropagationMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMSSHPublicKeyAccessCredential) contextValidatePropagationMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.PropagationMethod != nil {

		if err := m.PropagationMethod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("propagationMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("propagationMethod")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMSSHPublicKeyAccessCredential) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMSSHPublicKeyAccessCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMSSHPublicKeyAccessCredential) UnmarshalBinary(b []byte) error {
	var res V1VMSSHPublicKeyAccessCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
