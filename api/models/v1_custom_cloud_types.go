// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CustomCloudTypes Custom cloudType content response
//
// swagger:model v1CustomCloudTypes
type V1CustomCloudTypes struct {

	// Array of custom cloud types
	CloudTypes []*V1CustomCloudType `json:"cloudTypes"`
}

// Validate validates this v1 custom cloud types
func (m *V1CustomCloudTypes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CustomCloudTypes) validateCloudTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.CloudTypes); i++ {
		if swag.IsZero(m.CloudTypes[i]) { // not required
			continue
		}

		if m.CloudTypes[i] != nil {
			if err := m.CloudTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CustomCloudTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CustomCloudTypes) UnmarshalBinary(b []byte) error {
	var res V1CustomCloudTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}