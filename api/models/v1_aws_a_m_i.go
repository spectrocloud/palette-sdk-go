// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AwsAMI v1 aws a m i
//
// swagger:model v1AwsAMI
type V1AwsAMI struct {

	// id
	ID string `json:"id,omitempty"`

	// os
	Os string `json:"os,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 aws a m i
func (m *V1AwsAMI) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 aws a m i based on context it is used
func (m *V1AwsAMI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsAMI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsAMI) UnmarshalBinary(b []byte) error {
	var res V1AwsAMI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
