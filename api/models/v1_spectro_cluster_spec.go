// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SpectroClusterSpec SpectroClusterSpec defines the desired state of SpectroCluster
//
// swagger:model v1SpectroClusterSpec
type V1SpectroClusterSpec struct {

	// CloudConfigRef point to the cloud configuration for the cluster, input by user Ref types are: AwsCloudConfig/VsphereCloudConfig/BaremetalConfig/ etc this user config will be used to generate cloud specific cluster/machine spec for cluster-api For VM, it will contain information needed to launch VMs, like cloud account, instance type For BM, it will contain actual baremetal machines
	CloudConfigRef *V1ObjectReference `json:"cloudConfigRef,omitempty"`

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// ClusterConfig is the configuration related to a general cluster. Configuration related to the health of the cluster.
	ClusterConfig *V1ClusterConfig `json:"clusterConfig,omitempty"`

	// When a cluster created from a clusterprofile at t1, ClusterProfileTemplate is a copy of the draft version or latest published version of the clusterprofileSpec.clusterprofileTemplate then clusterprofile may evolve to v2 at t2, but before user decide to upgrade the cluster, it will stay as it is when user decide to upgrade, clusterProfileTemplate will be updated from the clusterprofile pointed by ClusterProfileRef
	ClusterProfileTemplates []*V1ClusterProfileTemplate `json:"clusterProfileTemplates"`

	// cluster type
	// Enum: [PureManage AlloyMonitor AlloyAssist AlloyExtend]
	ClusterType string `json:"clusterType,omitempty"`
}

// Validate validates this v1 spectro cluster spec
func (m *V1SpectroClusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudConfigRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterProfileTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterSpec) validateCloudConfigRef(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudConfigRef) { // not required
		return nil
	}

	if m.CloudConfigRef != nil {
		if err := m.CloudConfigRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudConfigRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSpec) validateClusterConfig(formats strfmt.Registry) error {

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

func (m *V1SpectroClusterSpec) validateClusterProfileTemplates(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterProfileTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterProfileTemplates); i++ {
		if swag.IsZero(m.ClusterProfileTemplates[i]) { // not required
			continue
		}

		if m.ClusterProfileTemplates[i] != nil {
			if err := m.ClusterProfileTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterProfileTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var v1SpectroClusterSpecTypeClusterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PureManage","AlloyMonitor","AlloyAssist","AlloyExtend"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SpectroClusterSpecTypeClusterTypePropEnum = append(v1SpectroClusterSpecTypeClusterTypePropEnum, v)
	}
}

const (

	// V1SpectroClusterSpecClusterTypePureManage captures enum value "PureManage"
	V1SpectroClusterSpecClusterTypePureManage string = "PureManage"

	// V1SpectroClusterSpecClusterTypeAlloyMonitor captures enum value "AlloyMonitor"
	V1SpectroClusterSpecClusterTypeAlloyMonitor string = "AlloyMonitor"

	// V1SpectroClusterSpecClusterTypeAlloyAssist captures enum value "AlloyAssist"
	V1SpectroClusterSpecClusterTypeAlloyAssist string = "AlloyAssist"

	// V1SpectroClusterSpecClusterTypeAlloyExtend captures enum value "AlloyExtend"
	V1SpectroClusterSpecClusterTypeAlloyExtend string = "AlloyExtend"
)

// prop value enum
func (m *V1SpectroClusterSpec) validateClusterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1SpectroClusterSpecTypeClusterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1SpectroClusterSpec) validateClusterType(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterType) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterTypeEnum("clusterType", "body", m.ClusterType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterSpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}