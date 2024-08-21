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

// V1ClusterConfig v1 cluster config
//
// swagger:model v1ClusterConfig
type V1ClusterConfig struct {

	// ClusterMetaAttribute contains additional cluster metadata information.
	ClusterMetaAttribute string `json:"clusterMetaAttribute,omitempty"`

	// Deprecated. Use clusterResources
	ClusterRbac []*V1ResourceReference `json:"clusterRbac"`

	// ClusterResources defines the managment of namespace resource allocations, role bindings.
	ClusterResources *V1ClusterResources `json:"clusterResources,omitempty"`

	// ControlPlaneHealthCheckTimeout is the timeout to check for ready state of the control plane nodes. If the node is not ready within the time out set, the node will be deleted and a new node will be launched.
	ControlPlaneHealthCheckTimeout string `json:"controlPlaneHealthCheckTimeout,omitempty"`

	// HostClusterConfiguration defines the configuration of host clusters, where virtual clusters be deployed
	HostClusterConfig *V1HostClusterConfig `json:"hostClusterConfig,omitempty"`

	// lifecycle config
	LifecycleConfig *V1LifecycleConfig `json:"lifecycleConfig,omitempty"`

	// MachineHealthCheckConfig defines the healthcheck timeouts for the node. The timeouts are configured by the user to overide the default healthchecks.
	MachineHealthConfig *V1MachineHealthCheckConfig `json:"machineHealthConfig,omitempty"`

	// MachineManagementConfig defines the management configurations for the node. Patching OS security updates etc can be configured by user.
	MachineManagementConfig *V1MachineManagementConfig `json:"machineManagementConfig,omitempty"`

	// UpdateWorkerPoolsInParallel is used to decide if the update of workerpools happen in parallel. When this flag is false, the workerpools are updated sequentially.
	UpdateWorkerPoolsInParallel bool `json:"updateWorkerPoolsInParallel,omitempty"`
}

// Validate validates this v1 cluster config
func (m *V1ClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterRbac(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifecycleConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineHealthConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineManagementConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterConfig) validateClusterRbac(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterRbac) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterRbac); i++ {
		if swag.IsZero(m.ClusterRbac[i]) { // not required
			continue
		}

		if m.ClusterRbac[i] != nil {
			if err := m.ClusterRbac[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterRbac" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterConfig) validateClusterResources(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterResources) { // not required
		return nil
	}

	if m.ClusterResources != nil {
		if err := m.ClusterResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterResources")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterConfig) validateHostClusterConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.HostClusterConfig) { // not required
		return nil
	}

	if m.HostClusterConfig != nil {
		if err := m.HostClusterConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostClusterConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterConfig) validateLifecycleConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.LifecycleConfig) { // not required
		return nil
	}

	if m.LifecycleConfig != nil {
		if err := m.LifecycleConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lifecycleConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterConfig) validateMachineHealthConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.MachineHealthConfig) { // not required
		return nil
	}

	if m.MachineHealthConfig != nil {
		if err := m.MachineHealthConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineHealthConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterConfig) validateMachineManagementConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.MachineManagementConfig) { // not required
		return nil
	}

	if m.MachineManagementConfig != nil {
		if err := m.MachineManagementConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineManagementConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterConfig) UnmarshalBinary(b []byte) error {
	var res V1ClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}