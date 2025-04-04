// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1FeatureUpdateSpec Feature update spec
//
// swagger:model v1FeatureUpdateSpec
type V1FeatureUpdateSpec struct {

	// Feature description
	Description string `json:"description,omitempty"`

	// Feature doc link
	DocLink string `json:"docLink,omitempty"`
}

// Validate validates this v1 feature update spec
func (m *V1FeatureUpdateSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 feature update spec based on context it is used
func (m *V1FeatureUpdateSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1FeatureUpdateSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FeatureUpdateSpec) UnmarshalBinary(b []byte) error {
	var res V1FeatureUpdateSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
