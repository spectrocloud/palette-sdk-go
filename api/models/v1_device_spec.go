// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1DeviceSpec DeviceSpec defines the desired state of Device
//
// swagger:model v1DeviceSpec
type V1DeviceSpec struct {

	// Architecture type of the edge host
	// Enum: ["arm64","amd64"]
	ArchType *string `json:"archType,omitempty"`

	// cpu
	CPU *V1CPU `json:"cpu,omitempty"`

	// disks
	Disks []*V1Disk `json:"disks"`

	// gpus
	Gpus []*V1GPUDeviceSpec `json:"gpus"`

	// memory
	Memory *V1Memory `json:"memory,omitempty"`

	// nics
	Nics []*V1Nic `json:"nics"`

	// os
	Os *V1OS `json:"os,omitempty"`
}

// Validate validates this v1 device spec
func (m *V1DeviceSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1DeviceSpecTypeArchTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["arm64","amd64"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1DeviceSpecTypeArchTypePropEnum = append(v1DeviceSpecTypeArchTypePropEnum, v)
	}
}

const (

	// V1DeviceSpecArchTypeArm64 captures enum value "arm64"
	V1DeviceSpecArchTypeArm64 string = "arm64"

	// V1DeviceSpecArchTypeAmd64 captures enum value "amd64"
	V1DeviceSpecArchTypeAmd64 string = "amd64"
)

// prop value enum
func (m *V1DeviceSpec) validateArchTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1DeviceSpecTypeArchTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1DeviceSpec) validateArchType(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchType) { // not required
		return nil
	}

	// value enum
	if err := m.validateArchTypeEnum("archType", "body", *m.ArchType); err != nil {
		return err
	}

	return nil
}

func (m *V1DeviceSpec) validateCPU(formats strfmt.Registry) error {
	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *V1DeviceSpec) validateDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1DeviceSpec) validateGpus(formats strfmt.Registry) error {
	if swag.IsZero(m.Gpus) { // not required
		return nil
	}

	for i := 0; i < len(m.Gpus); i++ {
		if swag.IsZero(m.Gpus[i]) { // not required
			continue
		}

		if m.Gpus[i] != nil {
			if err := m.Gpus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1DeviceSpec) validateMemory(formats strfmt.Registry) error {
	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if m.Memory != nil {
		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memory")
			}
			return err
		}
	}

	return nil
}

func (m *V1DeviceSpec) validateNics(formats strfmt.Registry) error {
	if swag.IsZero(m.Nics) { // not required
		return nil
	}

	for i := 0; i < len(m.Nics); i++ {
		if swag.IsZero(m.Nics[i]) { // not required
			continue
		}

		if m.Nics[i] != nil {
			if err := m.Nics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1DeviceSpec) validateOs(formats strfmt.Registry) error {
	if swag.IsZero(m.Os) { // not required
		return nil
	}

	if m.Os != nil {
		if err := m.Os.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("os")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("os")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 device spec based on the context it is used
func (m *V1DeviceSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCPU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGpus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DeviceSpec) contextValidateCPU(ctx context.Context, formats strfmt.Registry) error {

	if m.CPU != nil {

		if swag.IsZero(m.CPU) { // not required
			return nil
		}

		if err := m.CPU.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *V1DeviceSpec) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {

			if swag.IsZero(m.Disks[i]) { // not required
				return nil
			}

			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1DeviceSpec) contextValidateGpus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Gpus); i++ {

		if m.Gpus[i] != nil {

			if swag.IsZero(m.Gpus[i]) { // not required
				return nil
			}

			if err := m.Gpus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1DeviceSpec) contextValidateMemory(ctx context.Context, formats strfmt.Registry) error {

	if m.Memory != nil {

		if swag.IsZero(m.Memory) { // not required
			return nil
		}

		if err := m.Memory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memory")
			}
			return err
		}
	}

	return nil
}

func (m *V1DeviceSpec) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nics); i++ {

		if m.Nics[i] != nil {

			if swag.IsZero(m.Nics[i]) { // not required
				return nil
			}

			if err := m.Nics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1DeviceSpec) contextValidateOs(ctx context.Context, formats strfmt.Registry) error {

	if m.Os != nil {

		if swag.IsZero(m.Os) { // not required
			return nil
		}

		if err := m.Os.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("os")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("os")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1DeviceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DeviceSpec) UnmarshalBinary(b []byte) error {
	var res V1DeviceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
