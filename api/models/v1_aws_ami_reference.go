// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1AwsAmiReference AMI is the reference to the AMI from which to create the machine instance
//
// swagger:model v1AwsAmiReference
type V1AwsAmiReference struct {

	// EKSOptimizedLookupType If specified, will look up an EKS Optimized image in SSM Parameter store
	// Enum: ["AmazonLinux","AmazonLinuxGPU"]
	EksOptimizedLookupType string `json:"eksOptimizedLookupType,omitempty"`

	// ID of resource
	ID string `json:"id,omitempty"`
}

// Validate validates this v1 aws ami reference
func (m *V1AwsAmiReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEksOptimizedLookupType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1AwsAmiReferenceTypeEksOptimizedLookupTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AmazonLinux","AmazonLinuxGPU"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1AwsAmiReferenceTypeEksOptimizedLookupTypePropEnum = append(v1AwsAmiReferenceTypeEksOptimizedLookupTypePropEnum, v)
	}
}

const (

	// V1AwsAmiReferenceEksOptimizedLookupTypeAmazonLinux captures enum value "AmazonLinux"
	V1AwsAmiReferenceEksOptimizedLookupTypeAmazonLinux string = "AmazonLinux"

	// V1AwsAmiReferenceEksOptimizedLookupTypeAmazonLinuxGPU captures enum value "AmazonLinuxGPU"
	V1AwsAmiReferenceEksOptimizedLookupTypeAmazonLinuxGPU string = "AmazonLinuxGPU"
)

// prop value enum
func (m *V1AwsAmiReference) validateEksOptimizedLookupTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1AwsAmiReferenceTypeEksOptimizedLookupTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1AwsAmiReference) validateEksOptimizedLookupType(formats strfmt.Registry) error {
	if swag.IsZero(m.EksOptimizedLookupType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEksOptimizedLookupTypeEnum("eksOptimizedLookupType", "body", m.EksOptimizedLookupType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 aws ami reference based on context it is used
func (m *V1AwsAmiReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsAmiReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsAmiReference) UnmarshalBinary(b []byte) error {
	var res V1AwsAmiReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
