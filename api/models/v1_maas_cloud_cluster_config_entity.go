// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MaasCloudClusterConfigEntity Maas cloud cluster config entity
//
// swagger:model v1MaasCloudClusterConfigEntity
type V1MaasCloudClusterConfigEntity struct {

	// cluster config
	ClusterConfig *V1MaasClusterConfig `json:"clusterConfig,omitempty"`
}

// Validate validates this v1 maas cloud cluster config entity
func (m *V1MaasCloudClusterConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MaasCloudClusterConfigEntity) validateClusterConfig(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V1MaasCloudClusterConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MaasCloudClusterConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1MaasCloudClusterConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
