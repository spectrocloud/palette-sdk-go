// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1FilterString v1 filter string
//
// swagger:model v1FilterString
type V1FilterString struct {

	// begins with
	BeginsWith *string `json:"beginsWith,omitempty"`

	// contains
	Contains *string `json:"contains,omitempty"`

	// eq
	Eq *string `json:"eq,omitempty"`

	// ignore case
	IgnoreCase *bool `json:"ignoreCase,omitempty"`

	// ne
	Ne *string `json:"ne,omitempty"`
}

// Validate validates this v1 filter string
func (m *V1FilterString) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1FilterString) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FilterString) UnmarshalBinary(b []byte) error {
	var res V1FilterString
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
