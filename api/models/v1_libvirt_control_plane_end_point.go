// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1LibvirtControlPlaneEndPoint v1 libvirt control plane end point
//
// swagger:model v1LibvirtControlPlaneEndPoint
type V1LibvirtControlPlaneEndPoint struct {

	// DDNSSearchDomain is the search domain used for resolving IP addresses when the EndpointType is DDNS. This search domain is appended to the generated Hostname to obtain the complete DNS name for the endpoint. If Host is already a DDNS FQDN, DDNSSearchDomain is not required
	DdnsSearchDomain string `json:"ddnsSearchDomain,omitempty"`

	// Host is FQDN(DDNS) or IP
	Host string `json:"host,omitempty"`

	// Type indicates DDNS or VIP
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 libvirt control plane end point
func (m *V1LibvirtControlPlaneEndPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1LibvirtControlPlaneEndPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LibvirtControlPlaneEndPoint) UnmarshalBinary(b []byte) error {
	var res V1LibvirtControlPlaneEndPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}