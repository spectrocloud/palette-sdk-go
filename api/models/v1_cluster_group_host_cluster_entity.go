// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ClusterGroupHostClusterEntity Clusters and clusters config of cluster group
//
// swagger:model v1ClusterGroupHostClusterEntity
type V1ClusterGroupHostClusterEntity struct {

	// cluster refs
	// Unique: true
	ClusterRefs []*V1ClusterGroupClusterRef `json:"clusterRefs"`

	// clusters config
	ClustersConfig *V1ClusterGroupClustersConfig `json:"clustersConfig,omitempty"`
}

// Validate validates this v1 cluster group host cluster entity
func (m *V1ClusterGroupHostClusterEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterRefs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustersConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterGroupHostClusterEntity) validateClusterRefs(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterRefs) { // not required
		return nil
	}

	if err := validate.UniqueItems("clusterRefs", "body", m.ClusterRefs); err != nil {
		return err
	}

	for i := 0; i < len(m.ClusterRefs); i++ {
		if swag.IsZero(m.ClusterRefs[i]) { // not required
			continue
		}

		if m.ClusterRefs[i] != nil {
			if err := m.ClusterRefs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterRefs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterRefs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterGroupHostClusterEntity) validateClustersConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ClustersConfig) { // not required
		return nil
	}

	if m.ClustersConfig != nil {
		if err := m.ClustersConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clustersConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clustersConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 cluster group host cluster entity based on the context it is used
func (m *V1ClusterGroupHostClusterEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterRefs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClustersConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterGroupHostClusterEntity) contextValidateClusterRefs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterRefs); i++ {

		if m.ClusterRefs[i] != nil {

			if swag.IsZero(m.ClusterRefs[i]) { // not required
				return nil
			}

			if err := m.ClusterRefs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterRefs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterRefs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterGroupHostClusterEntity) contextValidateClustersConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ClustersConfig != nil {

		if swag.IsZero(m.ClustersConfig) { // not required
			return nil
		}

		if err := m.ClustersConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clustersConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clustersConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterGroupHostClusterEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterGroupHostClusterEntity) UnmarshalBinary(b []byte) error {
	var res V1ClusterGroupHostClusterEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
