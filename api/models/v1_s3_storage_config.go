// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1S3StorageConfig S3 storage config object
//
// swagger:model v1S3StorageConfig
type V1S3StorageConfig struct {

	// S3 storage bucket name
	// Required: true
	BucketName *string `json:"bucketName"`

	// CA Certificate
	CaCert string `json:"caCert,omitempty"`

	// AWS cloud account credentials
	// Required: true
	Credentials *V1AwsCloudAccount `json:"credentials"`

	// AWS region name
	// Required: true
	Region *string `json:"region"`

	// s3 force path style
	S3ForcePathStyle *bool `json:"s3ForcePathStyle,omitempty"`

	// Custom hosted S3 URL
	S3URL string `json:"s3Url,omitempty"`

	// Set to 'true', to use Restic plugin for the backup
	UseRestic *bool `json:"useRestic,omitempty"`
}

// Validate validates this v1 s3 storage config
func (m *V1S3StorageConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucketName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1S3StorageConfig) validateBucketName(formats strfmt.Registry) error {

	if err := validate.Required("bucketName", "body", m.BucketName); err != nil {
		return err
	}

	return nil
}

func (m *V1S3StorageConfig) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("credentials", "body", m.Credentials); err != nil {
		return err
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *V1S3StorageConfig) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 s3 storage config based on the context it is used
func (m *V1S3StorageConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1S3StorageConfig) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {

		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1S3StorageConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1S3StorageConfig) UnmarshalBinary(b []byte) error {
	var res V1S3StorageConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
