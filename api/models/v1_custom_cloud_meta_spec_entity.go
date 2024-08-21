// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CustomCloudMetaSpecEntity Custom cloud spec response entity
//
// swagger:model v1CustomCloudMetaSpecEntity
type V1CustomCloudMetaSpecEntity struct {

	// cloud category
	CloudCategory V1CloudCategory `json:"cloudCategory,omitempty"`

	// Custom cloud displayName
	DisplayName string `json:"displayName,omitempty"`

	// If the custom cloud is a managed cluster
	IsManaged bool `json:"isManaged,omitempty"`

	// Custom cloud logo
	Logo string `json:"logo,omitempty"`
}

// Validate validates this v1 custom cloud meta spec entity
func (m *V1CustomCloudMetaSpecEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CustomCloudMetaSpecEntity) validateCloudCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudCategory) { // not required
		return nil
	}

	if err := m.CloudCategory.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cloudCategory")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CustomCloudMetaSpecEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CustomCloudMetaSpecEntity) UnmarshalBinary(b []byte) error {
	var res V1CustomCloudMetaSpecEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}