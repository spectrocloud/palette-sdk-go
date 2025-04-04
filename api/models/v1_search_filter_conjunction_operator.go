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

// V1SearchFilterConjunctionOperator v1 search filter conjunction operator
//
// swagger:model v1SearchFilterConjunctionOperator
type V1SearchFilterConjunctionOperator string

func NewV1SearchFilterConjunctionOperator(value V1SearchFilterConjunctionOperator) *V1SearchFilterConjunctionOperator {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1SearchFilterConjunctionOperator.
func (m V1SearchFilterConjunctionOperator) Pointer() *V1SearchFilterConjunctionOperator {
	return &m
}

const (

	// V1SearchFilterConjunctionOperatorAnd captures enum value "and"
	V1SearchFilterConjunctionOperatorAnd V1SearchFilterConjunctionOperator = "and"

	// V1SearchFilterConjunctionOperatorOr captures enum value "or"
	V1SearchFilterConjunctionOperatorOr V1SearchFilterConjunctionOperator = "or"
)

// for schema
var v1SearchFilterConjunctionOperatorEnum []interface{}

func init() {
	var res []V1SearchFilterConjunctionOperator
	if err := json.Unmarshal([]byte(`["and","or"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SearchFilterConjunctionOperatorEnum = append(v1SearchFilterConjunctionOperatorEnum, v)
	}
}

func (m V1SearchFilterConjunctionOperator) validateV1SearchFilterConjunctionOperatorEnum(path, location string, value V1SearchFilterConjunctionOperator) error {
	if err := validate.EnumCase(path, location, value, v1SearchFilterConjunctionOperatorEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 search filter conjunction operator
func (m V1SearchFilterConjunctionOperator) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1SearchFilterConjunctionOperatorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 search filter conjunction operator based on context it is used
func (m V1SearchFilterConjunctionOperator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
