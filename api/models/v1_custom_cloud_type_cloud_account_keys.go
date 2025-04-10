// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CustomCloudTypeCloudAccountKeys Custom cloudType custom cloud account keys
//
// swagger:model v1CustomCloudTypeCloudAccountKeys
type V1CustomCloudTypeCloudAccountKeys struct {

	// Array of custom cloud type cloud account keys
	Keys []string `json:"keys"`
}

// Validate validates this v1 custom cloud type cloud account keys
func (m *V1CustomCloudTypeCloudAccountKeys) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 custom cloud type cloud account keys based on context it is used
func (m *V1CustomCloudTypeCloudAccountKeys) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CustomCloudTypeCloudAccountKeys) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CustomCloudTypeCloudAccountKeys) UnmarshalBinary(b []byte) error {
	var res V1CustomCloudTypeCloudAccountKeys
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
