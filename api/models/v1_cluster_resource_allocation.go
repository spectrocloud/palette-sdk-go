// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterResourceAllocation Workspace resource allocation
//
// swagger:model v1ClusterResourceAllocation
type V1ClusterResourceAllocation struct {

	// cluster Uid
	ClusterUID string `json:"clusterUid,omitempty"`

	// resource allocation
	ResourceAllocation *V1WorkspaceResourceAllocation `json:"resourceAllocation,omitempty"`
}

// Validate validates this v1 cluster resource allocation
func (m *V1ClusterResourceAllocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceAllocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterResourceAllocation) validateResourceAllocation(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceAllocation) { // not required
		return nil
	}

	if m.ResourceAllocation != nil {
		if err := m.ResourceAllocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceAllocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceAllocation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 cluster resource allocation based on the context it is used
func (m *V1ClusterResourceAllocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceAllocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterResourceAllocation) contextValidateResourceAllocation(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceAllocation != nil {

		if swag.IsZero(m.ResourceAllocation) { // not required
			return nil
		}

		if err := m.ResourceAllocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceAllocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceAllocation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterResourceAllocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterResourceAllocation) UnmarshalBinary(b []byte) error {
	var res V1ClusterResourceAllocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
