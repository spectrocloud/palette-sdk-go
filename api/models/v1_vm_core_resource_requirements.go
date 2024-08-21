// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMCoreResourceRequirements ResourceRequirements describes the compute resource requirements.
//
// swagger:model v1VmCoreResourceRequirements
type V1VMCoreResourceRequirements struct {

	// Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Limits map[string]V1VMQuantity `json:"limits,omitempty"`

	// Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests map[string]V1VMQuantity `json:"requests,omitempty"`
}

// Validate validates this v1 Vm core resource requirements
func (m *V1VMCoreResourceRequirements) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMCoreResourceRequirements) validateLimits(formats strfmt.Registry) error {

	if swag.IsZero(m.Limits) { // not required
		return nil
	}

	for k := range m.Limits {

		if val, ok := m.Limits[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V1VMCoreResourceRequirements) validateRequests(formats strfmt.Registry) error {

	if swag.IsZero(m.Requests) { // not required
		return nil
	}

	for k := range m.Requests {

		if val, ok := m.Requests[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMCoreResourceRequirements) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMCoreResourceRequirements) UnmarshalBinary(b []byte) error {
	var res V1VMCoreResourceRequirements
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}