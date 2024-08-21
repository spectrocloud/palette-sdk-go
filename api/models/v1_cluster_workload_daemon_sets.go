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

// V1ClusterWorkloadDaemonSets Cluster workload daemonset summary
//
// swagger:model v1ClusterWorkloadDaemonSets
type V1ClusterWorkloadDaemonSets struct {

	// daemon sets
	DaemonSets []*V1ClusterWorkloadDaemonSet `json:"daemonSets"`
}

// Validate validates this v1 cluster workload daemon sets
func (m *V1ClusterWorkloadDaemonSets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDaemonSets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterWorkloadDaemonSets) validateDaemonSets(formats strfmt.Registry) error {

	if swag.IsZero(m.DaemonSets) { // not required
		return nil
	}

	for i := 0; i < len(m.DaemonSets); i++ {
		if swag.IsZero(m.DaemonSets[i]) { // not required
			continue
		}

		if m.DaemonSets[i] != nil {
			if err := m.DaemonSets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("daemonSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterWorkloadDaemonSets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterWorkloadDaemonSets) UnmarshalBinary(b []byte) error {
	var res V1ClusterWorkloadDaemonSets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}