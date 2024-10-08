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

// V1LibvirtPlacementEntity Libvirt placement config
//
// swagger:model v1LibvirtPlacementEntity
type V1LibvirtPlacementEntity struct {

	// data storage pool
	DataStoragePool string `json:"dataStoragePool,omitempty"`

	// GPUDevices defines an array of gpu device for a specific edge host. This will be overridden by edge host GPU devices if configured during registration.
	GpuDevices []*V1GPUDeviceSpec `json:"gpuDevices"`

	// host Uid
	// Required: true
	HostUID *string `json:"hostUid"`

	// networks
	Networks []*V1LibvirtNetworkSpec `json:"networks"`

	// source storage pool
	SourceStoragePool string `json:"sourceStoragePool,omitempty"`

	// target storage pool
	TargetStoragePool string `json:"targetStoragePool,omitempty"`
}

// Validate validates this v1 libvirt placement entity
func (m *V1LibvirtPlacementEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGpuDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1LibvirtPlacementEntity) validateGpuDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.GpuDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.GpuDevices); i++ {
		if swag.IsZero(m.GpuDevices[i]) { // not required
			continue
		}

		if m.GpuDevices[i] != nil {
			if err := m.GpuDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpuDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1LibvirtPlacementEntity) validateHostUID(formats strfmt.Registry) error {

	if err := validate.Required("hostUid", "body", m.HostUID); err != nil {
		return err
	}

	return nil
}

func (m *V1LibvirtPlacementEntity) validateNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1LibvirtPlacementEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LibvirtPlacementEntity) UnmarshalBinary(b []byte) error {
	var res V1LibvirtPlacementEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
