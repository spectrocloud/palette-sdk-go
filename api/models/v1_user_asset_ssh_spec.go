// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UserAssetSSHSpec SSH key specification
//
// swagger:model v1UserAssetSshSpec
type V1UserAssetSSHSpec struct {

	// public key
	PublicKey string `json:"publicKey,omitempty"`
}

// Validate validates this v1 user asset Ssh spec
func (m *V1UserAssetSSHSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 user asset Ssh spec based on context it is used
func (m *V1UserAssetSSHSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UserAssetSSHSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UserAssetSSHSpec) UnmarshalBinary(b []byte) error {
	var res V1UserAssetSSHSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
