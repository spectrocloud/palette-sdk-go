// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpotMarketOptions SpotMarketOptions defines the options available to a user when configuring Machines to run on Spot instances. Most users should provide an empty struct.
//
// swagger:model v1SpotMarketOptions
type V1SpotMarketOptions struct {

	// MaxPrice defines the maximum price the user is willing to pay for Spot VM instances
	MaxPrice string `json:"maxPrice,omitempty"`
}

// Validate validates this v1 spot market options
func (m *V1SpotMarketOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 spot market options based on context it is used
func (m *V1SpotMarketOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SpotMarketOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpotMarketOptions) UnmarshalBinary(b []byte) error {
	var res V1SpotMarketOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
