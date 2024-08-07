// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1LibvirtCloudConfigSpec LibvirtCloudConfigSpec defines the desired state of LibvirtCloudConfig
//
// swagger:model v1LibvirtCloudConfigSpec
type V1LibvirtCloudConfigSpec struct {

	// cluster config
	// Required: true
	ClusterConfig *V1LibvirtClusterConfig `json:"clusterConfig"`

	// machine pool config
	// Required: true
	MachinePoolConfig []*V1LibvirtMachinePoolConfig `json:"machinePoolConfig"`
}

// Validate validates this v1 libvirt cloud config spec
func (m *V1LibvirtCloudConfigSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachinePoolConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1LibvirtCloudConfigSpec) validateClusterConfig(formats strfmt.Registry) error {

	if err := validate.Required("clusterConfig", "body", m.ClusterConfig); err != nil {
		return err
	}

	if m.ClusterConfig != nil {
		if err := m.ClusterConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1LibvirtCloudConfigSpec) validateMachinePoolConfig(formats strfmt.Registry) error {

	if err := validate.Required("machinePoolConfig", "body", m.MachinePoolConfig); err != nil {
		return err
	}

	for i := 0; i < len(m.MachinePoolConfig); i++ {
		if swag.IsZero(m.MachinePoolConfig[i]) { // not required
			continue
		}

		if m.MachinePoolConfig[i] != nil {
			if err := m.MachinePoolConfig[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machinePoolConfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1LibvirtCloudConfigSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LibvirtCloudConfigSpec) UnmarshalBinary(b []byte) error {
	var res V1LibvirtCloudConfigSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
