// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ZoneEntity Azure availability zone entity
//
// swagger:model v1ZoneEntity
type V1ZoneEntity struct {

	// Azure availability zone id
	ID string `json:"id,omitempty"`
}

// Validate validates this v1 zone entity
func (m *V1ZoneEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ZoneEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ZoneEntity) UnmarshalBinary(b []byte) error {
	var res V1ZoneEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}