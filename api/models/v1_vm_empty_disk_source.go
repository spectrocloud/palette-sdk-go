// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMEmptyDiskSource EmptyDisk represents a temporary disk which shares the vmis lifecycle.
//
// swagger:model v1VmEmptyDiskSource
type V1VMEmptyDiskSource struct {

	// capacity
	// Required: true
	Capacity V1VMQuantity `json:"capacity"`
}

// Validate validates this v1 Vm empty disk source
func (m *V1VMEmptyDiskSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMEmptyDiskSource) validateCapacity(formats strfmt.Registry) error {

	if err := m.Capacity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("capacity")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMEmptyDiskSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMEmptyDiskSource) UnmarshalBinary(b []byte) error {
	var res V1VMEmptyDiskSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
