// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMAffinity Affinity is a group of affinity scheduling rules.
//
// swagger:model v1VmAffinity
type V1VMAffinity struct {

	// node affinity
	NodeAffinity *V1VMNodeAffinity `json:"nodeAffinity,omitempty"`

	// pod affinity
	PodAffinity *V1VMPodAffinity `json:"podAffinity,omitempty"`

	// pod anti affinity
	PodAntiAffinity *V1PodAntiAffinity `json:"podAntiAffinity,omitempty"`
}

// Validate validates this v1 Vm affinity
func (m *V1VMAffinity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeAffinity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodAffinity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodAntiAffinity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMAffinity) validateNodeAffinity(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeAffinity) { // not required
		return nil
	}

	if m.NodeAffinity != nil {
		if err := m.NodeAffinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeAffinity")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMAffinity) validatePodAffinity(formats strfmt.Registry) error {

	if swag.IsZero(m.PodAffinity) { // not required
		return nil
	}

	if m.PodAffinity != nil {
		if err := m.PodAffinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAffinity")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMAffinity) validatePodAntiAffinity(formats strfmt.Registry) error {

	if swag.IsZero(m.PodAntiAffinity) { // not required
		return nil
	}

	if m.PodAntiAffinity != nil {
		if err := m.PodAntiAffinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAntiAffinity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMAffinity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMAffinity) UnmarshalBinary(b []byte) error {
	var res V1VMAffinity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
