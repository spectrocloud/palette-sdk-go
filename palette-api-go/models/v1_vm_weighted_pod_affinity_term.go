// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1VMWeightedPodAffinityTerm The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)
//
// swagger:model v1VmWeightedPodAffinityTerm
type V1VMWeightedPodAffinityTerm struct {

	// pod affinity term
	// Required: true
	PodAffinityTerm *V1VMPodAffinityTerm `json:"podAffinityTerm"`

	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	// Required: true
	Weight *int32 `json:"weight"`
}

// Validate validates this v1 Vm weighted pod affinity term
func (m *V1VMWeightedPodAffinityTerm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePodAffinityTerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMWeightedPodAffinityTerm) validatePodAffinityTerm(formats strfmt.Registry) error {

	if err := validate.Required("podAffinityTerm", "body", m.PodAffinityTerm); err != nil {
		return err
	}

	if m.PodAffinityTerm != nil {
		if err := m.PodAffinityTerm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podAffinityTerm")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMWeightedPodAffinityTerm) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("weight", "body", m.Weight); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMWeightedPodAffinityTerm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMWeightedPodAffinityTerm) UnmarshalBinary(b []byte) error {
	var res V1VMWeightedPodAffinityTerm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
