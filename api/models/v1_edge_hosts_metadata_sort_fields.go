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

// V1EdgeHostsMetadataSortFields v1 edge hosts metadata sort fields
//
// swagger:model v1EdgeHostsMetadataSortFields
type V1EdgeHostsMetadataSortFields string

func NewV1EdgeHostsMetadataSortFields(value V1EdgeHostsMetadataSortFields) *V1EdgeHostsMetadataSortFields {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1EdgeHostsMetadataSortFields.
func (m V1EdgeHostsMetadataSortFields) Pointer() *V1EdgeHostsMetadataSortFields {
	return &m
}

const (

	// V1EdgeHostsMetadataSortFieldsName captures enum value "name"
	V1EdgeHostsMetadataSortFieldsName V1EdgeHostsMetadataSortFields = "name"

	// V1EdgeHostsMetadataSortFieldsState captures enum value "state"
	V1EdgeHostsMetadataSortFieldsState V1EdgeHostsMetadataSortFields = "state"

	// V1EdgeHostsMetadataSortFieldsCreationTimestamp captures enum value "creationTimestamp"
	V1EdgeHostsMetadataSortFieldsCreationTimestamp V1EdgeHostsMetadataSortFields = "creationTimestamp"

	// V1EdgeHostsMetadataSortFieldsLastModifiedTimestamp captures enum value "lastModifiedTimestamp"
	V1EdgeHostsMetadataSortFieldsLastModifiedTimestamp V1EdgeHostsMetadataSortFields = "lastModifiedTimestamp"
)

// for schema
var v1EdgeHostsMetadataSortFieldsEnum []interface{}

func init() {
	var res []V1EdgeHostsMetadataSortFields
	if err := json.Unmarshal([]byte(`["name","state","creationTimestamp","lastModifiedTimestamp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1EdgeHostsMetadataSortFieldsEnum = append(v1EdgeHostsMetadataSortFieldsEnum, v)
	}
}

func (m V1EdgeHostsMetadataSortFields) validateV1EdgeHostsMetadataSortFieldsEnum(path, location string, value V1EdgeHostsMetadataSortFields) error {
	if err := validate.EnumCase(path, location, value, v1EdgeHostsMetadataSortFieldsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 edge hosts metadata sort fields
func (m V1EdgeHostsMetadataSortFields) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1EdgeHostsMetadataSortFieldsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 edge hosts metadata sort fields based on context it is used
func (m V1EdgeHostsMetadataSortFields) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
