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

// V1ClusterWorkloadStatefulSets Cluster workload statefulsets summary
//
// swagger:model v1ClusterWorkloadStatefulSets
type V1ClusterWorkloadStatefulSets struct {

	// stateful sets
	StatefulSets []*V1ClusterWorkloadStatefulSet `json:"statefulSets"`
}

// Validate validates this v1 cluster workload stateful sets
func (m *V1ClusterWorkloadStatefulSets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatefulSets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterWorkloadStatefulSets) validateStatefulSets(formats strfmt.Registry) error {

	if swag.IsZero(m.StatefulSets) { // not required
		return nil
	}

	for i := 0; i < len(m.StatefulSets); i++ {
		if swag.IsZero(m.StatefulSets[i]) { // not required
			continue
		}

		if m.StatefulSets[i] != nil {
			if err := m.StatefulSets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statefulSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterWorkloadStatefulSets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterWorkloadStatefulSets) UnmarshalBinary(b []byte) error {
	var res V1ClusterWorkloadStatefulSets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}