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

// V1SpectroCoxEdgeClusterEntity CoxEdge cluster request payload for create and update
//
// swagger:model v1SpectroCoxEdgeClusterEntity
type V1SpectroCoxEdgeClusterEntity struct {

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// spec
	Spec *V1SpectroCoxEdgeClusterEntitySpec `json:"spec,omitempty"`
}

// Validate validates this v1 spectro cox edge cluster entity
func (m *V1SpectroCoxEdgeClusterEntity) Validate(formats strfmt.Registry) error {
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

func (m *V1SpectroCoxEdgeClusterEntity) validateMetadata(formats strfmt.Registry) error {

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

func (m *V1SpectroCoxEdgeClusterEntity) validateSpec(formats strfmt.Registry) error {

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
func (m *V1SpectroCoxEdgeClusterEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroCoxEdgeClusterEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroCoxEdgeClusterEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroCoxEdgeClusterEntitySpec v1 spectro cox edge cluster entity spec
//
// swagger:model V1SpectroCoxEdgeClusterEntitySpec
type V1SpectroCoxEdgeClusterEntitySpec struct {

	// Cloud account uid to be used for cluster provisioning
	// Required: true
	CloudAccountUID *string `json:"cloudAccountUid"`

	// cloud config
	// Required: true
	CloudConfig *V1CoxEdgeClusterConfig `json:"cloudConfig"`

	// cloud type
	// Required: true
	CloudType *string `json:"cloudType"`

	// General cluster configuration like health, patching settings, namespace resource allocation, rbac
	ClusterConfig *V1ClusterConfigEntity `json:"clusterConfig,omitempty"`

	// machinepoolconfig
	Machinepoolconfig []*V1CoxEdgeMachinePoolConfigEntity `json:"machinepoolconfig"`

	// policies
	Policies *V1SpectroClusterPolicies `json:"policies,omitempty"`

	// profiles
	Profiles []*V1SpectroClusterProfileEntity `json:"profiles"`

	// variables
	Variables []*V1SpectroClusterVariable `json:"variables"`
}

// Validate validates this v1 spectro cox edge cluster entity spec
func (m *V1SpectroCoxEdgeClusterEntitySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudAccountUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachinepoolconfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroCoxEdgeClusterEntitySpec) validateCloudAccountUID(formats strfmt.Registry) error {

	if err := validate.Required("spec"+"."+"cloudAccountUid", "body", m.CloudAccountUID); err != nil {
		return err
	}

	return nil
}

func (m *V1SpectroCoxEdgeClusterEntitySpec) validateCloudConfig(formats strfmt.Registry) error {

	if err := validate.Required("spec"+"."+"cloudConfig", "body", m.CloudConfig); err != nil {
		return err
	}

	if m.CloudConfig != nil {
		if err := m.CloudConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "cloudConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroCoxEdgeClusterEntitySpec) validateCloudType(formats strfmt.Registry) error {

	if err := validate.Required("spec"+"."+"cloudType", "body", m.CloudType); err != nil {
		return err
	}

	return nil
}

func (m *V1SpectroCoxEdgeClusterEntitySpec) validateClusterConfig(formats strfmt.Registry) error {

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

func (m *V1SpectroCoxEdgeClusterEntitySpec) validateMachinepoolconfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Machinepoolconfig) { // not required
		return nil
	}

	for i := 0; i < len(m.Machinepoolconfig); i++ {
		if swag.IsZero(m.Machinepoolconfig[i]) { // not required
			continue
		}

		if m.Machinepoolconfig[i] != nil {
			if err := m.Machinepoolconfig[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "machinepoolconfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroCoxEdgeClusterEntitySpec) validatePolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	if m.Policies != nil {
		if err := m.Policies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "policies")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroCoxEdgeClusterEntitySpec) validateProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Profiles) { // not required
		return nil
	}

	for i := 0; i < len(m.Profiles); i++ {
		if swag.IsZero(m.Profiles[i]) { // not required
			continue
		}

		if m.Profiles[i] != nil {
			if err := m.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroCoxEdgeClusterEntitySpec) validateVariables(formats strfmt.Registry) error {

	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for i := 0; i < len(m.Variables); i++ {
		if swag.IsZero(m.Variables[i]) { // not required
			continue
		}

		if m.Variables[i] != nil {
			if err := m.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroCoxEdgeClusterEntitySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroCoxEdgeClusterEntitySpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroCoxEdgeClusterEntitySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
