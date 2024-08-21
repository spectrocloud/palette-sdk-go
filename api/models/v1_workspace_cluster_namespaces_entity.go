// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1WorkspaceClusterNamespacesEntity Workspace cluster namespaces update entity
//
// swagger:model v1WorkspaceClusterNamespacesEntity
type V1WorkspaceClusterNamespacesEntity struct {

	// cluster namespaces
	// Unique: true
	ClusterNamespaces []*V1WorkspaceClusterNamespace `json:"clusterNamespaces"`

	// cluster refs
	// Unique: true
	ClusterRefs []*V1WorkspaceClusterRef `json:"clusterRefs"`

	// quota
	Quota *V1WorkspaceQuota `json:"quota,omitempty"`
}

// Validate validates this v1 workspace cluster namespaces entity
func (m *V1WorkspaceClusterNamespacesEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterRefs(formats); err != nil {
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

func (m *V1WorkspaceClusterNamespacesEntity) validateClusterNamespaces(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNamespaces) { // not required
		return nil
	}

	if err := validate.UniqueItems("clusterNamespaces", "body", m.ClusterNamespaces); err != nil {
		return err
	}

	for i := 0; i < len(m.ClusterNamespaces); i++ {
		if swag.IsZero(m.ClusterNamespaces[i]) { // not required
			continue
		}

		if m.ClusterNamespaces[i] != nil {
			if err := m.ClusterNamespaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterNamespaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1WorkspaceClusterNamespacesEntity) validateClusterRefs(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1WorkspaceClusterNamespacesEntity) validateQuota(formats strfmt.Registry) error {

	if swag.IsZero(m.Quota) { // not required
		return nil
	}

	if m.Quota != nil {
		if err := m.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspaceClusterNamespacesEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceClusterNamespacesEntity) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceClusterNamespacesEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}