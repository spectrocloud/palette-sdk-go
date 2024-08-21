// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EdgeHostSpecHost Host specifications
//
// swagger:model v1EdgeHostSpecHost
type V1EdgeHostSpecHost struct {

	// HostAddress is a FQDN or IP address of the Host
	HostAddress string `json:"hostAddress,omitempty"`

	// mac address
	MacAddress string `json:"macAddress,omitempty"`
}

// Validate validates this v1 edge host spec host
func (m *V1EdgeHostSpecHost) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeHostSpecHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeHostSpecHost) UnmarshalBinary(b []byte) error {
	var res V1EdgeHostSpecHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}