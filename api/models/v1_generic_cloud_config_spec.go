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

// V1GenericCloudConfigSpec Generic CloudConfig spec for all cloud types
//
// swagger:model v1GenericCloudConfigSpec
type V1GenericCloudConfigSpec struct {

	// Cloud account reference is optional and dynamically handled based on the kind
	CloudAccountRef *V1ObjectReference `json:"cloudAccountRef,omitempty"`

	// cluster config
	ClusterConfig *V1GenericClusterConfig `json:"clusterConfig,omitempty"`

	// Appliances (Edge Host) uids
	EdgeHostRefs []*V1ObjectReference `json:"edgeHostRefs"`

	// machine pool config
	MachinePoolConfig []*V1GenericMachinePoolConfig `json:"machinePoolConfig"`
}

// Validate validates this v1 generic cloud config spec
func (m *V1GenericCloudConfigSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudAccountRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeHostRefs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachinePoolConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GenericCloudConfigSpec) validateCloudAccountRef(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudAccountRef) { // not required
		return nil
	}

	if m.CloudAccountRef != nil {
		if err := m.CloudAccountRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudAccountRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1GenericCloudConfigSpec) validateClusterConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterConfig) { // not required
		return nil
	}

	if m.ClusterConfig != nil {
		if err := m.ClusterConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1GenericCloudConfigSpec) validateEdgeHostRefs(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeHostRefs) { // not required
		return nil
	}

	for i := 0; i < len(m.EdgeHostRefs); i++ {
		if swag.IsZero(m.EdgeHostRefs[i]) { // not required
			continue
		}

		if m.EdgeHostRefs[i] != nil {
			if err := m.EdgeHostRefs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeHostRefs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1GenericCloudConfigSpec) validateMachinePoolConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.MachinePoolConfig) { // not required
		return nil
	}

	for i := 0; i < len(m.MachinePoolConfig); i++ {
		if swag.IsZero(m.MachinePoolConfig[i]) { // not required
			continue
		}

		if m.MachinePoolConfig[i] != nil {
			if err := m.MachinePoolConfig[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machinePoolConfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GenericCloudConfigSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GenericCloudConfigSpec) UnmarshalBinary(b []byte) error {
	var res V1GenericCloudConfigSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}