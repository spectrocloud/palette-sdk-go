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

// V1VsphereInstanceType v1 vsphere instance type
//
// swagger:model v1VsphereInstanceType
type V1VsphereInstanceType struct {

	// DiskGiB is the size of a virtual machine's disk, in GiB. Defaults to the analogue property value in the template from which this machine is cloned.
	// Required: true
	DiskGiB *int32 `json:"diskGiB"`

	// MemoryMiB is the size of a virtual machine's memory, in MiB. Defaults to the analogue property value in the template from which this machine is cloned.
	// Required: true
	MemoryMiB *int64 `json:"memoryMiB"`

	// NumCPUs is the number of virtual processors in a virtual machine. Defaults to the analogue property value in the template from which this machine is cloned.
	// Required: true
	NumCPUs *int32 `json:"numCPUs"`
}

// Validate validates this v1 vsphere instance type
func (m *V1VsphereInstanceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiskGiB(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryMiB(formats); err != nil {
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

func (m *V1VsphereInstanceType) validateDiskGiB(formats strfmt.Registry) error {

	if err := validate.Required("diskGiB", "body", m.DiskGiB); err != nil {
		return err
	}

	return nil
}

func (m *V1VsphereInstanceType) validateMemoryMiB(formats strfmt.Registry) error {

	if err := validate.Required("memoryMiB", "body", m.MemoryMiB); err != nil {
		return err
	}

	return nil
}

func (m *V1VsphereInstanceType) validateNumCPUs(formats strfmt.Registry) error {

	if err := validate.Required("numCPUs", "body", m.NumCPUs); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VsphereInstanceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VsphereInstanceType) UnmarshalBinary(b []byte) error {
	var res V1VsphereInstanceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
