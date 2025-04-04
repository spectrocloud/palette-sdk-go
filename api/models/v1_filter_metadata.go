// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1FilterMetadata Filter metadata object
//
// swagger:model v1FilterMetadata
type V1FilterMetadata struct {

	// filter type
	FilterType string `json:"filterType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 filter metadata
func (m *V1FilterMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 filter metadata based on context it is used
func (m *V1FilterMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1FilterMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FilterMetadata) UnmarshalBinary(b []byte) error {
	var res V1FilterMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
