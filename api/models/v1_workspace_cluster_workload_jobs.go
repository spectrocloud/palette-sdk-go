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

// V1WorkspaceClusterWorkloadJobs Workspace cluster workload jobs summary
//
// swagger:model v1WorkspaceClusterWorkloadJobs
type V1WorkspaceClusterWorkloadJobs struct {

	// jobs
	Jobs []*V1ClusterWorkloadJob `json:"jobs"`

	// metadata
	Metadata *V1RelatedObject `json:"metadata,omitempty"`
}

// Validate validates this v1 workspace cluster workload jobs
func (m *V1WorkspaceClusterWorkloadJobs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceClusterWorkloadJobs) validateJobs(formats strfmt.Registry) error {

	if swag.IsZero(m.Jobs) { // not required
		return nil
	}

	for i := 0; i < len(m.Jobs); i++ {
		if swag.IsZero(m.Jobs[i]) { // not required
			continue
		}

		if m.Jobs[i] != nil {
			if err := m.Jobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1WorkspaceClusterWorkloadJobs) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspaceClusterWorkloadJobs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceClusterWorkloadJobs) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceClusterWorkloadJobs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
