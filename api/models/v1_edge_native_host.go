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

// V1EdgeNativeHost EdgeNativeHost is the underlying appliance
//
// swagger:model v1EdgeNativeHost
type V1EdgeNativeHost struct {

	// Is Edge host nominated as candidate
	IsCandidateCaption *bool `json:"IsCandidateCaption"`

	// CACert for TLS connections
	CaCert string `json:"caCert,omitempty"`

	// HostAddress is a FQDN or IP address of the Host
	// Required: true
	HostAddress *string `json:"hostAddress"`

	// Qualified name of host
	HostName string `json:"hostName,omitempty"`

	// HostUid is the ID of the EdgeHost
	// Required: true
	HostUID *string `json:"hostUid"`

	// Edge native nic
	Nic *V1Nic `json:"nic,omitempty"`

	// Deprecated. Edge host nic name
	NicName string `json:"nicName,omitempty"`

	// Deprecated. Edge host static IP
	StaticIP string `json:"staticIP,omitempty"`

	// Set the edgehost candidate priority as primary or secondary, if the edgehost is nominated as two node candidate
	// Enum: ["primary","secondary"]
	TwoNodeCandidatePriority string `json:"twoNodeCandidatePriority,omitempty"`
}

// Validate validates this v1 edge native host
func (m *V1EdgeNativeHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostAddress(formats); err != nil {
		res = append(res, err)
	}

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

func (m *V1EdgeNativeHost) validateHostAddress(formats strfmt.Registry) error {

	if err := validate.Required("hostAddress", "body", m.HostAddress); err != nil {
		return err
	}

	return nil
}

func (m *V1EdgeNativeHost) validateHostUID(formats strfmt.Registry) error {

	if err := validate.Required("hostUid", "body", m.HostUID); err != nil {
		return err
	}

	return nil
}

func (m *V1EdgeNativeHost) validateNic(formats strfmt.Registry) error {
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

var v1EdgeNativeHostTypeTwoNodeCandidatePriorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","secondary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1EdgeNativeHostTypeTwoNodeCandidatePriorityPropEnum = append(v1EdgeNativeHostTypeTwoNodeCandidatePriorityPropEnum, v)
	}
}

const (

	// V1EdgeNativeHostTwoNodeCandidatePriorityPrimary captures enum value "primary"
	V1EdgeNativeHostTwoNodeCandidatePriorityPrimary string = "primary"

	// V1EdgeNativeHostTwoNodeCandidatePrioritySecondary captures enum value "secondary"
	V1EdgeNativeHostTwoNodeCandidatePrioritySecondary string = "secondary"
)

// prop value enum
func (m *V1EdgeNativeHost) validateTwoNodeCandidatePriorityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1EdgeNativeHostTypeTwoNodeCandidatePriorityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1EdgeNativeHost) validateTwoNodeCandidatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.TwoNodeCandidatePriority) { // not required
		return nil
	}

	// value enum
	if err := m.validateTwoNodeCandidatePriorityEnum("twoNodeCandidatePriority", "body", m.TwoNodeCandidatePriority); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 edge native host based on the context it is used
func (m *V1EdgeNativeHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeNativeHost) contextValidateNic(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1EdgeNativeHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeNativeHost) UnmarshalBinary(b []byte) error {
	var res V1EdgeNativeHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
