// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Subnet v1 subnet
//
// swagger:model v1Subnet
type V1Subnet struct {

	// CidrBlock is the CIDR block to be used when the provider creates a managed Vnet.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Network Security Group(NSG) to be attached to subnet. NSG for a control plane subnet, should allow inbound to port 6443, as port 6443 is used by kubeadm to bootstrap the control planes
	SecurityGroupName string `json:"securityGroupName,omitempty"`
}

// Validate validates this v1 subnet
func (m *V1Subnet) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Subnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Subnet) UnmarshalBinary(b []byte) error {
	var res V1Subnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
