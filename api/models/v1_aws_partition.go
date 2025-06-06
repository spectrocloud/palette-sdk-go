// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1AwsPartition AWS accounts are scoped to a single partition. Allowed values [aws, aws-us-gov], Default values
//
// swagger:model v1AwsPartition
type V1AwsPartition string

func NewV1AwsPartition(value V1AwsPartition) *V1AwsPartition {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1AwsPartition.
func (m V1AwsPartition) Pointer() *V1AwsPartition {
	return &m
}

const (

	// V1AwsPartitionAws captures enum value "aws"
	V1AwsPartitionAws V1AwsPartition = "aws"

	// V1AwsPartitionAwsDashUsDashGov captures enum value "aws-us-gov"
	V1AwsPartitionAwsDashUsDashGov V1AwsPartition = "aws-us-gov"
)

// for schema
var v1AwsPartitionEnum []interface{}

func init() {
	var res []V1AwsPartition
	if err := json.Unmarshal([]byte(`["aws","aws-us-gov"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1AwsPartitionEnum = append(v1AwsPartitionEnum, v)
	}
}

func (m V1AwsPartition) validateV1AwsPartitionEnum(path, location string, value V1AwsPartition) error {
	if err := validate.EnumCase(path, location, value, v1AwsPartitionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 aws partition
func (m V1AwsPartition) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1AwsPartitionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 aws partition based on context it is used
func (m V1AwsPartition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
