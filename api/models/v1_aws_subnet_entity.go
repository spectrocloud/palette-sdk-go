// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AwsSubnetEntity v1 aws subnet entity
//
// swagger:model v1AwsSubnetEntity
type V1AwsSubnetEntity struct {

	// az
	Az string `json:"az,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this v1 aws subnet entity
func (m *V1AwsSubnetEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsSubnetEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsSubnetEntity) UnmarshalBinary(b []byte) error {
	var res V1AwsSubnetEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
