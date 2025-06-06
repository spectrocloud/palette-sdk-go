// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroEksClusterRateEntity Spectro EKS cluster request payload for estimating rate
//
// swagger:model v1SpectroEksClusterRateEntity
type V1SpectroEksClusterRateEntity struct {

	// cloud config
	CloudConfig *V1EksClusterConfig `json:"cloudConfig,omitempty"`

	// machinepoolconfig
	Machinepoolconfig []*V1EksMachinePoolConfigEntity `json:"machinepoolconfig"`
}

// Validate validates this v1 spectro eks cluster rate entity
func (m *V1SpectroEksClusterRateEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachinepoolconfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroEksClusterRateEntity) validateCloudConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudConfig) { // not required
		return nil
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

func (m *V1SpectroEksClusterRateEntity) validateMachinepoolconfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Machinepoolconfig) { // not required
		return nil
	}

	for i := 0; i < len(m.Machinepoolconfig); i++ {
		if swag.IsZero(m.Machinepoolconfig[i]) { // not required
			continue
		}

		if m.Machinepoolconfig[i] != nil {
			if err := m.Machinepoolconfig[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machinepoolconfig" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machinepoolconfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 spectro eks cluster rate entity based on the context it is used
func (m *V1SpectroEksClusterRateEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachinepoolconfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroEksClusterRateEntity) contextValidateCloudConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudConfig != nil {

		if swag.IsZero(m.CloudConfig) { // not required
			return nil
		}

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

func (m *V1SpectroEksClusterRateEntity) contextValidateMachinepoolconfig(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Machinepoolconfig); i++ {

		if m.Machinepoolconfig[i] != nil {

			if swag.IsZero(m.Machinepoolconfig[i]) { // not required
				return nil
			}

			if err := m.Machinepoolconfig[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machinepoolconfig" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machinepoolconfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroEksClusterRateEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroEksClusterRateEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroEksClusterRateEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
