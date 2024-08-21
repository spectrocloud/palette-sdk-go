// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CloudSpotPrice Spot price entity of a particular cloud type
//
// swagger:model v1CloudSpotPrice
type V1CloudSpotPrice struct {

	// Spot price of a resource for a particular cloud
	SpotPrice float64 `json:"spotPrice"`
}

// Validate validates this v1 cloud spot price
func (m *V1CloudSpotPrice) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CloudSpotPrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CloudSpotPrice) UnmarshalBinary(b []byte) error {
	var res V1CloudSpotPrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}