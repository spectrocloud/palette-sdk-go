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

// V1ClusterProfilePacksManifests Cluster profile pack manifests
//
// swagger:model v1ClusterProfilePacksManifests
type V1ClusterProfilePacksManifests struct {

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// spec
	Spec *V1ClusterProfilePacksManifestsSpec `json:"spec,omitempty"`
}

// Validate validates this v1 cluster profile packs manifests
func (m *V1ClusterProfilePacksManifests) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterProfilePacksManifests) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterProfilePacksManifests) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterProfilePacksManifests) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterProfilePacksManifests) UnmarshalBinary(b []byte) error {
	var res V1ClusterProfilePacksManifests
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1ClusterProfilePacksManifestsSpec v1 cluster profile packs manifests spec
//
// swagger:model V1ClusterProfilePacksManifestsSpec
type V1ClusterProfilePacksManifestsSpec struct {

	// Cluster profile packs array
	// Unique: true
	Packs []*V1ClusterProfilePackManifests `json:"packs"`
}

// Validate validates this v1 cluster profile packs manifests spec
func (m *V1ClusterProfilePacksManifestsSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePacks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterProfilePacksManifestsSpec) validatePacks(formats strfmt.Registry) error {

	if swag.IsZero(m.Packs) { // not required
		return nil
	}

	if err := validate.UniqueItems("spec"+"."+"packs", "body", m.Packs); err != nil {
		return err
	}

	for i := 0; i < len(m.Packs); i++ {
		if swag.IsZero(m.Packs[i]) { // not required
			continue
		}

		if m.Packs[i] != nil {
			if err := m.Packs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "packs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterProfilePacksManifestsSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterProfilePacksManifestsSpec) UnmarshalBinary(b []byte) error {
	var res V1ClusterProfilePacksManifestsSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
