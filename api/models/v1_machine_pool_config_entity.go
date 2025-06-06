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

// V1MachinePoolConfigEntity Machine pool configuration for the cluster
//
// swagger:model v1MachinePoolConfigEntity
type V1MachinePoolConfigEntity struct {

	// Additional labels to be part of the machine pool
	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	// AdditionalTags is an optional set of tags to add to resources managed by the provider, in addition to the ones added by default. For eg., tags for EKS nodeGroup or EKS NodegroupIAMRole
	AdditionalTags map[string]string `json:"additionalTags,omitempty"`

	// Whether this pool is for control plane
	IsControlPlane bool `json:"isControlPlane"`

	// Labels for this machine pool, example: control-plane/worker, gpu, windows
	// Required: true
	Labels []string `json:"labels"`

	// machine pool properties
	MachinePoolProperties *V1MachinePoolProperties `json:"machinePoolProperties,omitempty"`

	// Max size of the pool, for scaling
	MaxSize int32 `json:"maxSize,omitempty"`

	// Min size of the pool, for scaling
	MinSize int32 `json:"minSize,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Minimum number of seconds a node should be Ready, before the next node is selected for repave. Applicable only for workerpools in infrastructure cluster
	NodeRepaveInterval int32 `json:"nodeRepaveInterval,omitempty"`

	// Size of the pool, number of nodes/machines
	// Required: true
	Size *int32 `json:"size"`

	// control plane or worker taints
	// Unique: true
	Taints []*V1Taint `json:"taints"`

	// Rolling update strategy for this machine pool if not specified, will use ScaleOut
	UpdateStrategy *V1UpdateStrategy `json:"updateStrategy,omitempty"`

	// If IsControlPlane==true && useControlPlaneAsWorker==true, then will remove control plane taint this will not be used for worker pools
	UseControlPlaneAsWorker bool `json:"useControlPlaneAsWorker"`
}

// Validate validates this v1 machine pool config entity
func (m *V1MachinePoolConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachinePoolProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
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

func (m *V1MachinePoolConfigEntity) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("labels", "body", m.Labels); err != nil {
		return err
	}

	return nil
}

func (m *V1MachinePoolConfigEntity) validateMachinePoolProperties(formats strfmt.Registry) error {
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

func (m *V1MachinePoolConfigEntity) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1MachinePoolConfigEntity) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *V1MachinePoolConfigEntity) validateTaints(formats strfmt.Registry) error {
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

func (m *V1MachinePoolConfigEntity) validateUpdateStrategy(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 machine pool config entity based on the context it is used
func (m *V1MachinePoolConfigEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMachinePoolProperties(ctx, formats); err != nil {
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

func (m *V1MachinePoolConfigEntity) contextValidateMachinePoolProperties(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1MachinePoolConfigEntity) contextValidateTaints(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1MachinePoolConfigEntity) contextValidateUpdateStrategy(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1MachinePoolConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachinePoolConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1MachinePoolConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
