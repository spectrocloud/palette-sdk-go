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
)

// V1PackImportEntity Pack import request payload
//
// swagger:model v1PackImportEntity
type V1PackImportEntity struct {

	// Pack layer [ "os", "k8s", "cni", "csi", "addon" ]
	Layer string `json:"layer,omitempty"`

	// Pack manifests array
	Manifests []*V1PackManifestImportEntity `json:"manifests"`

	// Pack name
	Name string `json:"name,omitempty"`

	// registry
	Registry *V1PackRegistryImportEntity `json:"registry,omitempty"`

	// Pack version tag
	Tag string `json:"tag,omitempty"`

	// Pack type [ "spectro", "helm", "manifest", "oci" ]
	Type string `json:"type,omitempty"`

	// Pack values are the customizable configurations for the pack
	Values string `json:"values,omitempty"`

	// Pack version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 pack import entity
func (m *V1PackImportEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManifests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackImportEntity) validateManifests(formats strfmt.Registry) error {
	if swag.IsZero(m.Manifests) { // not required
		return nil
	}

	for i := 0; i < len(m.Manifests); i++ {
		if swag.IsZero(m.Manifests[i]) { // not required
			continue
		}

		if m.Manifests[i] != nil {
			if err := m.Manifests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("manifests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("manifests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PackImportEntity) validateRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.Registry) { // not required
		return nil
	}

	if m.Registry != nil {
		if err := m.Registry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 pack import entity based on the context it is used
func (m *V1PackImportEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateManifests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackImportEntity) contextValidateManifests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Manifests); i++ {

		if m.Manifests[i] != nil {

			if swag.IsZero(m.Manifests[i]) { // not required
				return nil
			}

			if err := m.Manifests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("manifests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("manifests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PackImportEntity) contextValidateRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.Registry != nil {

		if swag.IsZero(m.Registry) { // not required
			return nil
		}

		if err := m.Registry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PackImportEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackImportEntity) UnmarshalBinary(b []byte) error {
	var res V1PackImportEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
