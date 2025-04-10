// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpotVMOptions SpotVMOptions defines the options relevant to running the Machine on Spot VMs
//
// swagger:model v1SpotVMOptions
type V1SpotVMOptions struct {

	// MaxPrice defines the maximum price the user is willing to pay for Spot VM instances
	MaxPrice string `json:"maxPrice,omitempty"`
}

// Validate validates this v1 spot VM options
func (m *V1SpotVMOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 spot VM options based on context it is used
func (m *V1SpotVMOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SpotVMOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpotVMOptions) UnmarshalBinary(b []byte) error {
	var res V1SpotVMOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
