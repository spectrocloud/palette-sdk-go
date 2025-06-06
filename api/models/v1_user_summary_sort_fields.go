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

// V1UserSummarySortFields v1 user summary sort fields
//
// swagger:model v1UserSummarySortFields
type V1UserSummarySortFields string

func NewV1UserSummarySortFields(value V1UserSummarySortFields) *V1UserSummarySortFields {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1UserSummarySortFields.
func (m V1UserSummarySortFields) Pointer() *V1UserSummarySortFields {
	return &m
}

const (

	// V1UserSummarySortFieldsName captures enum value "name"
	V1UserSummarySortFieldsName V1UserSummarySortFields = "name"

	// V1UserSummarySortFieldsCreationTimestamp captures enum value "creationTimestamp"
	V1UserSummarySortFieldsCreationTimestamp V1UserSummarySortFields = "creationTimestamp"
)

// for schema
var v1UserSummarySortFieldsEnum []interface{}

func init() {
	var res []V1UserSummarySortFields
	if err := json.Unmarshal([]byte(`["name","creationTimestamp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1UserSummarySortFieldsEnum = append(v1UserSummarySortFieldsEnum, v)
	}
}

func (m V1UserSummarySortFields) validateV1UserSummarySortFieldsEnum(path, location string, value V1UserSummarySortFields) error {
	if err := validate.EnumCase(path, location, value, v1UserSummarySortFieldsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 user summary sort fields
func (m V1UserSummarySortFields) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1UserSummarySortFieldsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 user summary sort fields based on context it is used
func (m V1UserSummarySortFields) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
