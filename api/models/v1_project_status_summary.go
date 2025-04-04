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

// V1ProjectStatusSummary Project status summary
//
// swagger:model v1ProjectStatusSummary
type V1ProjectStatusSummary struct {

	// clusters health
	ClustersHealth *V1SpectroClustersHealth `json:"clustersHealth,omitempty"`

	// status
	Status *V1ProjectStatus `json:"status,omitempty"`

	// usage
	Usage *V1ProjectUsageSummary `json:"usage,omitempty"`
}

// Validate validates this v1 project status summary
func (m *V1ProjectStatusSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClustersHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectStatusSummary) validateClustersHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.ClustersHealth) { // not required
		return nil
	}

	if m.ClustersHealth != nil {
		if err := m.ClustersHealth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clustersHealth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clustersHealth")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectStatusSummary) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectStatusSummary) validateUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.Usage) { // not required
		return nil
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 project status summary based on the context it is used
func (m *V1ProjectStatusSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClustersHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectStatusSummary) contextValidateClustersHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.ClustersHealth != nil {

		if swag.IsZero(m.ClustersHealth) { // not required
			return nil
		}

		if err := m.ClustersHealth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clustersHealth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clustersHealth")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectStatusSummary) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectStatusSummary) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.Usage != nil {

		if swag.IsZero(m.Usage) { // not required
			return nil
		}

		if err := m.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectStatusSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectStatusSummary) UnmarshalBinary(b []byte) error {
	var res V1ProjectStatusSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
