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

// V1CustomCloudType v1 custom cloud type
//
// swagger:model v1CustomCloudType
type V1CustomCloudType struct {

	// cloud category
	CloudCategory *V1CloudCategory `json:"cloudCategory,omitempty"`

	// Cloud grouping as family
	CloudFamily string `json:"cloudFamily,omitempty"`

	// Custom cloudtype displayName
	DisplayName string `json:"displayName,omitempty"`

	// If it is a custom cloudtype
	IsCustom bool `json:"isCustom"`

	// If custom cloudtype is managed
	IsManaged bool `json:"isManaged"`

	// If cloud is support for Vertex env
	IsVertex bool `json:"isVertex"`

	// Custom cloudtype logo
	Logo string `json:"logo,omitempty"`

	// Custom cloudtype name
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 custom cloud type
func (m *V1CustomCloudType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CustomCloudType) validateCloudCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudCategory) { // not required
		return nil
	}

	if m.CloudCategory != nil {
		if err := m.CloudCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudCategory")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 custom cloud type based on the context it is used
func (m *V1CustomCloudType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CustomCloudType) contextValidateCloudCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudCategory != nil {

		if swag.IsZero(m.CloudCategory) { // not required
			return nil
		}

		if err := m.CloudCategory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudCategory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CustomCloudType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CustomCloudType) UnmarshalBinary(b []byte) error {
	var res V1CustomCloudType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
