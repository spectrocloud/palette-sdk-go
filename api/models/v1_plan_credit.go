// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1PlanCredit Plan Credit
//
// swagger:model v1PlanCredit
type V1PlanCredit struct {

	// cpu core hours
	CPUCoreHours int64 `json:"cpuCoreHours"`

	// credit Uid
	CreditUID string `json:"creditUid,omitempty"`

	// credit expiry time
	// Format: date-time
	Expiry V1Time `json:"expiry,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// credit start time
	// Format: date-time
	Start V1Time `json:"start,omitempty"`

	// type
	// Required: true
	// Enum: ["Pure","Alloy"]
	Type *string `json:"type"`
}

// Validate validates this v1 plan credit
func (m *V1PlanCredit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PlanCredit) validateExpiry(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiry) { // not required
		return nil
	}

	if err := m.Expiry.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("expiry")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("expiry")
		}
		return err
	}

	return nil
}

func (m *V1PlanCredit) validateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := m.Start.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("start")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("start")
		}
		return err
	}

	return nil
}

var v1PlanCreditTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pure","Alloy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1PlanCreditTypeTypePropEnum = append(v1PlanCreditTypeTypePropEnum, v)
	}
}

const (

	// V1PlanCreditTypePure captures enum value "Pure"
	V1PlanCreditTypePure string = "Pure"

	// V1PlanCreditTypeAlloy captures enum value "Alloy"
	V1PlanCreditTypeAlloy string = "Alloy"
)

// prop value enum
func (m *V1PlanCredit) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1PlanCreditTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1PlanCredit) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 plan credit based on the context it is used
func (m *V1PlanCredit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExpiry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PlanCredit) contextValidateExpiry(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Expiry) { // not required
		return nil
	}

	if err := m.Expiry.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("expiry")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("expiry")
		}
		return err
	}

	return nil
}

func (m *V1PlanCredit) contextValidateStart(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := m.Start.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("start")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("start")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PlanCredit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PlanCredit) UnmarshalBinary(b []byte) error {
	var res V1PlanCredit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
