// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Nic v1 nic
//
// swagger:model v1Nic
type V1Nic struct {

	// dns
	DNS []string `json:"dns"`

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// ip
	IP string `json:"ip,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// mac addr
	MacAddr string `json:"macAddr,omitempty"`

	// nic name
	NicName string `json:"nicName,omitempty"`

	// subnet
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this v1 nic
func (m *V1Nic) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 nic based on context it is used
func (m *V1Nic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Nic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Nic) UnmarshalBinary(b []byte) error {
	var res V1Nic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
