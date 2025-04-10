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
	"github.com/go-openapi/validate"
)

// V1AwsRegions v1 aws regions
//
// swagger:model v1AwsRegions
type V1AwsRegions struct {

	// List of AWS regions
	// Required: true
	Regions []*V1AwsRegion `json:"regions"`
}

// Validate validates this v1 aws regions
func (m *V1AwsRegions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsRegions) validateRegions(formats strfmt.Registry) error {

	if err := validate.Required("regions", "body", m.Regions); err != nil {
		return err
	}

	for i := 0; i < len(m.Regions); i++ {
		if swag.IsZero(m.Regions[i]) { // not required
			continue
		}

		if m.Regions[i] != nil {
			if err := m.Regions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 aws regions based on the context it is used
func (m *V1AwsRegions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsRegions) contextValidateRegions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Regions); i++ {

		if m.Regions[i] != nil {

			if swag.IsZero(m.Regions[i]) { // not required
				return nil
			}

			if err := m.Regions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsRegions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsRegions) UnmarshalBinary(b []byte) error {
	var res V1AwsRegions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
