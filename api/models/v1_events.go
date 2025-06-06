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

// V1Events An array of component events items
//
// swagger:model v1Events
type V1Events struct {

	// Describes a list of returned component events
	// Required: true
	// Unique: true
	Items []*V1Event `json:"items"`

	// Describes the meta information about the component event lists
	Listmeta *V1ListMetaData `json:"listmeta,omitempty"`
}

// Validate validates this v1 events
func (m *V1Events) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListmeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Events) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	if err := validate.UniqueItems("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Events) validateListmeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Listmeta) { // not required
		return nil
	}

	if m.Listmeta != nil {
		if err := m.Listmeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listmeta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("listmeta")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 events based on the context it is used
func (m *V1Events) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateListmeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Events) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {

			if swag.IsZero(m.Items[i]) { // not required
				return nil
			}

			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Events) contextValidateListmeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Listmeta != nil {

		if swag.IsZero(m.Listmeta) { // not required
			return nil
		}

		if err := m.Listmeta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listmeta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("listmeta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Events) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Events) UnmarshalBinary(b []byte) error {
	var res V1Events
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
