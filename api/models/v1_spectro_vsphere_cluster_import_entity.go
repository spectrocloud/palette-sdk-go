// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroVsphereClusterImportEntity Spectro Vsphere cluster import request payload
//
// swagger:model v1SpectroVsphereClusterImportEntity
type V1SpectroVsphereClusterImportEntity struct {

	// metadata
	Metadata *V1ObjectMetaInputEntity `json:"metadata,omitempty"`

	// spec
	Spec *V1SpectroVsphereClusterImportEntitySpec `json:"spec,omitempty"`
}

// Validate validates this v1 spectro vsphere cluster import entity
func (m *V1SpectroVsphereClusterImportEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroVsphereClusterImportEntity) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroVsphereClusterImportEntity) validateSpec(formats strfmt.Registry) error {

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
func (m *V1SpectroVsphereClusterImportEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroVsphereClusterImportEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroVsphereClusterImportEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroVsphereClusterImportEntitySpec v1 spectro vsphere cluster import entity spec
//
// swagger:model V1SpectroVsphereClusterImportEntitySpec
type V1SpectroVsphereClusterImportEntitySpec struct {

	// cluster config
	ClusterConfig *V1ImportClusterConfig `json:"clusterConfig,omitempty"`
}

// Validate validates this v1 spectro vsphere cluster import entity spec
func (m *V1SpectroVsphereClusterImportEntitySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroVsphereClusterImportEntitySpec) validateClusterConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterConfig) { // not required
		return nil
	}

	if m.ClusterConfig != nil {
		if err := m.ClusterConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "clusterConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroVsphereClusterImportEntitySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroVsphereClusterImportEntitySpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroVsphereClusterImportEntitySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
