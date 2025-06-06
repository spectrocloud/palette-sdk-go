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

// V1VMDHCPPrivateOptions DHCPExtraOptions defines Extra DHCP options for a VM.
//
// swagger:model v1VmDHCPPrivateOptions
type V1VMDHCPPrivateOptions struct {

	// Option is an Integer value from 224-254 Required.
	// Required: true
	Option *int32 `json:"option"`

	// Value is a String value for the Option provided Required.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this v1 Vm d h c p private options
func (m *V1VMDHCPPrivateOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMDHCPPrivateOptions) validateOption(formats strfmt.Registry) error {

	if err := validate.Required("option", "body", m.Option); err != nil {
		return err
	}

	return nil
}

func (m *V1VMDHCPPrivateOptions) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 Vm d h c p private options based on context it is used
func (m *V1VMDHCPPrivateOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMDHCPPrivateOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMDHCPPrivateOptions) UnmarshalBinary(b []byte) error {
	var res V1VMDHCPPrivateOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
