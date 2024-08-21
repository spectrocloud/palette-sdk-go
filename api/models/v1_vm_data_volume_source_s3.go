// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1VMDataVolumeSourceS3 DataVolumeSourceS3 provides the parameters to create a Data Volume from an S3 source
//
// swagger:model v1VmDataVolumeSourceS3
type V1VMDataVolumeSourceS3 struct {

	// CertConfigMap is a configmap reference, containing a Certificate Authority(CA) public key, and a base64 encoded pem certificate
	CertConfigMap string `json:"certConfigMap,omitempty"`

	// SecretRef provides the secret reference needed to access the S3 source
	SecretRef string `json:"secretRef,omitempty"`

	// URL is the url of the S3 source
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this v1 Vm data volume source s3
func (m *V1VMDataVolumeSourceS3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMDataVolumeSourceS3) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMDataVolumeSourceS3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMDataVolumeSourceS3) UnmarshalBinary(b []byte) error {
	var res V1VMDataVolumeSourceS3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}