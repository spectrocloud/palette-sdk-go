// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MetricPoint Metric Info
//
// swagger:model v1MetricPoint
type V1MetricPoint struct {

	// avg
	Avg float64 `json:"avg,omitempty"`

	// count
	Count int64 `json:"count,omitempty"`

	// max
	Max float64 `json:"max,omitempty"`

	// min
	Min float64 `json:"min,omitempty"`

	// sum
	Sum float64 `json:"sum,omitempty"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`

	// value
	Value float64 `json:"value"`
}

// Validate validates this v1 metric point
func (m *V1MetricPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 metric point based on context it is used
func (m *V1MetricPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MetricPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MetricPoint) UnmarshalBinary(b []byte) error {
	var res V1MetricPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
