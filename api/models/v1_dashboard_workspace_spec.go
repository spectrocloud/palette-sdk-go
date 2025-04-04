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

// V1DashboardWorkspaceSpec Workspace spec summary
//
// swagger:model v1DashboardWorkspaceSpec
type V1DashboardWorkspaceSpec struct {

	// cluster refs
	// Unique: true
	ClusterRefs []*V1DashboardWorkspaceClusterRef `json:"clusterRefs"`

	// namespaces
	// Unique: true
	Namespaces []string `json:"namespaces"`

	// quota
	Quota *V1DashboardWorkspaceQuota `json:"quota,omitempty"`
}

// Validate validates this v1 dashboard workspace spec
func (m *V1DashboardWorkspaceSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterRefs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DashboardWorkspaceSpec) validateClusterRefs(formats strfmt.Registry) error {
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

func (m *V1DashboardWorkspaceSpec) validateNamespaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespaces) { // not required
		return nil
	}

	if err := validate.UniqueItems("namespaces", "body", m.Namespaces); err != nil {
		return err
	}

	return nil
}

func (m *V1DashboardWorkspaceSpec) validateQuota(formats strfmt.Registry) error {
	if swag.IsZero(m.Quota) { // not required
		return nil
	}

	if m.Quota != nil {
		if err := m.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 dashboard workspace spec based on the context it is used
func (m *V1DashboardWorkspaceSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterRefs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DashboardWorkspaceSpec) contextValidateClusterRefs(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1DashboardWorkspaceSpec) contextValidateQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.Quota != nil {

		if swag.IsZero(m.Quota) { // not required
			return nil
		}

		if err := m.Quota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1DashboardWorkspaceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DashboardWorkspaceSpec) UnmarshalBinary(b []byte) error {
	var res V1DashboardWorkspaceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
