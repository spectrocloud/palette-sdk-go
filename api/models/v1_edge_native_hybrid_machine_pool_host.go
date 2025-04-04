// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1EdgeNativeHybridMachinePoolHost v1EdgeNativeHybridMachinePoolHostEntity defines Edge Native machine pool's host configuration
//
// swagger:model v1EdgeNativeHybridMachinePoolHost
type V1EdgeNativeHybridMachinePoolHost struct {

	// Edge host name
	HostName string `json:"hostName,omitempty"`

	// Edge host id
	// Required: true
	HostUID *string `json:"hostUid"`

	// Edge native nic
	Nic *V1Nic `json:"nic,omitempty"`

	// Set the Edge Host candidate priority as primary or secondary, if the Edge Host is nominated as two node candidate
	// Enum: ["primary","secondary"]
	TwoNodeCandidatePriority string `json:"twoNodeCandidatePriority,omitempty"`

	// Vpn server IP
	VpnServerIP string `json:"vpnServerIp,omitempty"`
}

// Validate validates this v1 edge native hybrid machine pool host
func (m *V1EdgeNativeHybridMachinePoolHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTwoNodeCandidatePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeNativeHybridMachinePoolHost) validateHostUID(formats strfmt.Registry) error {

	if err := validate.Required("hostUid", "body", m.HostUID); err != nil {
		return err
	}

	return nil
}

func (m *V1EdgeNativeHybridMachinePoolHost) validateNic(formats strfmt.Registry) error {
	if swag.IsZero(m.Nic) { // not required
		return nil
	}

	if m.Nic != nil {
		if err := m.Nic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nic")
			}
			return err
		}
	}

	return nil
}

var v1EdgeNativeHybridMachinePoolHostTypeTwoNodeCandidatePriorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","secondary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1EdgeNativeHybridMachinePoolHostTypeTwoNodeCandidatePriorityPropEnum = append(v1EdgeNativeHybridMachinePoolHostTypeTwoNodeCandidatePriorityPropEnum, v)
	}
}

const (

	// V1EdgeNativeHybridMachinePoolHostTwoNodeCandidatePriorityPrimary captures enum value "primary"
	V1EdgeNativeHybridMachinePoolHostTwoNodeCandidatePriorityPrimary string = "primary"

	// V1EdgeNativeHybridMachinePoolHostTwoNodeCandidatePrioritySecondary captures enum value "secondary"
	V1EdgeNativeHybridMachinePoolHostTwoNodeCandidatePrioritySecondary string = "secondary"
)

// prop value enum
func (m *V1EdgeNativeHybridMachinePoolHost) validateTwoNodeCandidatePriorityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1EdgeNativeHybridMachinePoolHostTypeTwoNodeCandidatePriorityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1EdgeNativeHybridMachinePoolHost) validateTwoNodeCandidatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.TwoNodeCandidatePriority) { // not required
		return nil
	}

	// value enum
	if err := m.validateTwoNodeCandidatePriorityEnum("twoNodeCandidatePriority", "body", m.TwoNodeCandidatePriority); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 edge native hybrid machine pool host based on the context it is used
func (m *V1EdgeNativeHybridMachinePoolHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeNativeHybridMachinePoolHost) contextValidateNic(ctx context.Context, formats strfmt.Registry) error {

	if m.Nic != nil {

		if swag.IsZero(m.Nic) { // not required
			return nil
		}

		if err := m.Nic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nic")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeNativeHybridMachinePoolHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeNativeHybridMachinePoolHost) UnmarshalBinary(b []byte) error {
	var res V1EdgeNativeHybridMachinePoolHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
