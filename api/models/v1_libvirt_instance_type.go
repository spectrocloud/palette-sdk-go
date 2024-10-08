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

// V1LibvirtInstanceType LibvirtInstanceType defines the instance configuration for a virtual machine
//
// swagger:model v1LibvirtInstanceType
type V1LibvirtInstanceType struct {

	// Defines CPU Passthrough Spec. A not null value enables CPU Passthrough for the libvirt domain. Further cache passthrough can be enabled with the CPU passthrough spec.
	CPUPassthroughSpec *V1CPUPassthroughSpec `json:"cpuPassthroughSpec,omitempty"`

	// CPUSet defines cpuset for an instance which allows allocation specific set of cpus E.g  cpuset="1-4,^3,6" See https://libvirt.org/formatdomain.html#cpu-allocation
	Cpuset string `json:"cpuset,omitempty"`

	// GPU configuration
	GpuConfig *V1GPUConfig `json:"gpuConfig,omitempty"`

	// MemoryinMB is the memory in megabytes
	// Required: true
	MemoryInMB *int32 `json:"memoryInMB"`

	// NumCPUs is the number of CPUs
	// Required: true
	NumCPUs *int32 `json:"numCPUs"`
}

// Validate validates this v1 libvirt instance type
func (m *V1LibvirtInstanceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUPassthroughSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpuConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryInMB(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumCPUs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1LibvirtInstanceType) validateCPUPassthroughSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.CPUPassthroughSpec) { // not required
		return nil
	}

	if m.CPUPassthroughSpec != nil {
		if err := m.CPUPassthroughSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpuPassthroughSpec")
			}
			return err
		}
	}

	return nil
}

func (m *V1LibvirtInstanceType) validateGpuConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.GpuConfig) { // not required
		return nil
	}

	if m.GpuConfig != nil {
		if err := m.GpuConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gpuConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1LibvirtInstanceType) validateMemoryInMB(formats strfmt.Registry) error {

	if err := validate.Required("memoryInMB", "body", m.MemoryInMB); err != nil {
		return err
	}

	return nil
}

func (m *V1LibvirtInstanceType) validateNumCPUs(formats strfmt.Registry) error {

	if err := validate.Required("numCPUs", "body", m.NumCPUs); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1LibvirtInstanceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LibvirtInstanceType) UnmarshalBinary(b []byte) error {
	var res V1LibvirtInstanceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
