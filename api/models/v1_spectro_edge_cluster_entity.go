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

// V1SpectroEdgeClusterEntity Edge cluster request payload for create and update
//
// swagger:model v1SpectroEdgeClusterEntity
type V1SpectroEdgeClusterEntity struct {

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// spec
	Spec *V1SpectroEdgeClusterEntitySpec `json:"spec,omitempty"`
}

// Validate validates this v1 spectro edge cluster entity
func (m *V1SpectroEdgeClusterEntity) Validate(formats strfmt.Registry) error {
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

func (m *V1SpectroEdgeClusterEntity) validateMetadata(formats strfmt.Registry) error {

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

func (m *V1SpectroEdgeClusterEntity) validateSpec(formats strfmt.Registry) error {

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
func (m *V1SpectroEdgeClusterEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroEdgeClusterEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroEdgeClusterEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroEdgeClusterEntitySpec v1 spectro edge cluster entity spec
//
// swagger:model V1SpectroEdgeClusterEntitySpec
type V1SpectroEdgeClusterEntitySpec struct {

	// cloud config
	CloudConfig *V1EdgeClusterConfig `json:"cloudConfig,omitempty"`

	// General cluster configuration like health, patching settings, namespace resource allocation, rbac
	ClusterConfig *V1ClusterConfigEntity `json:"clusterConfig,omitempty"`

	// machinepoolconfig
	Machinepoolconfig []*V1EdgeMachinePoolConfigEntity `json:"machinepoolconfig"`

	// policies
	Policies *V1SpectroClusterPolicies `json:"policies,omitempty"`

	// profiles
	Profiles []*V1SpectroClusterProfileEntity `json:"profiles"`

	// variables
	Variables []*V1SpectroClusterVariable `json:"variables"`
}

// Validate validates this v1 spectro edge cluster entity spec
func (m *V1SpectroEdgeClusterEntitySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudConfig(formats); err != nil {
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

func (m *V1SpectroEdgeClusterEntitySpec) validateCloudConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudConfig) { // not required
		return nil
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

func (m *V1SpectroEdgeClusterEntitySpec) validateClusterConfig(formats strfmt.Registry) error {

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

func (m *V1SpectroEdgeClusterEntitySpec) validateMachinepoolconfig(formats strfmt.Registry) error {

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

func (m *V1SpectroEdgeClusterEntitySpec) validatePolicies(formats strfmt.Registry) error {

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

func (m *V1SpectroEdgeClusterEntitySpec) validateProfiles(formats strfmt.Registry) error {

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

func (m *V1SpectroEdgeClusterEntitySpec) validateVariables(formats strfmt.Registry) error {

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
func (m *V1SpectroEdgeClusterEntitySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroEdgeClusterEntitySpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroEdgeClusterEntitySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}