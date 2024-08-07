// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroClusterAssetEntity Cluster asset
//
// swagger:model v1SpectroClusterAssetEntity
type V1SpectroClusterAssetEntity struct {

	// spec
	Spec *V1SpectroClusterAssetEntitySpec `json:"spec,omitempty"`
}

// Validate validates this v1 spectro cluster asset entity
func (m *V1SpectroClusterAssetEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterAssetEntity) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterAssetEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterAssetEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterAssetEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroClusterAssetEntitySpec v1 spectro cluster asset entity spec
//
// swagger:model V1SpectroClusterAssetEntitySpec
type V1SpectroClusterAssetEntitySpec struct {

	// frp kubeconfig
	FrpKubeconfig string `json:"frpKubeconfig,omitempty"`

	// kubeconfig
	Kubeconfig string `json:"kubeconfig,omitempty"`

	// kubeconfigclient
	Kubeconfigclient string `json:"kubeconfigclient,omitempty"`

	// manifest
	Manifest string `json:"manifest,omitempty"`
}

// Validate validates this v1 spectro cluster asset entity spec
func (m *V1SpectroClusterAssetEntitySpec) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterAssetEntitySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterAssetEntitySpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterAssetEntitySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
