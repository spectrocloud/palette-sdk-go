// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AwsHybridMachinePool Machine pool reference of cloud config of cluster deployed by hybrid cluster
//
// swagger:model v1AwsHybridMachinePool
type V1AwsHybridMachinePool struct {

	// pool cloud type
	PoolCloudType V1HybridPoolClusterCloudType `json:"poolCloudType,omitempty"`

	// Machine pool name
	PoolName string `json:"poolName,omitempty"`

	// Machine pool uid
	PoolUID string `json:"poolUid,omitempty"`
}

// Validate validates this v1 aws hybrid machine pool
func (m *V1AwsHybridMachinePool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePoolCloudType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsHybridMachinePool) validatePoolCloudType(formats strfmt.Registry) error {

	if swag.IsZero(m.PoolCloudType) { // not required
		return nil
	}

	if err := m.PoolCloudType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("poolCloudType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsHybridMachinePool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsHybridMachinePool) UnmarshalBinary(b []byte) error {
	var res V1AwsHybridMachinePool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
