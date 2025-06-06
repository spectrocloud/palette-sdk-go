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

// V1VsphereMachinePoolConfig v1 vsphere machine pool config
//
// swagger:model v1VsphereMachinePoolConfig
type V1VsphereMachinePoolConfig struct {

	// additionalLabels
	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	// AdditionalTags is an optional set of tags to add to resources managed by the provider, in addition to the ones added by default. For eg., tags for EKS nodeGroup or EKS NodegroupIAMRole
	AdditionalTags map[string]string `json:"additionalTags,omitempty"`

	// InstanceType defines the required CPU, Memory, Storage
	// Required: true
	InstanceType *V1VsphereInstanceType `json:"instanceType"`

	// whether this pool is for control plane
	// Required: true
	IsControlPlane *bool `json:"isControlPlane"`

	// labels for this pool, example: control-plane/worker, gpu, windows
	Labels []string `json:"labels"`

	// machine pool properties
	MachinePoolProperties *V1MachinePoolProperties `json:"machinePoolProperties,omitempty"`

	// max size of the pool, for scaling
	MaxSize int32 `json:"maxSize,omitempty"`

	// min size of the pool, for scaling
	MinSize int32 `json:"minSize,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Minimum number of seconds a node should be Ready, before the next node is selected for repave. Applicable only for workerpools in infrastructure cluster
	NodeRepaveInterval int32 `json:"nodeRepaveInterval,omitempty"`

	// Placements configuration Placements If defined, will replace default values defined in VsphereClusterConfig Array means one MachinePool can span across multiple vsphere compute cluster
	Placements []*V1VspherePlacementConfig `json:"placements"`

	// size of the pool, number of machines
	Size int32 `json:"size,omitempty"`

	// control plane or worker taints
	// Unique: true
	Taints []*V1Taint `json:"taints"`

	// rolling update strategy for this machinepool if not specified, will use ScaleOut
	UpdateStrategy *V1UpdateStrategy `json:"updateStrategy,omitempty"`

	// if IsControlPlane==true && useControlPlaneAsWorker==true, then will remove control plane taint this will not be used for worker pools
	UseControlPlaneAsWorker bool `json:"useControlPlaneAsWorker"`
}

// Validate validates this v1 vsphere machine pool config
func (m *V1VsphereMachinePoolConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsControlPlane(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachinePoolProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlacements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VsphereMachinePoolConfig) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instanceType", "body", m.InstanceType); err != nil {
		return err
	}

	if m.InstanceType != nil {
		if err := m.InstanceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instanceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instanceType")
			}
			return err
		}
	}

	return nil
}

func (m *V1VsphereMachinePoolConfig) validateIsControlPlane(formats strfmt.Registry) error {

	if err := validate.Required("isControlPlane", "body", m.IsControlPlane); err != nil {
		return err
	}

	return nil
}

func (m *V1VsphereMachinePoolConfig) validateMachinePoolProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.MachinePoolProperties) { // not required
		return nil
	}

	if m.MachinePoolProperties != nil {
		if err := m.MachinePoolProperties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machinePoolProperties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machinePoolProperties")
			}
			return err
		}
	}

	return nil
}

func (m *V1VsphereMachinePoolConfig) validatePlacements(formats strfmt.Registry) error {
	if swag.IsZero(m.Placements) { // not required
		return nil
	}

	for i := 0; i < len(m.Placements); i++ {
		if swag.IsZero(m.Placements[i]) { // not required
			continue
		}

		if m.Placements[i] != nil {
			if err := m.Placements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("placements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("placements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VsphereMachinePoolConfig) validateTaints(formats strfmt.Registry) error {
	if swag.IsZero(m.Taints) { // not required
		return nil
	}

	if err := validate.UniqueItems("taints", "body", m.Taints); err != nil {
		return err
	}

	for i := 0; i < len(m.Taints); i++ {
		if swag.IsZero(m.Taints[i]) { // not required
			continue
		}

		if m.Taints[i] != nil {
			if err := m.Taints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VsphereMachinePoolConfig) validateUpdateStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateStrategy) { // not required
		return nil
	}

	if m.UpdateStrategy != nil {
		if err := m.UpdateStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateStrategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateStrategy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 vsphere machine pool config based on the context it is used
func (m *V1VsphereMachinePoolConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstanceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachinePoolProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlacements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VsphereMachinePoolConfig) contextValidateInstanceType(ctx context.Context, formats strfmt.Registry) error {

	if m.InstanceType != nil {

		if err := m.InstanceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instanceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instanceType")
			}
			return err
		}
	}

	return nil
}

func (m *V1VsphereMachinePoolConfig) contextValidateMachinePoolProperties(ctx context.Context, formats strfmt.Registry) error {

	if m.MachinePoolProperties != nil {

		if swag.IsZero(m.MachinePoolProperties) { // not required
			return nil
		}

		if err := m.MachinePoolProperties.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machinePoolProperties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machinePoolProperties")
			}
			return err
		}
	}

	return nil
}

func (m *V1VsphereMachinePoolConfig) contextValidatePlacements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Placements); i++ {

		if m.Placements[i] != nil {

			if swag.IsZero(m.Placements[i]) { // not required
				return nil
			}

			if err := m.Placements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("placements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("placements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VsphereMachinePoolConfig) contextValidateTaints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Taints); i++ {

		if m.Taints[i] != nil {

			if swag.IsZero(m.Taints[i]) { // not required
				return nil
			}

			if err := m.Taints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VsphereMachinePoolConfig) contextValidateUpdateStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdateStrategy != nil {

		if swag.IsZero(m.UpdateStrategy) { // not required
			return nil
		}

		if err := m.UpdateStrategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateStrategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateStrategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VsphereMachinePoolConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VsphereMachinePoolConfig) UnmarshalBinary(b []byte) error {
	var res V1VsphereMachinePoolConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
