// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1GenericMachinePoolConfig v1 generic machine pool config
//
// swagger:model v1GenericMachinePoolConfig
type V1GenericMachinePoolConfig struct {

	// instance type
	InstanceType string `json:"instanceType,omitempty"`

	// whether this pool is for control plane
	// Required: true
	IsControlPlane *bool `json:"isControlPlane"`

	// labels for this pool, example: control-plane/worker, gpu, windows
	Labels []string `json:"labels"`

	// name
	Name string `json:"name,omitempty"`

	// Size of root volume in GB. Default is 30GB
	RootDeviceSize int64 `json:"rootDeviceSize,omitempty"`

	// size of the pool, number of machines
	Size int32 `json:"size,omitempty"`

	// if IsControlPlane==true && useControlPlaneAsWorker==true, then will remove control plane taint this will not be used for worker pools
	UseControlPlaneAsWorker bool `json:"useControlPlaneAsWorker,omitempty"`
}

// Validate validates this v1 generic machine pool config
func (m *V1GenericMachinePoolConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsControlPlane(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GenericMachinePoolConfig) validateIsControlPlane(formats strfmt.Registry) error {

	if err := validate.Required("isControlPlane", "body", m.IsControlPlane); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GenericMachinePoolConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GenericMachinePoolConfig) UnmarshalBinary(b []byte) error {
	var res V1GenericMachinePoolConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}