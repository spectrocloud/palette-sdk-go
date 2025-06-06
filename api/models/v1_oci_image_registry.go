// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OciImageRegistry Oci Image Registry
//
// swagger:model v1OciImageRegistry
type V1OciImageRegistry struct {

	// baseContentPath is the root path for the registry content
	BaseContentPath string `json:"baseContentPath,omitempty"`

	// ca cert
	CaCert string `json:"caCert,omitempty"`

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// insecure skip verify
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`

	// mirrorRegistries contains the array of image sources like gcr.io, ghcr.io, docker.io
	MirrorRegistries string `json:"mirrorRegistries,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this v1 oci image registry
func (m *V1OciImageRegistry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 oci image registry based on context it is used
func (m *V1OciImageRegistry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1OciImageRegistry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OciImageRegistry) UnmarshalBinary(b []byte) error {
	var res V1OciImageRegistry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
