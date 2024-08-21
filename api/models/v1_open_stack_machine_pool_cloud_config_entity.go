// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1OpenStackMachinePoolCloudConfigEntity v1 open stack machine pool cloud config entity
//
// swagger:model v1OpenStackMachinePoolCloudConfigEntity
type V1OpenStackMachinePoolCloudConfigEntity struct {

	// for control plane pool, this will be the failure domains for kcp
	Azs []string `json:"azs"`

	// Root disk size
	DiskGiB int32 `json:"diskGiB,omitempty"`

	// flavor config
	// Required: true
	FlavorConfig *V1OpenstackFlavorConfig `json:"flavorConfig"`

	// subnet
	Subnet *V1OpenStackResource `json:"subnet,omitempty"`
}

// Validate validates this v1 open stack machine pool cloud config entity
func (m *V1OpenStackMachinePoolCloudConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlavorConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OpenStackMachinePoolCloudConfigEntity) validateFlavorConfig(formats strfmt.Registry) error {

	if err := validate.Required("flavorConfig", "body", m.FlavorConfig); err != nil {
		return err
	}

	if m.FlavorConfig != nil {
		if err := m.FlavorConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flavorConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1OpenStackMachinePoolCloudConfigEntity) validateSubnet(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {
		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OpenStackMachinePoolCloudConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OpenStackMachinePoolCloudConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1OpenStackMachinePoolCloudConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}