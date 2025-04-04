// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ResourceUsageMeteringDataPoint min and max count for machines & edgehost for the given period
//
// swagger:model v1ResourceUsageMeteringDataPoint
type V1ResourceUsageMeteringDataPoint struct {

	// active edgehosts
	ActiveEdgehosts int64 `json:"activeEdgehosts,omitempty"`

	// active machines
	ActiveMachines int64 `json:"activeMachines,omitempty"`

	// max edgehosts
	MaxEdgehosts int64 `json:"maxEdgehosts,omitempty"`

	// max machines
	MaxMachines int64 `json:"maxMachines,omitempty"`
}

// Validate validates this v1 resource usage metering data point
func (m *V1ResourceUsageMeteringDataPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 resource usage metering data point based on context it is used
func (m *V1ResourceUsageMeteringDataPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ResourceUsageMeteringDataPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ResourceUsageMeteringDataPoint) UnmarshalBinary(b []byte) error {
	var res V1ResourceUsageMeteringDataPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
