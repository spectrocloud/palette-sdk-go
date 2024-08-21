// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterWorkloadJobStatus Cluster workload job status
//
// swagger:model v1ClusterWorkloadJobStatus
type V1ClusterWorkloadJobStatus struct {

	// completion time
	// Format: date-time
	CompletionTime V1Time `json:"completionTime,omitempty"`

	// conditions
	Conditions []*V1ClusterWorkloadCondition `json:"conditions"`

	// start time
	// Format: date-time
	StartTime V1Time `json:"startTime,omitempty"`

	// succeeded
	Succeeded int32 `json:"succeeded,omitempty"`
}

// Validate validates this v1 cluster workload job status
func (m *V1ClusterWorkloadJobStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterWorkloadJobStatus) validateCompletionTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletionTime) { // not required
		return nil
	}

	if err := m.CompletionTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("completionTime")
		}
		return err
	}

	return nil
}

func (m *V1ClusterWorkloadJobStatus) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterWorkloadJobStatus) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := m.StartTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("startTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterWorkloadJobStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterWorkloadJobStatus) UnmarshalBinary(b []byte) error {
	var res V1ClusterWorkloadJobStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}