// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMDataVolumeSourceRegistry DataVolumeSourceRegistry provides the parameters to create a Data Volume from an registry source
//
// swagger:model v1VmDataVolumeSourceRegistry
type V1VMDataVolumeSourceRegistry struct {

	// CertConfigMap provides a reference to the Registry certs
	CertConfigMap string `json:"certConfigMap,omitempty"`

	// ImageStream is the name of image stream for import
	ImageStream string `json:"imageStream,omitempty"`

	// PullMethod can be either "pod" (default import), or "node" (node docker cache based import)
	PullMethod string `json:"pullMethod,omitempty"`

	// SecretRef provides the secret reference needed to access the Registry source
	SecretRef string `json:"secretRef,omitempty"`

	// URL is the url of the registry source (starting with the scheme: docker, oci-archive)
	URL string `json:"url,omitempty"`
}

// Validate validates this v1 Vm data volume source registry
func (m *V1VMDataVolumeSourceRegistry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMDataVolumeSourceRegistry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMDataVolumeSourceRegistry) UnmarshalBinary(b []byte) error {
	var res V1VMDataVolumeSourceRegistry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
