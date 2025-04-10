// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AwsImage AWS image name and ami
//
// swagger:model v1AwsImage
type V1AwsImage struct {

	// AWS image id
	ID string `json:"id,omitempty"`

	// AWS image name
	Name string `json:"name,omitempty"`

	// AWS image owner id
	Owner string `json:"owner,omitempty"`
}

// Validate validates this v1 aws image
func (m *V1AwsImage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 aws image based on context it is used
func (m *V1AwsImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsImage) UnmarshalBinary(b []byte) error {
	var res V1AwsImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
