// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ManifestSpec Manifest spec
//
// swagger:model v1ManifestSpec
type V1ManifestSpec struct {

	// draft
	Draft *V1ManifestData `json:"draft,omitempty"`

	// published
	Published *V1ManifestData `json:"published,omitempty"`
}

// Validate validates this v1 manifest spec
func (m *V1ManifestSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDraft(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublished(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ManifestSpec) validateDraft(formats strfmt.Registry) error {

	if swag.IsZero(m.Draft) { // not required
		return nil
	}

	if m.Draft != nil {
		if err := m.Draft.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("draft")
			}
			return err
		}
	}

	return nil
}

func (m *V1ManifestSpec) validatePublished(formats strfmt.Registry) error {

	if swag.IsZero(m.Published) { // not required
		return nil
	}

	if m.Published != nil {
		if err := m.Published.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("published")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ManifestSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ManifestSpec) UnmarshalBinary(b []byte) error {
	var res V1ManifestSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
