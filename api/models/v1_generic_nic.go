// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GenericNic Generic network interface
//
// swagger:model v1GenericNic
type V1GenericNic struct {

	// index
	Index int8 `json:"index,omitempty"`

	// network name
	NetworkName string `json:"networkName,omitempty"`

	// private i ps
	PrivateIPs []string `json:"privateIPs"`

	// public Ip
	PublicIP string `json:"publicIp,omitempty"`
}

// Validate validates this v1 generic nic
func (m *V1GenericNic) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 generic nic based on context it is used
func (m *V1GenericNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GenericNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GenericNic) UnmarshalBinary(b []byte) error {
	var res V1GenericNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
