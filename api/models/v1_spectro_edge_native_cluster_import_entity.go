// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroEdgeNativeClusterImportEntity Spectro EdgeNative cluster import request payload
//
// swagger:model v1SpectroEdgeNativeClusterImportEntity
type V1SpectroEdgeNativeClusterImportEntity struct {

	// metadata
	Metadata *V1ObjectMetaInputEntity `json:"metadata,omitempty"`

	// spec
	Spec *V1SpectroEdgeNativeClusterImportEntitySpec `json:"spec,omitempty"`
}

// Validate validates this v1 spectro edge native cluster import entity
func (m *V1SpectroEdgeNativeClusterImportEntity) Validate(formats strfmt.Registry) error {
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

func (m *V1SpectroEdgeNativeClusterImportEntity) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroEdgeNativeClusterImportEntity) validateSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 spectro edge native cluster import entity based on the context it is used
func (m *V1SpectroEdgeNativeClusterImportEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroEdgeNativeClusterImportEntity) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroEdgeNativeClusterImportEntity) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.Spec != nil {

		if swag.IsZero(m.Spec) { // not required
			return nil
		}

		if err := m.Spec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroEdgeNativeClusterImportEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroEdgeNativeClusterImportEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroEdgeNativeClusterImportEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroEdgeNativeClusterImportEntitySpec v1 spectro edge native cluster import entity spec
//
// swagger:model V1SpectroEdgeNativeClusterImportEntitySpec
type V1SpectroEdgeNativeClusterImportEntitySpec struct {

	// cluster config
	ClusterConfig *V1ImportClusterConfig `json:"clusterConfig,omitempty"`
}

// Validate validates this v1 spectro edge native cluster import entity spec
func (m *V1SpectroEdgeNativeClusterImportEntitySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroEdgeNativeClusterImportEntitySpec) validateClusterConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterConfig) { // not required
		return nil
	}

	if m.ClusterConfig != nil {
		if err := m.ClusterConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "clusterConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec" + "." + "clusterConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 spectro edge native cluster import entity spec based on the context it is used
func (m *V1SpectroEdgeNativeClusterImportEntitySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroEdgeNativeClusterImportEntitySpec) contextValidateClusterConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterConfig != nil {

		if swag.IsZero(m.ClusterConfig) { // not required
			return nil
		}

		if err := m.ClusterConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "clusterConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec" + "." + "clusterConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroEdgeNativeClusterImportEntitySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroEdgeNativeClusterImportEntitySpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroEdgeNativeClusterImportEntitySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
