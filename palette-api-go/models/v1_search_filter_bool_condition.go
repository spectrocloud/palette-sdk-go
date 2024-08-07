// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SearchFilterBoolCondition v1 search filter bool condition
//
// swagger:model v1SearchFilterBoolCondition
type V1SearchFilterBoolCondition struct {

	// value
	Value bool `json:"value,omitempty"`
}

// Validate validates this v1 search filter bool condition
func (m *V1SearchFilterBoolCondition) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SearchFilterBoolCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchFilterBoolCondition) UnmarshalBinary(b []byte) error {
	var res V1SearchFilterBoolCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
