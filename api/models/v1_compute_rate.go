// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ComputeRate Compute estimated rate information
//
// swagger:model v1ComputeRate
type V1ComputeRate struct {

	// rate
	Rate float64 `json:"rate"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 compute rate
func (m *V1ComputeRate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 compute rate based on context it is used
func (m *V1ComputeRate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ComputeRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ComputeRate) UnmarshalBinary(b []byte) error {
	var res V1ComputeRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
