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

// V1PackRegistrySpec Pack registry credentials spec
//
// swagger:model v1PackRegistrySpec
type V1PackRegistrySpec struct {

	// auth
	// Required: true
	Auth *V1RegistryAuth `json:"auth"`

	// endpoint
	// Required: true
	Endpoint *string `json:"endpoint"`

	// name
	Name string `json:"name,omitempty"`

	// private
	Private bool `json:"private"`

	// Pack registry uid
	RegistryUID string `json:"registryUid,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`
}

// Validate validates this v1 pack registry spec
func (m *V1PackRegistrySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackRegistrySpec) validateAuth(formats strfmt.Registry) error {

	if err := validate.Required("auth", "body", m.Auth); err != nil {
		return err
	}

	if m.Auth != nil {
		if err := m.Auth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth")
			}
			return err
		}
	}

	return nil
}

func (m *V1PackRegistrySpec) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 pack registry spec based on the context it is used
func (m *V1PackRegistrySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackRegistrySpec) contextValidateAuth(ctx context.Context, formats strfmt.Registry) error {

	if m.Auth != nil {

		if err := m.Auth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PackRegistrySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackRegistrySpec) UnmarshalBinary(b []byte) error {
	var res V1PackRegistrySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
