// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AwsKeyPairs List of AWS keypairs
//
// swagger:model v1AwsKeyPairs
type V1AwsKeyPairs struct {

	// Array of Aws Keypair names
	KeyNames []string `json:"keyNames"`
}

// Validate validates this v1 aws key pairs
func (m *V1AwsKeyPairs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 aws key pairs based on context it is used
func (m *V1AwsKeyPairs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsKeyPairs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsKeyPairs) UnmarshalBinary(b []byte) error {
	var res V1AwsKeyPairs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
