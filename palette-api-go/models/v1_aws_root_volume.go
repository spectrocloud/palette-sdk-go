// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AwsRootVolume Volume encapsulates the configuration options for the storage device.
//
// swagger:model v1AwsRootVolume
type V1AwsRootVolume struct {

	// Device name
	DeviceName string `json:"deviceName,omitempty"`

	// EncryptionKey is the KMS key to use to encrypt the volume. Can be either a KMS key ID or ARN
	Encrypted bool `json:"encrypted,omitempty"`

	// EncryptionKey is the KMS key to use to encrypt the volume. Can be either a KMS key ID or ARN
	EncryptionKey string `json:"encryptionKey,omitempty"`

	// IOPS is the number of IOPS requested for the disk. Not applicable to all types
	Iops int64 `json:"iops,omitempty"`

	// Throughput to provision in MiB/s supported for the volume type. Not applicable to all types.
	Throughput int64 `json:"throughput,omitempty"`

	// Type is the type of the volume (e.g. gp2, io1, etc...)
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 aws root volume
func (m *V1AwsRootVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsRootVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsRootVolume) UnmarshalBinary(b []byte) error {
	var res V1AwsRootVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
