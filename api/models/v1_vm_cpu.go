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

// V1VMCPU CPU allows specifying the CPU topology.
//
// swagger:model v1VmCpu
type V1VMCPU struct {

	// Cores specifies the number of cores inside the vmi. Must be a value greater or equal 1.
	Cores int64 `json:"cores,omitempty"`

	// DedicatedCPUPlacement requests the scheduler to place the VirtualMachineInstance on a node with enough dedicated pCPUs and pin the vCPUs to it.
	DedicatedCPUPlacement bool `json:"dedicatedCpuPlacement,omitempty"`

	// Features specifies the CPU features list inside the VMI.
	Features []*V1VMCPUFeature `json:"features"`

	// IsolateEmulatorThread requests one more dedicated pCPU to be allocated for the VMI to place the emulator thread on it.
	IsolateEmulatorThread bool `json:"isolateEmulatorThread,omitempty"`

	// Model specifies the CPU model inside the VMI. List of available models https://github.com/libvirt/libvirt/tree/master/src/cpu_map. It is possible to specify special cases like "host-passthrough" to get the same CPU as the node and "host-model" to get CPU closest to the node one. Defaults to host-model.
	Model string `json:"model,omitempty"`

	// numa
	Numa *V1VMNUMA `json:"numa,omitempty"`

	// realtime
	Realtime *V1VMRealtime `json:"realtime,omitempty"`

	// Sockets specifies the number of sockets inside the vmi. Must be a value greater or equal 1.
	Sockets int64 `json:"sockets,omitempty"`

	// Threads specifies the number of threads inside the vmi. Must be a value greater or equal 1.
	Threads int64 `json:"threads,omitempty"`
}

// Validate validates this v1 Vm Cpu
func (m *V1VMCPU) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNuma(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRealtime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMCPU) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	for i := 0; i < len(m.Features); i++ {
		if swag.IsZero(m.Features[i]) { // not required
			continue
		}

		if m.Features[i] != nil {
			if err := m.Features[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VMCPU) validateNuma(formats strfmt.Registry) error {

	if swag.IsZero(m.Numa) { // not required
		return nil
	}

	if m.Numa != nil {
		if err := m.Numa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("numa")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMCPU) validateRealtime(formats strfmt.Registry) error {

	if swag.IsZero(m.Realtime) { // not required
		return nil
	}

	if m.Realtime != nil {
		if err := m.Realtime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realtime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMCPU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMCPU) UnmarshalBinary(b []byte) error {
	var res V1VMCPU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
