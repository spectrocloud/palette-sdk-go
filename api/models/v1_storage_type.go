// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1StorageType Cloud cloud Storage type details
//
// swagger:model v1StorageType
type V1StorageType struct {

	// cost
	Cost *V1StorageCost `json:"cost,omitempty"`

	// iops cost
	IopsCost *V1StorageCost `json:"iopsCost,omitempty"`

	// kind of storage type
	Kind string `json:"kind,omitempty"`

	// Name of the storage type
	Name string `json:"name,omitempty"`

	// throughput cost
	ThroughputCost *V1StorageCost `json:"throughputCost,omitempty"`
}

// Validate validates this v1 storage type
func (m *V1StorageType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughputCost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1StorageType) validateCost(formats strfmt.Registry) error {

	if swag.IsZero(m.Cost) { // not required
		return nil
	}

	if m.Cost != nil {
		if err := m.Cost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost")
			}
			return err
		}
	}

	return nil
}

func (m *V1StorageType) validateIopsCost(formats strfmt.Registry) error {

	if swag.IsZero(m.IopsCost) { // not required
		return nil
	}

	if m.IopsCost != nil {
		if err := m.IopsCost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iopsCost")
			}
			return err
		}
	}

	return nil
}

func (m *V1StorageType) validateThroughputCost(formats strfmt.Registry) error {

	if swag.IsZero(m.ThroughputCost) { // not required
		return nil
	}

	if m.ThroughputCost != nil {
		if err := m.ThroughputCost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughputCost")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1StorageType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1StorageType) UnmarshalBinary(b []byte) error {
	var res V1StorageType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
