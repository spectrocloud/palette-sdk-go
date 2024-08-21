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

// V1EdgeMachinePoolHostEntity v1 edge machine pool host entity
//
// swagger:model v1EdgeMachinePoolHostEntity
type V1EdgeMachinePoolHostEntity struct {

	// host Uid
	// Required: true
	HostUID *string `json:"hostUid"`
}

// Validate validates this v1 edge machine pool host entity
func (m *V1EdgeMachinePoolHostEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeMachinePoolHostEntity) validateHostUID(formats strfmt.Registry) error {

	if err := validate.Required("hostUid", "body", m.HostUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeMachinePoolHostEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeMachinePoolHostEntity) UnmarshalBinary(b []byte) error {
	var res V1EdgeMachinePoolHostEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}