// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AwsCloudCostSummary AWS cloud account usage cost summary response data
//
// swagger:model v1AwsCloudCostSummary
type V1AwsCloudCostSummary struct {

	// cost
	Cost *V1AwsCloudCostSummaryCloudCost `json:"cost,omitempty"`
}

// Validate validates this v1 aws cloud cost summary
func (m *V1AwsCloudCostSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsCloudCostSummary) validateCost(formats strfmt.Registry) error {

	if swag.IsZero(m.Cost) { // not required
		return nil
	}

	if m.Cost != nil {
		if err := m.Cost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsCloudCostSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsCloudCostSummary) UnmarshalBinary(b []byte) error {
	var res V1AwsCloudCostSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
