// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1WorkspaceClusterNamespace Workspace cluster namespace
//
// swagger:model v1WorkspaceClusterNamespace
type V1WorkspaceClusterNamespace struct {

	// image
	Image *V1WorkspaceNamespaceImage `json:"image,omitempty"`

	// is regex
	IsRegex bool `json:"isRegex"`

	// name
	Name string `json:"name,omitempty"`

	// namespace resource allocation
	NamespaceResourceAllocation *V1WorkspaceNamespaceResourceAllocation `json:"namespaceResourceAllocation,omitempty"`
}

// Validate validates this v1 workspace cluster namespace
func (m *V1WorkspaceClusterNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceResourceAllocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceClusterNamespace) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *V1WorkspaceClusterNamespace) validateNamespaceResourceAllocation(formats strfmt.Registry) error {

	if swag.IsZero(m.NamespaceResourceAllocation) { // not required
		return nil
	}

	if m.NamespaceResourceAllocation != nil {
		if err := m.NamespaceResourceAllocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespaceResourceAllocation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspaceClusterNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceClusterNamespace) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceClusterNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
