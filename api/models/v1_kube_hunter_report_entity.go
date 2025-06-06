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

// V1KubeHunterReportEntity KubeHunter report
//
// swagger:model v1KubeHunterReportEntity
type V1KubeHunterReportEntity struct {

	// logs
	Logs []*V1KubeHunterLogEntity `json:"logs"`

	// time
	// Format: date-time
	Time V1Time `json:"time,omitempty"`

	// vulnerabilities
	Vulnerabilities *V1KubeHunterVulnerabilityDataEntity `json:"vulnerabilities,omitempty"`
}

// Validate validates this v1 kube hunter report entity
func (m *V1KubeHunterReportEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1KubeHunterReportEntity) validateLogs(formats strfmt.Registry) error {
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

func (m *V1KubeHunterReportEntity) validateTime(formats strfmt.Registry) error {
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

func (m *V1KubeHunterReportEntity) validateVulnerabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Vulnerabilities) { // not required
		return nil
	}

	if m.Vulnerabilities != nil {
		if err := m.Vulnerabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vulnerabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vulnerabilities")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 kube hunter report entity based on the context it is used
func (m *V1KubeHunterReportEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVulnerabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1KubeHunterReportEntity) contextValidateLogs(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1KubeHunterReportEntity) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1KubeHunterReportEntity) contextValidateVulnerabilities(ctx context.Context, formats strfmt.Registry) error {

	if m.Vulnerabilities != nil {

		if swag.IsZero(m.Vulnerabilities) { // not required
			return nil
		}

		if err := m.Vulnerabilities.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vulnerabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vulnerabilities")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1KubeHunterReportEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1KubeHunterReportEntity) UnmarshalBinary(b []byte) error {
	var res V1KubeHunterReportEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
