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

// V1SearchFilterKeyValueOperator v1 search filter key value operator
//
// swagger:model v1SearchFilterKeyValueOperator
type V1SearchFilterKeyValueOperator string

func NewV1SearchFilterKeyValueOperator(value V1SearchFilterKeyValueOperator) *V1SearchFilterKeyValueOperator {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1SearchFilterKeyValueOperator.
func (m V1SearchFilterKeyValueOperator) Pointer() *V1SearchFilterKeyValueOperator {
	return &m
}

const (

	// V1SearchFilterKeyValueOperatorEq captures enum value "eq"
	V1SearchFilterKeyValueOperatorEq V1SearchFilterKeyValueOperator = "eq"
)

// for schema
var v1SearchFilterKeyValueOperatorEnum []interface{}

func init() {
	var res []V1SearchFilterKeyValueOperator
	if err := json.Unmarshal([]byte(`["eq"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SearchFilterKeyValueOperatorEnum = append(v1SearchFilterKeyValueOperatorEnum, v)
	}
}

func (m V1SearchFilterKeyValueOperator) validateV1SearchFilterKeyValueOperatorEnum(path, location string, value V1SearchFilterKeyValueOperator) error {
	if err := validate.EnumCase(path, location, value, v1SearchFilterKeyValueOperatorEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 search filter key value operator
func (m V1SearchFilterKeyValueOperator) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1SearchFilterKeyValueOperatorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 search filter key value operator based on context it is used
func (m V1SearchFilterKeyValueOperator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
