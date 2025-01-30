// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1EdgeNativeMachinePoolHostEntity v1 edge native machine pool host entity
//
// swagger:model v1EdgeNativeMachinePoolHostEntity
type V1EdgeNativeMachinePoolHostEntity struct {

	// Edge host name
	HostName string `json:"hostName,omitempty"`

	// Edge host id
	// Required: true
	HostUID *string `json:"hostUid"`

	// Edge native nic
	Nic *V1Nic `json:"nic,omitempty"`

	// Deprecated - Edge host nic name
	NicName string `json:"nicName,omitempty"`

	// Deprecated - Edge host static IP
	StaticIP string `json:"staticIP,omitempty"`

	// Set the edgehost candidate priority as primary or secondary, if the edgehost is nominated as two node candidate
	// Enum: [primary secondary]
	TwoNodeCandidatePriority string `json:"twoNodeCandidatePriority,omitempty"`
}

// Validate validates this v1 edge native machine pool host entity
func (m *V1EdgeNativeMachinePoolHostEntity) Validate(formats strfmt.Registry) error {
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

func (m *V1EdgeNativeMachinePoolHostEntity) validateHostUID(formats strfmt.Registry) error {

	if err := validate.Required("hostUid", "body", m.HostUID); err != nil {
		return err
	}

	return nil
}

func (m *V1EdgeNativeMachinePoolHostEntity) validateNic(formats strfmt.Registry) error {

	if swag.IsZero(m.Nic) { // not required
		return nil
	}

	if m.Nic != nil {
		if err := m.Nic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nic")
			}
			return err
		}
	}

	return nil
}

var v1EdgeNativeMachinePoolHostEntityTypeTwoNodeCandidatePriorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","secondary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1EdgeNativeMachinePoolHostEntityTypeTwoNodeCandidatePriorityPropEnum = append(v1EdgeNativeMachinePoolHostEntityTypeTwoNodeCandidatePriorityPropEnum, v)
	}
}

const (

	// V1EdgeNativeMachinePoolHostEntityTwoNodeCandidatePriorityPrimary captures enum value "primary"
	V1EdgeNativeMachinePoolHostEntityTwoNodeCandidatePriorityPrimary string = "primary"

	// V1EdgeNativeMachinePoolHostEntityTwoNodeCandidatePrioritySecondary captures enum value "secondary"
	V1EdgeNativeMachinePoolHostEntityTwoNodeCandidatePrioritySecondary string = "secondary"
)

// prop value enum
func (m *V1EdgeNativeMachinePoolHostEntity) validateTwoNodeCandidatePriorityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1EdgeNativeMachinePoolHostEntityTypeTwoNodeCandidatePriorityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1EdgeNativeMachinePoolHostEntity) validateTwoNodeCandidatePriority(formats strfmt.Registry) error {

	if swag.IsZero(m.TwoNodeCandidatePriority) { // not required
		return nil
	}

	// value enum
	if err := m.validateTwoNodeCandidatePriorityEnum("twoNodeCandidatePriority", "body", m.TwoNodeCandidatePriority); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeNativeMachinePoolHostEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeNativeMachinePoolHostEntity) UnmarshalBinary(b []byte) error {
	var res V1EdgeNativeMachinePoolHostEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
