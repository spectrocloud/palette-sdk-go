// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1EdgeHostsMetadataFilter Edge host metadata spec
//
// swagger:model v1EdgeHostsMetadataFilter
type V1EdgeHostsMetadataFilter struct {

	// filter
	Filter *V1EdgeHostsMetadataFilterSpec `json:"filter,omitempty"`

	// sort
	// Unique: true
	Sort []*V1EdgeHostsMetadataSortSpec `json:"sort"`
}

// Validate validates this v1 edge hosts metadata filter
func (m *V1EdgeHostsMetadataFilter) Validate(formats strfmt.Registry) error {
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

func (m *V1EdgeHostsMetadataFilter) validateFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *V1EdgeHostsMetadataFilter) validateSort(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 edge hosts metadata filter based on the context it is used
func (m *V1EdgeHostsMetadataFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeHostsMetadataFilter) contextValidateFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.Filter != nil {

		if swag.IsZero(m.Filter) { // not required
			return nil
		}

		if err := m.Filter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *V1EdgeHostsMetadataFilter) contextValidateSort(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sort); i++ {

		if m.Sort[i] != nil {

			if swag.IsZero(m.Sort[i]) { // not required
				return nil
			}

			if err := m.Sort[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeHostsMetadataFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeHostsMetadataFilter) UnmarshalBinary(b []byte) error {
	var res V1EdgeHostsMetadataFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
