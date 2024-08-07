// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EdgeHostDeviceMetaUpdateEntity Edge host device uid and name
//
// swagger:model v1EdgeHostDeviceMetaUpdateEntity
type V1EdgeHostDeviceMetaUpdateEntity struct {

	// metadata
	Metadata *V1ObjectTagsEntity `json:"metadata,omitempty"`
}

// Validate validates this v1 edge host device meta update entity
func (m *V1EdgeHostDeviceMetaUpdateEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeHostDeviceMetaUpdateEntity) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeHostDeviceMetaUpdateEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeHostDeviceMetaUpdateEntity) UnmarshalBinary(b []byte) error {
	var res V1EdgeHostDeviceMetaUpdateEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
