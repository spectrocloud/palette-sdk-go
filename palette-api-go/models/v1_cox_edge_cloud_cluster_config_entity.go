// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CoxEdgeCloudClusterConfigEntity CoxEdge cloud cluster config entity
//
// swagger:model v1CoxEdgeCloudClusterConfigEntity
type V1CoxEdgeCloudClusterConfigEntity struct {

	// cluster config
	ClusterConfig *V1CoxEdgeClusterConfig `json:"clusterConfig,omitempty"`
}

// Validate validates this v1 cox edge cloud cluster config entity
func (m *V1CoxEdgeCloudClusterConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CoxEdgeCloudClusterConfigEntity) validateClusterConfig(formats strfmt.Registry) error {

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
func (m *V1CoxEdgeCloudClusterConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CoxEdgeCloudClusterConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1CoxEdgeCloudClusterConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
