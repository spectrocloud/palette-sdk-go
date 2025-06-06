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

// V1DataSinkCloudWatchConfig Data sink cloud watch config
//
// swagger:model v1.DataSinkCloudWatchConfig
type V1DataSinkCloudWatchConfig struct {

	// payload
	Payload V1DataSinkPayloads `json:"payload,omitempty"`

	// spec
	Spec *V1CloudWatchConfig `json:"spec,omitempty"`
}

// Validate validates this v1 data sink cloud watch config
func (m *V1DataSinkCloudWatchConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DataSinkCloudWatchConfig) validatePayload(formats strfmt.Registry) error {
	if swag.IsZero(m.Payload) { // not required
		return nil
	}

	if err := m.Payload.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("payload")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("payload")
		}
		return err
	}

	return nil
}

func (m *V1DataSinkCloudWatchConfig) validateSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 data sink cloud watch config based on the context it is used
func (m *V1DataSinkCloudWatchConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePayload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DataSinkCloudWatchConfig) contextValidatePayload(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Payload.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("payload")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("payload")
		}
		return err
	}

	return nil
}

func (m *V1DataSinkCloudWatchConfig) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.Spec != nil {

		if swag.IsZero(m.Spec) { // not required
			return nil
		}

		if err := m.Spec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1DataSinkCloudWatchConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DataSinkCloudWatchConfig) UnmarshalBinary(b []byte) error {
	var res V1DataSinkCloudWatchConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
