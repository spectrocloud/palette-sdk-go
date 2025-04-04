// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GcpZone A zone is a deployment area for Google Cloud resources within a region
//
// swagger:model v1GcpZone
type V1GcpZone struct {

	// GCP zone name
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 gcp zone
func (m *V1GcpZone) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 gcp zone based on context it is used
func (m *V1GcpZone) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GcpZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GcpZone) UnmarshalBinary(b []byte) error {
	var res V1GcpZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
