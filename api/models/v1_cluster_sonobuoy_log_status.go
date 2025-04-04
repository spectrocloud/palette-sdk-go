// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ClusterSonobuoyLogStatus Cluster compliance scan Sonobuoy Log Status
//
// swagger:model v1ClusterSonobuoyLogStatus
type V1ClusterSonobuoyLogStatus struct {

	// actor
	Actor *V1ClusterFeatureActor `json:"actor,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// reports
	Reports map[string]V1SonobuoyReport `json:"reports,omitempty"`

	// request Uid
	RequestUID string `json:"requestUid,omitempty"`

	// scan time
	ScanTime *V1ClusterScanTime `json:"scanTime,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 cluster sonobuoy log status
func (m *V1ClusterSonobuoyLogStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterSonobuoyLogStatus) validateActor(formats strfmt.Registry) error {
	if swag.IsZero(m.Actor) { // not required
		return nil
	}

	if m.Actor != nil {
		if err := m.Actor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actor")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterSonobuoyLogStatus) validateReports(formats strfmt.Registry) error {
	if swag.IsZero(m.Reports) { // not required
		return nil
	}

	for k := range m.Reports {

		if err := validate.Required("reports"+"."+k, "body", m.Reports[k]); err != nil {
			return err
		}
		if val, ok := m.Reports[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reports" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reports" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterSonobuoyLogStatus) validateScanTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ScanTime) { // not required
		return nil
	}

	if m.ScanTime != nil {
		if err := m.ScanTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scanTime")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 cluster sonobuoy log status based on the context it is used
func (m *V1ClusterSonobuoyLogStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScanTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterSonobuoyLogStatus) contextValidateActor(ctx context.Context, formats strfmt.Registry) error {

	if m.Actor != nil {

		if swag.IsZero(m.Actor) { // not required
			return nil
		}

		if err := m.Actor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actor")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterSonobuoyLogStatus) contextValidateReports(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Reports {

		if val, ok := m.Reports[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V1ClusterSonobuoyLogStatus) contextValidateScanTime(ctx context.Context, formats strfmt.Registry) error {

	if m.ScanTime != nil {

		if swag.IsZero(m.ScanTime) { // not required
			return nil
		}

		if err := m.ScanTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scanTime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterSonobuoyLogStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterSonobuoyLogStatus) UnmarshalBinary(b []byte) error {
	var res V1ClusterSonobuoyLogStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
