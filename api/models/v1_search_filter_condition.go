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

// V1SearchFilterCondition v1 search filter condition
//
// swagger:model v1SearchFilterCondition
type V1SearchFilterCondition struct {

	// bool
	Bool *V1SearchFilterBoolCondition `json:"bool,omitempty"`

	// date
	Date *V1SearchFilterDateCondition `json:"date,omitempty"`

	// float
	Float *V1SearchFilterFloatCondition `json:"float,omitempty"`

	// int
	Int *V1SearchFilterIntegerCondition `json:"int,omitempty"`

	// key value
	KeyValue *V1SearchFilterKeyValueCondition `json:"keyValue,omitempty"`

	// string
	String *V1SearchFilterStringCondition `json:"string,omitempty"`
}

// Validate validates this v1 search filter condition
func (m *V1SearchFilterCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFloat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateString(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchFilterCondition) validateBool(formats strfmt.Registry) error {
	if swag.IsZero(m.Bool) { // not required
		return nil
	}

	if m.Bool != nil {
		if err := m.Bool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bool")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterCondition) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if m.Date != nil {
		if err := m.Date.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("date")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("date")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterCondition) validateFloat(formats strfmt.Registry) error {
	if swag.IsZero(m.Float) { // not required
		return nil
	}

	if m.Float != nil {
		if err := m.Float.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("float")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("float")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterCondition) validateInt(formats strfmt.Registry) error {
	if swag.IsZero(m.Int) { // not required
		return nil
	}

	if m.Int != nil {
		if err := m.Int.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("int")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("int")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterCondition) validateKeyValue(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyValue) { // not required
		return nil
	}

	if m.KeyValue != nil {
		if err := m.KeyValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyValue")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterCondition) validateString(formats strfmt.Registry) error {
	if swag.IsZero(m.String) { // not required
		return nil
	}

	if m.String != nil {
		if err := m.String.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("string")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("string")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 search filter condition based on the context it is used
func (m *V1SearchFilterCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFloat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateString(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchFilterCondition) contextValidateBool(ctx context.Context, formats strfmt.Registry) error {

	if m.Bool != nil {

		if swag.IsZero(m.Bool) { // not required
			return nil
		}

		if err := m.Bool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bool")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterCondition) contextValidateDate(ctx context.Context, formats strfmt.Registry) error {

	if m.Date != nil {

		if swag.IsZero(m.Date) { // not required
			return nil
		}

		if err := m.Date.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("date")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("date")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterCondition) contextValidateFloat(ctx context.Context, formats strfmt.Registry) error {

	if m.Float != nil {

		if swag.IsZero(m.Float) { // not required
			return nil
		}

		if err := m.Float.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("float")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("float")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterCondition) contextValidateInt(ctx context.Context, formats strfmt.Registry) error {

	if m.Int != nil {

		if swag.IsZero(m.Int) { // not required
			return nil
		}

		if err := m.Int.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("int")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("int")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterCondition) contextValidateKeyValue(ctx context.Context, formats strfmt.Registry) error {

	if m.KeyValue != nil {

		if swag.IsZero(m.KeyValue) { // not required
			return nil
		}

		if err := m.KeyValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyValue")
			}
			return err
		}
	}

	return nil
}

func (m *V1SearchFilterCondition) contextValidateString(ctx context.Context, formats strfmt.Registry) error {

	if m.String != nil {

		if swag.IsZero(m.String) { // not required
			return nil
		}

		if err := m.String.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("string")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("string")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SearchFilterCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchFilterCondition) UnmarshalBinary(b []byte) error {
	var res V1SearchFilterCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
