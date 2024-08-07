// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroGenericClusterRateEntity Generic cluster request payload for estimating rate
//
// swagger:model v1SpectroGenericClusterRateEntity
type V1SpectroGenericClusterRateEntity struct {

	// cloud config
	CloudConfig *V1GenericClusterConfig `json:"cloudConfig,omitempty"`

	// machinepoolconfig
	Machinepoolconfig []*V1GenericMachinePoolConfigEntity `json:"machinepoolconfig"`
}

// Validate validates this v1 spectro generic cluster rate entity
func (m *V1SpectroGenericClusterRateEntity) Validate(formats strfmt.Registry) error {
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

func (m *V1SpectroGenericClusterRateEntity) validateCloudConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudConfig) { // not required
		return nil
	}

	if m.CloudConfig != nil {
		if err := m.CloudConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroGenericClusterRateEntity) validateMachinepoolconfig(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroGenericClusterRateEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroGenericClusterRateEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroGenericClusterRateEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
