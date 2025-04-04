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

// V1ClusterType v1 cluster type
//
// swagger:model v1ClusterType
type V1ClusterType string

func NewV1ClusterType(value V1ClusterType) *V1ClusterType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1ClusterType.
func (m V1ClusterType) Pointer() *V1ClusterType {
	return &m
}

const (

	// V1ClusterTypePureManage captures enum value "PureManage"
	V1ClusterTypePureManage V1ClusterType = "PureManage"

	// V1ClusterTypePureAttach captures enum value "PureAttach"
	V1ClusterTypePureAttach V1ClusterType = "PureAttach"
)

// for schema
var v1ClusterTypeEnum []interface{}

func init() {
	var res []V1ClusterType
	if err := json.Unmarshal([]byte(`["PureManage","PureAttach"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ClusterTypeEnum = append(v1ClusterTypeEnum, v)
	}
}

func (m V1ClusterType) validateV1ClusterTypeEnum(path, location string, value V1ClusterType) error {
	if err := validate.EnumCase(path, location, value, v1ClusterTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 cluster type
func (m V1ClusterType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ClusterTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 cluster type based on context it is used
func (m V1ClusterType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
