// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMPodNetwork Represents the stock pod network interface.
//
// swagger:model v1VmPodNetwork
type V1VMPodNetwork struct {

	// IPv6 CIDR for the vm network. Defaults to fd10:0:2::/120 if not specified.
	VMIPV6NetworkCIDR string `json:"vmIPv6NetworkCIDR,omitempty"`

	// CIDR for vm network. Default 10.0.2.0/24 if not specified.
	VMNetworkCIDR string `json:"vmNetworkCIDR,omitempty"`
}

// Validate validates this v1 Vm pod network
func (m *V1VMPodNetwork) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 Vm pod network based on context it is used
func (m *V1VMPodNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMPodNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMPodNetwork) UnmarshalBinary(b []byte) error {
	var res V1VMPodNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
