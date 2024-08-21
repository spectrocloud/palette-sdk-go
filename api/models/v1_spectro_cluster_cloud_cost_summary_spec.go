// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroClusterCloudCostSummarySpec Spectro cluster cloud cost summary spec
//
// swagger:model v1SpectroClusterCloudCostSummarySpec
type V1SpectroClusterCloudCostSummarySpec struct {

	// filter
	Filter *V1SpectroClusterCloudCostSummaryFilter `json:"filter,omitempty"`

	// options
	Options *V1SpectroClusterCloudCostSummaryOptions `json:"options,omitempty"`
}

// Validate validates this v1 spectro cluster cloud cost summary spec
func (m *V1SpectroClusterCloudCostSummarySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterCloudCostSummarySpec) validateFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterCloudCostSummarySpec) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {
		if err := m.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterCloudCostSummarySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterCloudCostSummarySpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterCloudCostSummarySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}