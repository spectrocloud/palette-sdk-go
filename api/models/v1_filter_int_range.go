// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1FilterIntRange v1 filter int range
//
// swagger:model v1FilterIntRange
type V1FilterIntRange struct {

	// eq
	Eq *int32 `json:"eq,omitempty"`

	// gt
	Gt *int32 `json:"gt,omitempty"`

	// gte
	Gte *int32 `json:"gte,omitempty"`

	// lt
	Lt *int32 `json:"lt,omitempty"`

	// lte
	Lte *int32 `json:"lte,omitempty"`

	// ne
	Ne *int32 `json:"ne,omitempty"`
}

// Validate validates this v1 filter int range
func (m *V1FilterIntRange) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1FilterIntRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FilterIntRange) UnmarshalBinary(b []byte) error {
	var res V1FilterIntRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
