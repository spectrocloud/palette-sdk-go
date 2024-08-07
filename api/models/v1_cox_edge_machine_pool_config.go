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

// V1CoxEdgeMachinePoolConfig v1 cox edge machine pool config
//
// swagger:model v1CoxEdgeMachinePoolConfig
type V1CoxEdgeMachinePoolConfig struct {

	// additionalLabels
	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	// AdditionalTags is an optional set of tags to add to resources managed by the provider, in addition to the ones added by default. For eg., tags for EKS nodeGroup or EKS NodegroupIAMRole
	AdditionalTags map[string]string `json:"additionalTags,omitempty"`

	// deployments
	Deployments []*V1CoxEdgeDeployment `json:"deployments"`

	// instance config
	InstanceConfig *V1InstanceConfig `json:"instanceConfig,omitempty"`

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

	// Array of coxedge load persistent storages
	// Unique: true
	PersistentStorages []*V1CoxEdgeLoadPersistentStorage `json:"persistentStorages"`

	// security group rules
	SecurityGroupRules []*V1CoxEdgeSecurityGroupRule `json:"securityGroupRules"`

	// size of the pool, number of machines
	Size int32 `json:"size,omitempty"`

	// spec
	Spec string `json:"spec,omitempty"`

	// control plane or worker taints
	// Unique: true
	Taints []*V1Taint `json:"taints"`

	// rolling update strategy for this machinepool if not specified, will use ScaleOut
	UpdateStrategy *V1UpdateStrategy `json:"updateStrategy,omitempty"`

	// if IsControlPlane==true && useControlPlaneAsWorker==true, then will remove control plane taint this will not be used for worker pools
	UseControlPlaneAsWorker bool `json:"useControlPlaneAsWorker,omitempty"`
}

// Validate validates this v1 cox edge machine pool config
func (m *V1CoxEdgeMachinePoolConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsControlPlane(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachinePoolProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersistentStorages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupRules(formats); err != nil {
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

func (m *V1CoxEdgeMachinePoolConfig) validateDeployments(formats strfmt.Registry) error {

	if swag.IsZero(m.Deployments) { // not required
		return nil
	}

	for i := 0; i < len(m.Deployments); i++ {
		if swag.IsZero(m.Deployments[i]) { // not required
			continue
		}

		if m.Deployments[i] != nil {
			if err := m.Deployments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CoxEdgeMachinePoolConfig) validateInstanceConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceConfig) { // not required
		return nil
	}

	if m.InstanceConfig != nil {
		if err := m.InstanceConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instanceConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1CoxEdgeMachinePoolConfig) validateIsControlPlane(formats strfmt.Registry) error {

	if err := validate.Required("isControlPlane", "body", m.IsControlPlane); err != nil {
		return err
	}

	return nil
}

func (m *V1CoxEdgeMachinePoolConfig) validateMachinePoolProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.MachinePoolProperties) { // not required
		return nil
	}

	if m.MachinePoolProperties != nil {
		if err := m.MachinePoolProperties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machinePoolProperties")
			}
			return err
		}
	}

	return nil
}

func (m *V1CoxEdgeMachinePoolConfig) validatePersistentStorages(formats strfmt.Registry) error {

	if swag.IsZero(m.PersistentStorages) { // not required
		return nil
	}

	if err := validate.UniqueItems("persistentStorages", "body", m.PersistentStorages); err != nil {
		return err
	}

	for i := 0; i < len(m.PersistentStorages); i++ {
		if swag.IsZero(m.PersistentStorages[i]) { // not required
			continue
		}

		if m.PersistentStorages[i] != nil {
			if err := m.PersistentStorages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("persistentStorages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CoxEdgeMachinePoolConfig) validateSecurityGroupRules(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroupRules) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityGroupRules); i++ {
		if swag.IsZero(m.SecurityGroupRules[i]) { // not required
			continue
		}

		if m.SecurityGroupRules[i] != nil {
			if err := m.SecurityGroupRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityGroupRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CoxEdgeMachinePoolConfig) validateTaints(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CoxEdgeMachinePoolConfig) validateUpdateStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateStrategy) { // not required
		return nil
	}

	if m.UpdateStrategy != nil {
		if err := m.UpdateStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateStrategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CoxEdgeMachinePoolConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CoxEdgeMachinePoolConfig) UnmarshalBinary(b []byte) error {
	var res V1CoxEdgeMachinePoolConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
