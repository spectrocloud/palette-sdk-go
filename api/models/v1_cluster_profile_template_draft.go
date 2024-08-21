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

// V1ClusterProfileTemplateDraft Cluster profile template spec
//
// swagger:model v1ClusterProfileTemplateDraft
type V1ClusterProfileTemplateDraft struct {

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// Cluster profile packs array
	// Unique: true
	Packs []*V1PackManifestEntity `json:"packs"`

	// type
	Type V1ProfileType `json:"type,omitempty"`
}

// Validate validates this v1 cluster profile template draft
func (m *V1ClusterProfileTemplateDraft) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterProfileTemplateDraft) validatePacks(formats strfmt.Registry) error {

	if swag.IsZero(m.Packs) { // not required
		return nil
	}

	if err := validate.UniqueItems("packs", "body", m.Packs); err != nil {
		return err
	}

	for i := 0; i < len(m.Packs); i++ {
		if swag.IsZero(m.Packs[i]) { // not required
			continue
		}

		if m.Packs[i] != nil {
			if err := m.Packs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterProfileTemplateDraft) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterProfileTemplateDraft) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterProfileTemplateDraft) UnmarshalBinary(b []byte) error {
	var res V1ClusterProfileTemplateDraft
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}