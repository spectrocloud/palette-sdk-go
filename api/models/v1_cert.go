// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Cert v1 cert
//
// swagger:model v1Cert
type V1Cert struct {

	// certificate
	Certificate string `json:"certificate"`

	// is c a
	IsCA bool `json:"isCA"`

	// key
	Key string `json:"key"`
}

// Validate validates this v1 cert
func (m *V1Cert) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 cert based on context it is used
func (m *V1Cert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Cert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Cert) UnmarshalBinary(b []byte) error {
	var res V1Cert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
