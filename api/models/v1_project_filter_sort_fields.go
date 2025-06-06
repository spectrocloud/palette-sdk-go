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

// V1ProjectFilterSortFields v1 project filter sort fields
//
// swagger:model v1ProjectFilterSortFields
type V1ProjectFilterSortFields string

func NewV1ProjectFilterSortFields(value V1ProjectFilterSortFields) *V1ProjectFilterSortFields {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1ProjectFilterSortFields.
func (m V1ProjectFilterSortFields) Pointer() *V1ProjectFilterSortFields {
	return &m
}

const (

	// V1ProjectFilterSortFieldsName captures enum value "name"
	V1ProjectFilterSortFieldsName V1ProjectFilterSortFields = "name"

	// V1ProjectFilterSortFieldsCreationTimestamp captures enum value "creationTimestamp"
	V1ProjectFilterSortFieldsCreationTimestamp V1ProjectFilterSortFields = "creationTimestamp"

	// V1ProjectFilterSortFieldsLastModifiedTimestamp captures enum value "lastModifiedTimestamp"
	V1ProjectFilterSortFieldsLastModifiedTimestamp V1ProjectFilterSortFields = "lastModifiedTimestamp"
)

// for schema
var v1ProjectFilterSortFieldsEnum []interface{}

func init() {
	var res []V1ProjectFilterSortFields
	if err := json.Unmarshal([]byte(`["name","creationTimestamp","lastModifiedTimestamp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ProjectFilterSortFieldsEnum = append(v1ProjectFilterSortFieldsEnum, v)
	}
}

func (m V1ProjectFilterSortFields) validateV1ProjectFilterSortFieldsEnum(path, location string, value V1ProjectFilterSortFields) error {
	if err := validate.EnumCase(path, location, value, v1ProjectFilterSortFieldsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 project filter sort fields
func (m V1ProjectFilterSortFields) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ProjectFilterSortFieldsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 project filter sort fields based on context it is used
func (m V1ProjectFilterSortFields) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
