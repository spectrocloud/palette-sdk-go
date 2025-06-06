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

// V1SpectroEksClusterEntity Spectro EKS cluster request payload for create and update
//
// swagger:model v1SpectroEksClusterEntity
type V1SpectroEksClusterEntity struct {

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// spec
	Spec *V1SpectroEksClusterEntitySpec `json:"spec,omitempty"`
}

// Validate validates this v1 spectro eks cluster entity
func (m *V1SpectroEksClusterEntity) Validate(formats strfmt.Registry) error {
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

func (m *V1SpectroEksClusterEntity) validateMetadata(formats strfmt.Registry) error {
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

func (m *V1SpectroEksClusterEntity) validateSpec(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 spectro eks cluster entity based on the context it is used
func (m *V1SpectroEksClusterEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *V1SpectroEksClusterEntity) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1SpectroEksClusterEntity) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1SpectroEksClusterEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroEksClusterEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroEksClusterEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroEksClusterEntitySpec v1 spectro eks cluster entity spec
//
// swagger:model V1SpectroEksClusterEntitySpec
type V1SpectroEksClusterEntitySpec struct {

	// Cloud account uid to be used for cluster provisioning
	// Required: true
	CloudAccountUID *string `json:"cloudAccountUid"`

	// cloud config
	// Required: true
	CloudConfig *V1EksClusterConfig `json:"cloudConfig"`

	// General cluster configuration like health, patching settings, namespace resource allocation, rbac
	ClusterConfig *V1ClusterConfigEntity `json:"clusterConfig,omitempty"`

	// fargate profiles
	FargateProfiles []*V1FargateProfile `json:"fargateProfiles"`

	// machinepoolconfig
	Machinepoolconfig []*V1EksMachinePoolConfigEntity `json:"machinepoolconfig"`

	// policies
	Policies *V1SpectroClusterPolicies `json:"policies,omitempty"`

	// profiles
	Profiles []*V1SpectroClusterProfileEntity `json:"profiles"`
}

// Validate validates this v1 spectro eks cluster entity spec
func (m *V1SpectroEksClusterEntitySpec) Validate(formats strfmt.Registry) error {
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

	if err := m.validateFargateProfiles(formats); err != nil {
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

func (m *V1SpectroEksClusterEntitySpec) validateCloudAccountUID(formats strfmt.Registry) error {

	if err := validate.Required("spec"+"."+"cloudAccountUid", "body", m.CloudAccountUID); err != nil {
		return err
	}

	return nil
}

func (m *V1SpectroEksClusterEntitySpec) validateCloudConfig(formats strfmt.Registry) error {

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

func (m *V1SpectroEksClusterEntitySpec) validateClusterConfig(formats strfmt.Registry) error {
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

func (m *V1SpectroEksClusterEntitySpec) validateFargateProfiles(formats strfmt.Registry) error {
	if swag.IsZero(m.FargateProfiles) { // not required
		return nil
	}

	for i := 0; i < len(m.FargateProfiles); i++ {
		if swag.IsZero(m.FargateProfiles[i]) { // not required
			continue
		}

		if m.FargateProfiles[i] != nil {
			if err := m.FargateProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "fargateProfiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "fargateProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroEksClusterEntitySpec) validateMachinepoolconfig(formats strfmt.Registry) error {
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

func (m *V1SpectroEksClusterEntitySpec) validatePolicies(formats strfmt.Registry) error {
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

func (m *V1SpectroEksClusterEntitySpec) validateProfiles(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 spectro eks cluster entity spec based on the context it is used
func (m *V1SpectroEksClusterEntitySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFargateProfiles(ctx, formats); err != nil {
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

func (m *V1SpectroEksClusterEntitySpec) contextValidateCloudConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1SpectroEksClusterEntitySpec) contextValidateClusterConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1SpectroEksClusterEntitySpec) contextValidateFargateProfiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FargateProfiles); i++ {

		if m.FargateProfiles[i] != nil {

			if swag.IsZero(m.FargateProfiles[i]) { // not required
				return nil
			}

			if err := m.FargateProfiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "fargateProfiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spec" + "." + "fargateProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroEksClusterEntitySpec) contextValidateMachinepoolconfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1SpectroEksClusterEntitySpec) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1SpectroEksClusterEntitySpec) contextValidateProfiles(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1SpectroEksClusterEntitySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroEksClusterEntitySpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroEksClusterEntitySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
