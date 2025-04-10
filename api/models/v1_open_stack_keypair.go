// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OpenStackKeypair OpenStack keypair. KeyPair is an SSH key known to the OpenStack Cloud that is available to be injected into servers
//
// swagger:model v1OpenStackKeypair
type V1OpenStackKeypair struct {

	// Name is used to refer to this keypair from other services within this region
	Name string `json:"name,omitempty"`

	// PublicKey is the public key from this pair, in OpenSSH format
	PublicKey string `json:"publicKey,omitempty"`
}

// Validate validates this v1 open stack keypair
func (m *V1OpenStackKeypair) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 open stack keypair based on context it is used
func (m *V1OpenStackKeypair) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1OpenStackKeypair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OpenStackKeypair) UnmarshalBinary(b []byte) error {
	var res V1OpenStackKeypair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
