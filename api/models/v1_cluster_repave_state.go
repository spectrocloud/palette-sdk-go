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

// V1ClusterRepaveState v1 cluster repave state
//
// swagger:model v1ClusterRepaveState
type V1ClusterRepaveState string

func NewV1ClusterRepaveState(value V1ClusterRepaveState) *V1ClusterRepaveState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1ClusterRepaveState.
func (m V1ClusterRepaveState) Pointer() *V1ClusterRepaveState {
	return &m
}

const (

	// V1ClusterRepaveStatePending captures enum value "Pending"
	V1ClusterRepaveStatePending V1ClusterRepaveState = "Pending"

	// V1ClusterRepaveStateApproved captures enum value "Approved"
	V1ClusterRepaveStateApproved V1ClusterRepaveState = "Approved"

	// V1ClusterRepaveStateReverted captures enum value "Reverted"
	V1ClusterRepaveStateReverted V1ClusterRepaveState = "Reverted"
)

// for schema
var v1ClusterRepaveStateEnum []interface{}

func init() {
	var res []V1ClusterRepaveState
	if err := json.Unmarshal([]byte(`["Pending","Approved","Reverted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ClusterRepaveStateEnum = append(v1ClusterRepaveStateEnum, v)
	}
}

func (m V1ClusterRepaveState) validateV1ClusterRepaveStateEnum(path, location string, value V1ClusterRepaveState) error {
	if err := validate.EnumCase(path, location, value, v1ClusterRepaveStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 cluster repave state
func (m V1ClusterRepaveState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ClusterRepaveStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 cluster repave state based on context it is used
func (m V1ClusterRepaveState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
