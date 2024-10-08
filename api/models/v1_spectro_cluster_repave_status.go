// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroClusterRepaveStatus v1 spectro cluster repave status
//
// swagger:model v1SpectroClusterRepaveStatus
type V1SpectroClusterRepaveStatus struct {

	// message
	Message string `json:"message,omitempty"`

	// repave transition time
	// Format: date-time
	RepaveTransitionTime V1Time `json:"repaveTransitionTime,omitempty"`

	// state
	State V1ClusterRepaveState `json:"state,omitempty"`
}

// Validate validates this v1 spectro cluster repave status
func (m *V1SpectroClusterRepaveStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepaveTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterRepaveStatus) validateRepaveTransitionTime(formats strfmt.Registry) error {

	if swag.IsZero(m.RepaveTransitionTime) { // not required
		return nil
	}

	if err := m.RepaveTransitionTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("repaveTransitionTime")
		}
		return err
	}

	return nil
}

func (m *V1SpectroClusterRepaveStatus) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterRepaveStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterRepaveStatus) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterRepaveStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
