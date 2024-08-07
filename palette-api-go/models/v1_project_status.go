// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ProjectStatus Project status
//
// swagger:model v1ProjectStatus
type V1ProjectStatus struct {

	// clean up status
	CleanUpStatus *V1ProjectCleanUpStatus `json:"cleanUpStatus,omitempty"`

	// is disabled
	IsDisabled bool `json:"isDisabled,omitempty"`
}

// Validate validates this v1 project status
func (m *V1ProjectStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCleanUpStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectStatus) validateCleanUpStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.CleanUpStatus) { // not required
		return nil
	}

	if m.CleanUpStatus != nil {
		if err := m.CleanUpStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cleanUpStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectStatus) UnmarshalBinary(b []byte) error {
	var res V1ProjectStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
