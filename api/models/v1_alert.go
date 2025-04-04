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
)

// V1Alert v1 alert
//
// swagger:model v1Alert
type V1Alert struct {

	// channels
	Channels []*V1Channel `json:"channels"`

	// component
	Component string `json:"component,omitempty"`
}

// Validate validates this v1 alert
func (m *V1Alert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Alert) validateChannels(formats strfmt.Registry) error {
	if swag.IsZero(m.Channels) { // not required
		return nil
	}

	for i := 0; i < len(m.Channels); i++ {
		if swag.IsZero(m.Channels[i]) { // not required
			continue
		}

		if m.Channels[i] != nil {
			if err := m.Channels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("channels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("channels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 alert based on the context it is used
func (m *V1Alert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChannels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Alert) contextValidateChannels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Channels); i++ {

		if m.Channels[i] != nil {

			if swag.IsZero(m.Channels[i]) { // not required
				return nil
			}

			if err := m.Channels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("channels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("channels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Alert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Alert) UnmarshalBinary(b []byte) error {
	var res V1Alert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
