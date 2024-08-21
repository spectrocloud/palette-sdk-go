// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroGcpClusterImportEntity Spectro GCP cluster import request payload
//
// swagger:model v1SpectroGcpClusterImportEntity
type V1SpectroGcpClusterImportEntity struct {

	// metadata
	Metadata *V1ObjectMetaInputEntity `json:"metadata,omitempty"`

	// spec
	Spec *V1SpectroGcpClusterImportEntitySpec `json:"spec,omitempty"`
}

// Validate validates this v1 spectro gcp cluster import entity
func (m *V1SpectroGcpClusterImportEntity) Validate(formats strfmt.Registry) error {
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

func (m *V1SpectroGcpClusterImportEntity) validateMetadata(formats strfmt.Registry) error {

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

func (m *V1SpectroGcpClusterImportEntity) validateSpec(formats strfmt.Registry) error {

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
func (m *V1SpectroGcpClusterImportEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroGcpClusterImportEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroGcpClusterImportEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroGcpClusterImportEntitySpec v1 spectro gcp cluster import entity spec
//
// swagger:model V1SpectroGcpClusterImportEntitySpec
type V1SpectroGcpClusterImportEntitySpec struct {

	// cluster config
	ClusterConfig *V1ImportClusterConfig `json:"clusterConfig,omitempty"`
}

// Validate validates this v1 spectro gcp cluster import entity spec
func (m *V1SpectroGcpClusterImportEntitySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroGcpClusterImportEntitySpec) validateClusterConfig(formats strfmt.Registry) error {

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
func (m *V1SpectroGcpClusterImportEntitySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroGcpClusterImportEntitySpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroGcpClusterImportEntitySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}