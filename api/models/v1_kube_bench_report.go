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

// V1KubeBenchReport Compliance Scan KubeBench Report
//
// swagger:model v1KubeBenchReport
type V1KubeBenchReport struct {

	// fail
	Fail int32 `json:"fail,omitempty"`

	// info
	Info int32 `json:"info,omitempty"`

	// logs
	Logs []*V1KubeBenchLog `json:"logs"`

	// name
	Name string `json:"name,omitempty"`

	// pass
	Pass int32 `json:"pass,omitempty"`

	// time
	// Format: date-time
	Time V1Time `json:"time,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// warn
	Warn int32 `json:"warn,omitempty"`
}

// Validate validates this v1 kube bench report
func (m *V1KubeBenchReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1KubeBenchReport) validateLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.Logs) { // not required
		return nil
	}

	for i := 0; i < len(m.Logs); i++ {
		if swag.IsZero(m.Logs[i]) { // not required
			continue
		}

		if m.Logs[i] != nil {
			if err := m.Logs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1KubeBenchReport) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := m.Time.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("time")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 kube bench report based on the context it is used
func (m *V1KubeBenchReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1KubeBenchReport) contextValidateLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Logs); i++ {

		if m.Logs[i] != nil {

			if swag.IsZero(m.Logs[i]) { // not required
				return nil
			}

			if err := m.Logs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1KubeBenchReport) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := m.Time.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("time")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1KubeBenchReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1KubeBenchReport) UnmarshalBinary(b []byte) error {
	var res V1KubeBenchReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
