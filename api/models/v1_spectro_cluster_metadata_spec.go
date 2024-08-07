// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SpectroClusterMetadataSpec Spectro cluster metadata spec
//
// swagger:model v1SpectroClusterMetadataSpec
type V1SpectroClusterMetadataSpec struct {

	// filter
	Filter *V1SpectroClusterMetadataFilterSpec `json:"filter,omitempty"`

	// sort
	// Enum: [environment state name]
	Sort *string `json:"sort,omitempty"`
}

// Validate validates this v1 spectro cluster metadata spec
func (m *V1SpectroClusterMetadataSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterMetadataSpec) validateFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

var v1SpectroClusterMetadataSpecTypeSortPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["environment","state","name"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SpectroClusterMetadataSpecTypeSortPropEnum = append(v1SpectroClusterMetadataSpecTypeSortPropEnum, v)
	}
}

const (

	// V1SpectroClusterMetadataSpecSortEnvironment captures enum value "environment"
	V1SpectroClusterMetadataSpecSortEnvironment string = "environment"

	// V1SpectroClusterMetadataSpecSortState captures enum value "state"
	V1SpectroClusterMetadataSpecSortState string = "state"

	// V1SpectroClusterMetadataSpecSortName captures enum value "name"
	V1SpectroClusterMetadataSpecSortName string = "name"
)

// prop value enum
func (m *V1SpectroClusterMetadataSpec) validateSortEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1SpectroClusterMetadataSpecTypeSortPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1SpectroClusterMetadataSpec) validateSort(formats strfmt.Registry) error {

	if swag.IsZero(m.Sort) { // not required
		return nil
	}

	// value enum
	if err := m.validateSortEnum("sort", "body", *m.Sort); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterMetadataSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterMetadataSpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterMetadataSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
