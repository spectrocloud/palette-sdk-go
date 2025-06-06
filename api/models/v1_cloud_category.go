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

// V1CloudCategory Cloud category description
//
// swagger:model v1CloudCategory
type V1CloudCategory string

func NewV1CloudCategory(value V1CloudCategory) *V1CloudCategory {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1CloudCategory.
func (m V1CloudCategory) Pointer() *V1CloudCategory {
	return &m
}

const (

	// V1CloudCategoryDatacenter captures enum value "datacenter"
	V1CloudCategoryDatacenter V1CloudCategory = "datacenter"

	// V1CloudCategoryCloud captures enum value "cloud"
	V1CloudCategoryCloud V1CloudCategory = "cloud"

	// V1CloudCategoryEdge captures enum value "edge"
	V1CloudCategoryEdge V1CloudCategory = "edge"
)

// for schema
var v1CloudCategoryEnum []interface{}

func init() {
	var res []V1CloudCategory
	if err := json.Unmarshal([]byte(`["datacenter","cloud","edge"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1CloudCategoryEnum = append(v1CloudCategoryEnum, v)
	}
}

func (m V1CloudCategory) validateV1CloudCategoryEnum(path, location string, value V1CloudCategory) error {
	if err := validate.EnumCase(path, location, value, v1CloudCategoryEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 cloud category
func (m V1CloudCategory) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1CloudCategoryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 cloud category based on context it is used
func (m V1CloudCategory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
