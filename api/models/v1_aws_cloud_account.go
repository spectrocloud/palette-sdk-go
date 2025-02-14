// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1AwsCloudAccount AWS cloud account which includes access key and secret key in case of 'secret' credentials type. It includes policyARNS, ARN and externalId in case of sts. Partition is a group of AWS Region and Service objects
//
// swagger:model v1AwsCloudAccount
type V1AwsCloudAccount struct {

	// AWS account access key
	AccessKey string `json:"accessKey,omitempty"`

	// credential type
	CredentialType V1AwsCloudAccountCredentialType `json:"credentialType,omitempty"`

	// AWS accounts are scoped to a single partition. Allowed values [aws, aws-us-gov], Default values
	// Enum: [aws aws-us-gov aws-iso aws-iso-b]
	Partition *string `json:"partition,omitempty"`

	// List of policy ARNs required in case of credentialType sts.
	PolicyARNs []string `json:"policyARNs"`

	// AWS account secret key
	SecretKey string `json:"secretKey,omitempty"`

	// secret spec
	SecretSpec *V1AwsSecretSpec `json:"secretSpec,omitempty"`

	// AWS account secret token; in case of aws-iso and aws-iso-b
	SecretToken string `json:"secretToken,omitempty"`

	// AWS STS credentials in case of credentialType sts, will be empty in case of credential type secret
	Sts *V1AwsStsCredentials `json:"sts,omitempty"`
}

// Validate validates this v1 aws cloud account
func (m *V1AwsCloudAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsCloudAccount) validateCredentialType(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialType) { // not required
		return nil
	}

	if err := m.CredentialType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("credentialType")
		}
		return err
	}

	return nil
}

var v1AwsCloudAccountTypePartitionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aws","aws-us-gov","aws-iso","aws-iso-b"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1AwsCloudAccountTypePartitionPropEnum = append(v1AwsCloudAccountTypePartitionPropEnum, v)
	}
}

const (

	// V1AwsCloudAccountPartitionAws captures enum value "aws"
	V1AwsCloudAccountPartitionAws string = "aws"

	// V1AwsCloudAccountPartitionAwsUsGov captures enum value "aws-us-gov"
	V1AwsCloudAccountPartitionAwsUsGov string = "aws-us-gov"

	// V1AwsCloudAccountPartitionAwsIso captures enum value "aws-iso"
	V1AwsCloudAccountPartitionAwsIso string = "aws-iso"

	// V1AwsCloudAccountPartitionAwsIsob captures enum value "aws-iso-b"
	V1AwsCloudAccountPartitionAwsIsob string = "aws-iso-b"
)

// prop value enum
func (m *V1AwsCloudAccount) validatePartitionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1AwsCloudAccountTypePartitionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1AwsCloudAccount) validatePartition(formats strfmt.Registry) error {

	if swag.IsZero(m.Partition) { // not required
		return nil
	}

	// value enum
	if err := m.validatePartitionEnum("partition", "body", *m.Partition); err != nil {
		return err
	}

	return nil
}

func (m *V1AwsCloudAccount) validateSecretSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.SecretSpec) { // not required
		return nil
	}

	if m.SecretSpec != nil {
		if err := m.SecretSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretSpec")
			}
			return err
		}
	}

	return nil
}

func (m *V1AwsCloudAccount) validateSts(formats strfmt.Registry) error {

	if swag.IsZero(m.Sts) { // not required
		return nil
	}

	if m.Sts != nil {
		if err := m.Sts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsCloudAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsCloudAccount) UnmarshalBinary(b []byte) error {
	var res V1AwsCloudAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
