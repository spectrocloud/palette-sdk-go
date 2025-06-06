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

// V1EdgeClusterObjectEntity Object identity meta of the cluster
//
// swagger:model v1EdgeClusterObjectEntity
type V1EdgeClusterObjectEntity struct {

	// In case of hybrid edge clusters, it provides the object identity meta of hybrid cluster
	HybridCluster *V1HybridClusterMeta `json:"hybridCluster,omitempty"`

	// Name of the cluster
	Name string `json:"name,omitempty"`

	// Uid of the cluster
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 edge cluster object entity
func (m *V1EdgeClusterObjectEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHybridCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeClusterObjectEntity) validateHybridCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.HybridCluster) { // not required
		return nil
	}

	if m.HybridCluster != nil {
		if err := m.HybridCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hybridCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hybridCluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 edge cluster object entity based on the context it is used
func (m *V1EdgeClusterObjectEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHybridCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeClusterObjectEntity) contextValidateHybridCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.HybridCluster != nil {

		if swag.IsZero(m.HybridCluster) { // not required
			return nil
		}

		if err := m.HybridCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hybridCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hybridCluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeClusterObjectEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeClusterObjectEntity) UnmarshalBinary(b []byte) error {
	var res V1EdgeClusterObjectEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
