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
	"github.com/go-openapi/validate"
)

// V1VsphereMachineSpec vSphere cloud VM definition spec
//
// swagger:model v1VsphereMachineSpec
type V1VsphereMachineSpec struct {

	// images
	Images []*V1VsphereImage `json:"images"`

	// instance type
	InstanceType *V1VsphereInstanceType `json:"instanceType,omitempty"`

	// nics
	// Required: true
	Nics []*V1VsphereNic `json:"nics"`

	// NTPServers is a list of NTP servers to use instead of the machine image's default NTP server list.
	NtpServers []string `json:"ntpServers"`

	// Placement configuration
	// Required: true
	Placement *V1VspherePlacementConfig `json:"placement"`

	// VcenterServer is the address of the vSphere endpoint
	// Required: true
	VcenterServer *string `json:"vcenterServer"`
}

// Validate validates this v1 vsphere machine spec
func (m *V1VsphereMachineSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlacement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcenterServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VsphereMachineSpec) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VsphereMachineSpec) validateInstanceType(formats strfmt.Registry) error {
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

func (m *V1VsphereMachineSpec) validateNics(formats strfmt.Registry) error {

	if err := validate.Required("nics", "body", m.Nics); err != nil {
		return err
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

func (m *V1VsphereMachineSpec) validatePlacement(formats strfmt.Registry) error {

	if err := validate.Required("placement", "body", m.Placement); err != nil {
		return err
	}

	if m.Placement != nil {
		if err := m.Placement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("placement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("placement")
			}
			return err
		}
	}

	return nil
}

func (m *V1VsphereMachineSpec) validateVcenterServer(formats strfmt.Registry) error {

	if err := validate.Required("vcenterServer", "body", m.VcenterServer); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 vsphere machine spec based on the context it is used
func (m *V1VsphereMachineSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstanceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlacement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VsphereMachineSpec) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Images); i++ {

		if m.Images[i] != nil {

			if swag.IsZero(m.Images[i]) { // not required
				return nil
			}

			if err := m.Images[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VsphereMachineSpec) contextValidateInstanceType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1VsphereMachineSpec) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1VsphereMachineSpec) contextValidatePlacement(ctx context.Context, formats strfmt.Registry) error {

	if m.Placement != nil {

		if err := m.Placement.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("placement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("placement")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VsphereMachineSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VsphereMachineSpec) UnmarshalBinary(b []byte) error {
	var res V1VsphereMachineSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
