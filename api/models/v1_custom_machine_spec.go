// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CustomMachineSpec Custom cloud VM definition spec
//
// swagger:model v1CustomMachineSpec
type V1CustomMachineSpec struct {

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// host name
	HostName string `json:"hostName,omitempty"`

	// image Id
	ImageID string `json:"imageId,omitempty"`

	// instance type
	InstanceType *V1CustomInstanceType `json:"instanceType,omitempty"`

	// nics
	Nics []*V1CustomNic `json:"nics"`

	// ssh key name
	SSHKeyName string `json:"sshKeyName,omitempty"`
}

// Validate validates this v1 custom machine spec
func (m *V1CustomMachineSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CustomMachineSpec) validateInstanceType(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceType) { // not required
		return nil
	}

	if m.InstanceType != nil {
		if err := m.InstanceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instanceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instanceType")
			}
			return err
		}
	}

	return nil
}

func (m *V1CustomMachineSpec) validateNics(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 custom machine spec based on the context it is used
func (m *V1CustomMachineSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstanceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CustomMachineSpec) contextValidateInstanceType(ctx context.Context, formats strfmt.Registry) error {

	if m.InstanceType != nil {

		if swag.IsZero(m.InstanceType) { // not required
			return nil
		}

		if err := m.InstanceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instanceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instanceType")
			}
			return err
		}
	}

	return nil
}

func (m *V1CustomMachineSpec) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V1CustomMachineSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CustomMachineSpec) UnmarshalBinary(b []byte) error {
	var res V1CustomMachineSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
