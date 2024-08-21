// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CustomNic Custom network interface
//
// swagger:model v1CustomNic
type V1CustomNic struct {

	// index
	Index int8 `json:"index,omitempty"`

	// network name
	NetworkName string `json:"networkName,omitempty"`

	// private i ps
	PrivateIPs []string `json:"privateIPs"`

	// public Ip
	PublicIP string `json:"publicIp,omitempty"`
}

// Validate validates this v1 custom nic
func (m *V1CustomNic) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CustomNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CustomNic) UnmarshalBinary(b []byte) error {
	var res V1CustomNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}