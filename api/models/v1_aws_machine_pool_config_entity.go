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

// V1AwsMachinePoolConfigEntity v1 aws machine pool config entity
//
// swagger:model v1AwsMachinePoolConfigEntity
type V1AwsMachinePoolConfigEntity struct {

	// cloud config
	// Required: true
	CloudConfig *V1AwsMachinePoolCloudConfigEntity `json:"cloudConfig"`

	// pool config
	PoolConfig *V1MachinePoolConfigEntity `json:"poolConfig,omitempty"`
}

// Validate validates this v1 aws machine pool config entity
func (m *V1AwsMachinePoolConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoolConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsMachinePoolConfigEntity) validateCloudConfig(formats strfmt.Registry) error {

	if err := validate.Required("cloudConfig", "body", m.CloudConfig); err != nil {
		return err
	}

	if m.CloudConfig != nil {
		if err := m.CloudConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1AwsMachinePoolConfigEntity) validatePoolConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.PoolConfig) { // not required
		return nil
	}

	if m.PoolConfig != nil {
		if err := m.PoolConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("poolConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("poolConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 aws machine pool config entity based on the context it is used
func (m *V1AwsMachinePoolConfigEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePoolConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsMachinePoolConfigEntity) contextValidateCloudConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudConfig != nil {

		if err := m.CloudConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1AwsMachinePoolConfigEntity) contextValidatePoolConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.PoolConfig != nil {

		if swag.IsZero(m.PoolConfig) { // not required
			return nil
		}

		if err := m.PoolConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("poolConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("poolConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsMachinePoolConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsMachinePoolConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1AwsMachinePoolConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
