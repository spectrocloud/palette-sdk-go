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

// V1OverlordMaasCloudConfig v1 overlord maas cloud config
//
// swagger:model v1OverlordMaasCloudConfig
type V1OverlordMaasCloudConfig struct {

	// cluster config
	ClusterConfig *V1MaasClusterConfig `json:"clusterConfig,omitempty"`

	// Cluster profiles pack configuration for private gateway cluster
	ClusterProfiles []*V1SpectroClusterProfileEntity `json:"clusterProfiles"`

	// clusterSettings is the generic configuration related to a cluster like OS patch, Rbac, Namespace allocation
	ClusterSettings *V1ClusterConfigEntity `json:"clusterSettings,omitempty"`

	// machine config
	MachineConfig *V1MaasMachineConfigEntity `json:"machineConfig,omitempty"`

	// size of the pool, number of machines
	Size int32 `json:"size,omitempty"`
}

// Validate validates this v1 overlord maas cloud config
func (m *V1OverlordMaasCloudConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OverlordMaasCloudConfig) validateClusterConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterConfig) { // not required
		return nil
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

func (m *V1OverlordMaasCloudConfig) validateClusterProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterProfiles) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterProfiles); i++ {
		if swag.IsZero(m.ClusterProfiles[i]) { // not required
			continue
		}

		if m.ClusterProfiles[i] != nil {
			if err := m.ClusterProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1OverlordMaasCloudConfig) validateClusterSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterSettings) { // not required
		return nil
	}

	if m.ClusterSettings != nil {
		if err := m.ClusterSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterSettings")
			}
			return err
		}
	}

	return nil
}

func (m *V1OverlordMaasCloudConfig) validateMachineConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.MachineConfig) { // not required
		return nil
	}

	if m.MachineConfig != nil {
		if err := m.MachineConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OverlordMaasCloudConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OverlordMaasCloudConfig) UnmarshalBinary(b []byte) error {
	var res V1OverlordMaasCloudConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}