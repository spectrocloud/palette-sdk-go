// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1AppProfilesFilterSpec Application profile filter summary spec
//
// swagger:model v1AppProfilesFilterSpec
type V1AppProfilesFilterSpec struct {

	// filter
	Filter *V1AppProfileFilterSpec `json:"filter,omitempty"`

	// sort
	// Unique: true
	Sort []*V1AppProfileSortSpec `json:"sort"`
}

// Validate validates this v1 app profiles filter spec
func (m *V1AppProfilesFilterSpec) Validate(formats strfmt.Registry) error {
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

func (m *V1AppProfilesFilterSpec) validateFilter(formats strfmt.Registry) error {

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

func (m *V1AppProfilesFilterSpec) validateSort(formats strfmt.Registry) error {

	if swag.IsZero(m.Sort) { // not required
		return nil
	}

	if err := validate.UniqueItems("sort", "body", m.Sort); err != nil {
		return err
	}

	for i := 0; i < len(m.Sort); i++ {
		if swag.IsZero(m.Sort[i]) { // not required
			continue
		}

		if m.Sort[i] != nil {
			if err := m.Sort[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppProfilesFilterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppProfilesFilterSpec) UnmarshalBinary(b []byte) error {
	var res V1AppProfilesFilterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
