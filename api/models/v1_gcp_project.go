// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GcpProject GCP project organizes all Google Cloud resources
//
// swagger:model v1GcpProject
type V1GcpProject struct {

	// GCP project id
	ID string `json:"id,omitempty"`

	// GCP project name
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 gcp project
func (m *V1GcpProject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 gcp project based on context it is used
func (m *V1GcpProject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GcpProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GcpProject) UnmarshalBinary(b []byte) error {
	var res V1GcpProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
