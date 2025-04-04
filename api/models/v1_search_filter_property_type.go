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

// V1SearchFilterPropertyType v1 search filter property type
//
// swagger:model v1SearchFilterPropertyType
type V1SearchFilterPropertyType string

func NewV1SearchFilterPropertyType(value V1SearchFilterPropertyType) *V1SearchFilterPropertyType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1SearchFilterPropertyType.
func (m V1SearchFilterPropertyType) Pointer() *V1SearchFilterPropertyType {
	return &m
}

const (

	// V1SearchFilterPropertyTypeString captures enum value "string"
	V1SearchFilterPropertyTypeString V1SearchFilterPropertyType = "string"

	// V1SearchFilterPropertyTypeInt captures enum value "int"
	V1SearchFilterPropertyTypeInt V1SearchFilterPropertyType = "int"

	// V1SearchFilterPropertyTypeFloat captures enum value "float"
	V1SearchFilterPropertyTypeFloat V1SearchFilterPropertyType = "float"

	// V1SearchFilterPropertyTypeBool captures enum value "bool"
	V1SearchFilterPropertyTypeBool V1SearchFilterPropertyType = "bool"

	// V1SearchFilterPropertyTypeDate captures enum value "date"
	V1SearchFilterPropertyTypeDate V1SearchFilterPropertyType = "date"

	// V1SearchFilterPropertyTypeKeyValue captures enum value "keyValue"
	V1SearchFilterPropertyTypeKeyValue V1SearchFilterPropertyType = "keyValue"
)

// for schema
var v1SearchFilterPropertyTypeEnum []interface{}

func init() {
	var res []V1SearchFilterPropertyType
	if err := json.Unmarshal([]byte(`["string","int","float","bool","date","keyValue"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SearchFilterPropertyTypeEnum = append(v1SearchFilterPropertyTypeEnum, v)
	}
}

func (m V1SearchFilterPropertyType) validateV1SearchFilterPropertyTypeEnum(path, location string, value V1SearchFilterPropertyType) error {
	if err := validate.EnumCase(path, location, value, v1SearchFilterPropertyTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 search filter property type
func (m V1SearchFilterPropertyType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1SearchFilterPropertyTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 search filter property type based on context it is used
func (m V1SearchFilterPropertyType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
