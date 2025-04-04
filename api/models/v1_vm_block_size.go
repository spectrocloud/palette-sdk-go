// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMBlockSize BlockSize provides the option to change the block size presented to the VM for a disk. Only one of its members may be specified.
//
// swagger:model v1VmBlockSize
type V1VMBlockSize struct {

	// custom
	Custom *V1VMCustomBlockSize `json:"custom,omitempty"`

	// match volume
	MatchVolume *V1VMFeatureState `json:"matchVolume,omitempty"`
}

// Validate validates this v1 Vm block size
func (m *V1VMBlockSize) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMBlockSize) validateCustom(formats strfmt.Registry) error {
	if swag.IsZero(m.Custom) { // not required
		return nil
	}

	if m.Custom != nil {
		if err := m.Custom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMBlockSize) validateMatchVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.MatchVolume) { // not required
		return nil
	}

	if m.MatchVolume != nil {
		if err := m.MatchVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matchVolume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("matchVolume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 Vm block size based on the context it is used
func (m *V1VMBlockSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMatchVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMBlockSize) contextValidateCustom(ctx context.Context, formats strfmt.Registry) error {

	if m.Custom != nil {

		if swag.IsZero(m.Custom) { // not required
			return nil
		}

		if err := m.Custom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMBlockSize) contextValidateMatchVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.MatchVolume != nil {

		if swag.IsZero(m.MatchVolume) { // not required
			return nil
		}

		if err := m.MatchVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matchVolume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("matchVolume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMBlockSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMBlockSize) UnmarshalBinary(b []byte) error {
	var res V1VMBlockSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
