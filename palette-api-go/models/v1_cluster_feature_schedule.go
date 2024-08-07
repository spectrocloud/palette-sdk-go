// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterFeatureSchedule Cluster feature schedule
//
// swagger:model v1ClusterFeatureSchedule
type V1ClusterFeatureSchedule struct {

	// scheduled run time
	ScheduledRunTime string `json:"scheduledRunTime,omitempty"`
}

// Validate validates this v1 cluster feature schedule
func (m *V1ClusterFeatureSchedule) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterFeatureSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterFeatureSchedule) UnmarshalBinary(b []byte) error {
	var res V1ClusterFeatureSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
