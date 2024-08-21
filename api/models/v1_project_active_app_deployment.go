// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ProjectActiveAppDeployment Active app deployment
//
// swagger:model v1ProjectActiveAppDeployment
type V1ProjectActiveAppDeployment struct {

	// app ref
	AppRef *V1ObjectEntity `json:"appRef,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 project active app deployment
func (m *V1ProjectActiveAppDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectActiveAppDeployment) validateAppRef(formats strfmt.Registry) error {

	if swag.IsZero(m.AppRef) { // not required
		return nil
	}

	if m.AppRef != nil {
		if err := m.AppRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectActiveAppDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectActiveAppDeployment) UnmarshalBinary(b []byte) error {
	var res V1ProjectActiveAppDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}