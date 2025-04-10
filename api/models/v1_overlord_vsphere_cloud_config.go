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

// V1OverlordVsphereCloudConfig v1 overlord vsphere cloud config
//
// swagger:model v1OverlordVsphereCloudConfig
type V1OverlordVsphereCloudConfig struct {

	// cluster config
	ClusterConfig *V1VsphereOverlordClusterConfigEntity `json:"clusterConfig,omitempty"`

	// Cluster profiles pack configuration for private gateway cluster
	ClusterProfiles []*V1SpectroClusterProfileEntity `json:"clusterProfiles"`

	// clusterSettings is the generic configuration related to a cluster like OS patch, Rbac, Namespace allocation
	ClusterSettings *V1ClusterConfigEntity `json:"clusterSettings,omitempty"`

	// size of the pool, number of machines
	Size int32 `json:"size,omitempty"`
}

// Validate validates this v1 overlord vsphere cloud config
func (m *V1OverlordVsphereCloudConfig) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OverlordVsphereCloudConfig) validateClusterConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterConfig) { // not required
		return nil
	}

	if m.ClusterConfig != nil {
		if err := m.ClusterConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1OverlordVsphereCloudConfig) validateClusterProfiles(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1OverlordVsphereCloudConfig) validateClusterSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterSettings) { // not required
		return nil
	}

	if m.ClusterSettings != nil {
		if err := m.ClusterSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterSettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 overlord vsphere cloud config based on the context it is used
func (m *V1OverlordVsphereCloudConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterProfiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OverlordVsphereCloudConfig) contextValidateClusterConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterConfig != nil {

		if swag.IsZero(m.ClusterConfig) { // not required
			return nil
		}

		if err := m.ClusterConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1OverlordVsphereCloudConfig) contextValidateClusterProfiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterProfiles); i++ {

		if m.ClusterProfiles[i] != nil {

			if swag.IsZero(m.ClusterProfiles[i]) { // not required
				return nil
			}

			if err := m.ClusterProfiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterProfiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1OverlordVsphereCloudConfig) contextValidateClusterSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterSettings != nil {

		if swag.IsZero(m.ClusterSettings) { // not required
			return nil
		}

		if err := m.ClusterSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OverlordVsphereCloudConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OverlordVsphereCloudConfig) UnmarshalBinary(b []byte) error {
	var res V1OverlordVsphereCloudConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
