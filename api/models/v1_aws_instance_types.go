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

// V1AwsInstanceTypes List of AWS instance types
//
// swagger:model v1AwsInstanceTypes
type V1AwsInstanceTypes struct {

	// instance types
	InstanceTypes []*V1InstanceType `json:"instanceTypes"`
}

// Validate validates this v1 aws instance types
func (m *V1AwsInstanceTypes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsInstanceTypes) validateInstanceTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.InstanceTypes); i++ {
		if swag.IsZero(m.InstanceTypes[i]) { // not required
			continue
		}

		if m.InstanceTypes[i] != nil {
			if err := m.InstanceTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsInstanceTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsInstanceTypes) UnmarshalBinary(b []byte) error {
	var res V1AwsInstanceTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
