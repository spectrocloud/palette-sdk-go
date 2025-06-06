// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AwsVolumeSize AWS Volume Size entity
//
// swagger:model v1AwsVolumeSize
type V1AwsVolumeSize struct {

	// AWS volume size
	SizeGB int64 `json:"sizeGB,omitempty"`
}

// Validate validates this v1 aws volume size
func (m *V1AwsVolumeSize) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 aws volume size based on context it is used
func (m *V1AwsVolumeSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsVolumeSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsVolumeSize) UnmarshalBinary(b []byte) error {
	var res V1AwsVolumeSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
