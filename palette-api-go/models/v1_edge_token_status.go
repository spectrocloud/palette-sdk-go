// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EdgeTokenStatus Edge token status
//
// swagger:model v1EdgeTokenStatus
type V1EdgeTokenStatus struct {

	// Set to 'true', if the token is active
	IsActive bool `json:"isActive"`
}

// Validate validates this v1 edge token status
func (m *V1EdgeTokenStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeTokenStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeTokenStatus) UnmarshalBinary(b []byte) error {
	var res V1EdgeTokenStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
