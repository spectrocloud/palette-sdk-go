// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PlanLimitSpec Monthly Plan Limit spec
//
// swagger:model v1PlanLimitSpec
type V1PlanLimitSpec struct {

	// cpu cores hours
	CPUCoreHours int64 `json:"cpuCoreHours"`

	// overage limit in percentage
	OverageLimitPercentage *int8 `json:"overageLimitPercentage"`

	// warning limit in percentage
	WarnLimitPercentage *int8 `json:"warnLimitPercentage"`
}

// Validate validates this v1 plan limit spec
func (m *V1PlanLimitSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PlanLimitSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PlanLimitSpec) UnmarshalBinary(b []byte) error {
	var res V1PlanLimitSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
