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

// V1PackConfigSpec v1 pack config spec
//
// swagger:model v1PackConfigSpec
type V1PackConfigSpec struct {

	// associated object
	AssociatedObject string `json:"associatedObject,omitempty"`

	// is values overridden
	IsValuesOverridden bool `json:"isValuesOverridden"`

	// manifests
	Manifests []*V1PackManifestRef `json:"manifests"`

	// name
	Name string `json:"name,omitempty"`

	// pack Uid
	PackUID string `json:"packUid,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// values
	Values string `json:"values,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 pack config spec
func (m *V1PackConfigSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManifests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackConfigSpec) validateManifests(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 pack config spec based on the context it is used
func (m *V1PackConfigSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateManifests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackConfigSpec) contextValidateManifests(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V1PackConfigSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackConfigSpec) UnmarshalBinary(b []byte) error {
	var res V1PackConfigSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
