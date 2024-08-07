// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1FilterVersionString v1 filter version string
//
// swagger:model v1FilterVersionString
type V1FilterVersionString struct {

	// begins with
	BeginsWith *string `json:"beginsWith,omitempty"`

	// eq
	Eq *string `json:"eq,omitempty"`

	// gt
	Gt *string `json:"gt,omitempty"`

	// lt
	Lt *string `json:"lt,omitempty"`

	// ne
	Ne *string `json:"ne,omitempty"`
}

// Validate validates this v1 filter version string
func (m *V1FilterVersionString) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1FilterVersionString) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FilterVersionString) UnmarshalBinary(b []byte) error {
	var res V1FilterVersionString
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
