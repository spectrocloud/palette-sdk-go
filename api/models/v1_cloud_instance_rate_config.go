// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CloudInstanceRateConfig Cloud instance rate config
//
// swagger:model v1CloudInstanceRateConfig
type V1CloudInstanceRateConfig struct {

	// compute rate proportion
	ComputeRateProportion float32 `json:"computeRateProportion,omitempty"`

	// memory rate proportion
	MemoryRateProportion float32 `json:"memoryRateProportion,omitempty"`
}

// Validate validates this v1 cloud instance rate config
func (m *V1CloudInstanceRateConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 cloud instance rate config based on context it is used
func (m *V1CloudInstanceRateConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CloudInstanceRateConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CloudInstanceRateConfig) UnmarshalBinary(b []byte) error {
	var res V1CloudInstanceRateConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
