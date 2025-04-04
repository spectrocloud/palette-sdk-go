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

// V1SortOrder v1 sort order
//
// swagger:model v1SortOrder
type V1SortOrder string

func NewV1SortOrder(value V1SortOrder) *V1SortOrder {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1SortOrder.
func (m V1SortOrder) Pointer() *V1SortOrder {
	return &m
}

const (

	// V1SortOrderAsc captures enum value "asc"
	V1SortOrderAsc V1SortOrder = "asc"

	// V1SortOrderDesc captures enum value "desc"
	V1SortOrderDesc V1SortOrder = "desc"
)

// for schema
var v1SortOrderEnum []interface{}

func init() {
	var res []V1SortOrder
	if err := json.Unmarshal([]byte(`["asc","desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SortOrderEnum = append(v1SortOrderEnum, v)
	}
}

func (m V1SortOrder) validateV1SortOrderEnum(path, location string, value V1SortOrder) error {
	if err := validate.EnumCase(path, location, value, v1SortOrderEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 sort order
func (m V1SortOrder) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1SortOrderEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 sort order based on context it is used
func (m V1SortOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
