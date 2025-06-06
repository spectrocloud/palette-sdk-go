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
	"github.com/go-openapi/validate"
)

// V1SpectroMaasClusterEntity Spectro Maas cluster request payload for create and update
//
// swagger:model v1SpectroMaasClusterEntity
type V1SpectroMaasClusterEntity struct {

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// spec
	Spec *V1SpectroMaasClusterEntitySpec `json:"spec,omitempty"`
}

// Validate validates this v1 spectro maas cluster entity
func (m *V1SpectroMaasClusterEntity) Validate(formats strfmt.Registry) error {
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

func (m *V1SpectroMaasClusterEntity) validateMetadata(formats strfmt.Registry) error {
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

func (m *V1SpectroMaasClusterEntity) validateSpec(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 spectro maas cluster entity based on the context it is used
func (m *V1SpectroMaasClusterEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *V1SpectroMaasClusterEntity) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1SpectroMaasClusterEntity) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1SpectroMaasClusterEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroMaasClusterEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroMaasClusterEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroMaasClusterEntitySpec v1 spectro maas cluster entity spec
//
// swagger:model V1SpectroMaasClusterEntitySpec
type V1SpectroMaasClusterEntitySpec struct {

	// Cloud account uid to be used for cluster provisioning
	// Required: true
	CloudAccountUID *string `json:"cloudAccountUid"`

	// cloud config
	// Required: true
	CloudConfig *V1MaasClusterConfig `json:"cloudConfig"`

	// General cluster configuration like health, patching settings, namespace resource allocation, rbac
	ClusterConfig *V1ClusterConfigEntity `json:"clusterConfig,omitempty"`

	// machinepoolconfig
	Machinepoolconfig []*V1MaasMachinePoolConfigEntity `json:"machinepoolconfig"`

	// policies
	Policies *V1SpectroClusterPolicies `json:"policies,omitempty"`

	// profiles
	Profiles []*V1SpectroClusterProfileEntity `json:"profiles"`
}

// Validate validates this v1 spectro maas cluster entity spec
func (m *V1SpectroMaasClusterEntitySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudAccountUID(formats); err != nil {
		res = append(res, err)
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroMaasClusterEntitySpec) validateCloudAccountUID(formats strfmt.Registry) error {

	if err := validate.Required("spec"+"."+"cloudAccountUid", "body", m.CloudAccountUID); err != nil {
		return err
	}

	return nil
}

func (m *V1SpectroMaasClusterEntitySpec) validateCloudConfig(formats strfmt.Registry) error {

	if err := validate.Required("spec"+"."+"cloudConfig", "body", m.CloudConfig); err != nil {
		return err
	}

	if m.CloudConfig != nil {
		if err := m.CloudConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "cloudConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec" + "." + "cloudConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroMaasClusterEntitySpec) validateClusterConfig(formats strfmt.Registry) error {
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

func (m *V1SpectroMaasClusterEntitySpec) validateMachinepoolconfig(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "machinepoolconfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroMaasClusterEntitySpec) validatePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	if m.Policies != nil {
		if err := m.Policies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "policies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec" + "." + "policies")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroMaasClusterEntitySpec) validateProfiles(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 spectro maas cluster entity spec based on the context it is used
func (m *V1SpectroMaasClusterEntitySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachinepoolconfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProfiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroMaasClusterEntitySpec) contextValidateCloudConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudConfig != nil {

		if err := m.CloudConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "cloudConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec" + "." + "cloudConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroMaasClusterEntitySpec) contextValidateClusterConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1SpectroMaasClusterEntitySpec) contextValidateMachinepoolconfig(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Machinepoolconfig); i++ {

		if m.Machinepoolconfig[i] != nil {

			if swag.IsZero(m.Machinepoolconfig[i]) { // not required
				return nil
			}

			if err := m.Machinepoolconfig[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "machinepoolconfig" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "machinepoolconfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroMaasClusterEntitySpec) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

	if m.Policies != nil {

		if swag.IsZero(m.Policies) { // not required
			return nil
		}

		if err := m.Policies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "policies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spec" + "." + "policies")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroMaasClusterEntitySpec) contextValidateProfiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Profiles); i++ {

		if m.Profiles[i] != nil {

			if swag.IsZero(m.Profiles[i]) { // not required
				return nil
			}

			if err := m.Profiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "profiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroMaasClusterEntitySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroMaasClusterEntitySpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroMaasClusterEntitySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
