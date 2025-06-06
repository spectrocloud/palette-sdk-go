// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMClock Represents the clock and timers of a vmi.
//
// swagger:model v1VmClock
type V1VMClock struct {

	// timer
	Timer *V1VMTimer `json:"timer,omitempty"`

	// Timezone sets the guest clock to the specified timezone. Zone name follows the TZ environment variable format (e.g. 'America/New_York').
	Timezone string `json:"timezone,omitempty"`

	// utc
	Utc *V1VMClockOffsetUTC `json:"utc,omitempty"`
}

// Validate validates this v1 Vm clock
func (m *V1VMClock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUtc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMClock) validateTimer(formats strfmt.Registry) error {
	if swag.IsZero(m.Timer) { // not required
		return nil
	}

	if m.Timer != nil {
		if err := m.Timer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timer")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMClock) validateUtc(formats strfmt.Registry) error {
	if swag.IsZero(m.Utc) { // not required
		return nil
	}

	if m.Utc != nil {
		if err := m.Utc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("utc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("utc")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 Vm clock based on the context it is used
func (m *V1VMClock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTimer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUtc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMClock) contextValidateTimer(ctx context.Context, formats strfmt.Registry) error {

	if m.Timer != nil {

		if swag.IsZero(m.Timer) { // not required
			return nil
		}

		if err := m.Timer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timer")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMClock) contextValidateUtc(ctx context.Context, formats strfmt.Registry) error {

	if m.Utc != nil {

		if swag.IsZero(m.Utc) { // not required
			return nil
		}

		if err := m.Utc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("utc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("utc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMClock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMClock) UnmarshalBinary(b []byte) error {
	var res V1VMClock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
