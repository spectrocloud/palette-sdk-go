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

// V1InstanceCost Instance cost entity
//
// swagger:model v1InstanceCost
type V1InstanceCost struct {

	// Array of cloud instance price
	Price []*V1InstancePrice `json:"price"`
}

// Validate validates this v1 instance cost
func (m *V1InstanceCost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1InstanceCost) validatePrice(formats strfmt.Registry) error {

	if swag.IsZero(m.Price) { // not required
		return nil
	}

	for i := 0; i < len(m.Price); i++ {
		if swag.IsZero(m.Price[i]) { // not required
			continue
		}

		if m.Price[i] != nil {
			if err := m.Price[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("price" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1InstanceCost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1InstanceCost) UnmarshalBinary(b []byte) error {
	var res V1InstanceCost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}