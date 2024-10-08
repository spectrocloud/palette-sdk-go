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

// V1ClusterWorkloadPodMetadata Cluster workload pod metadata
//
// swagger:model v1ClusterWorkloadPodMetadata
type V1ClusterWorkloadPodMetadata struct {

	// associated refs
	AssociatedRefs []*V1ClusterWorkloadRef `json:"associatedRefs"`

	// creation timestamp
	// Format: date-time
	CreationTimestamp V1Time `json:"creationTimestamp,omitempty"`

	// entity
	Entity *V1ClusterWorkloadRef `json:"entity,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// machine Uid
	MachineUID string `json:"machineUid,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// nodename
	Nodename string `json:"nodename,omitempty"`
}

// Validate validates this v1 cluster workload pod metadata
func (m *V1ClusterWorkloadPodMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssociatedRefs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterWorkloadPodMetadata) validateAssociatedRefs(formats strfmt.Registry) error {

	if swag.IsZero(m.AssociatedRefs) { // not required
		return nil
	}

	for i := 0; i < len(m.AssociatedRefs); i++ {
		if swag.IsZero(m.AssociatedRefs[i]) { // not required
			continue
		}

		if m.AssociatedRefs[i] != nil {
			if err := m.AssociatedRefs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("associatedRefs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterWorkloadPodMetadata) validateCreationTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := m.CreationTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("creationTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1ClusterWorkloadPodMetadata) validateEntity(formats strfmt.Registry) error {

	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	if m.Entity != nil {
		if err := m.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterWorkloadPodMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterWorkloadPodMetadata) UnmarshalBinary(b []byte) error {
	var res V1ClusterWorkloadPodMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
