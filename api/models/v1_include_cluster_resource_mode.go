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

// V1IncludeClusterResourceMode Specifies the scope of cluster-wide resources to include in the backup based on the flag'--include-cluster-resources':
// - "Always": --include-cluster-resources=true, which includes all cluster-wide resources; restores only on the original cluster.
// - "Auto": don't specify --include-cluster-resources, which exclude general cluster-wide resources, but includes PersistentVolumes linked to selected namespaces.
// - "Never": --include-cluster-resources=false, exclude all cluster-wide resources, including PersistentVolumes.
//
// swagger:model v1IncludeClusterResourceMode
type V1IncludeClusterResourceMode string

func NewV1IncludeClusterResourceMode(value V1IncludeClusterResourceMode) *V1IncludeClusterResourceMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1IncludeClusterResourceMode.
func (m V1IncludeClusterResourceMode) Pointer() *V1IncludeClusterResourceMode {
	return &m
}

const (

	// V1IncludeClusterResourceModeAlways captures enum value "Always"
	V1IncludeClusterResourceModeAlways V1IncludeClusterResourceMode = "Always"

	// V1IncludeClusterResourceModeAuto captures enum value "Auto"
	V1IncludeClusterResourceModeAuto V1IncludeClusterResourceMode = "Auto"

	// V1IncludeClusterResourceModeNever captures enum value "Never"
	V1IncludeClusterResourceModeNever V1IncludeClusterResourceMode = "Never"
)

// for schema
var v1IncludeClusterResourceModeEnum []interface{}

func init() {
	var res []V1IncludeClusterResourceMode
	if err := json.Unmarshal([]byte(`["Always","Auto","Never"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1IncludeClusterResourceModeEnum = append(v1IncludeClusterResourceModeEnum, v)
	}
}

func (m V1IncludeClusterResourceMode) validateV1IncludeClusterResourceModeEnum(path, location string, value V1IncludeClusterResourceMode) error {
	if err := validate.EnumCase(path, location, value, v1IncludeClusterResourceModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 include cluster resource mode
func (m V1IncludeClusterResourceMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1IncludeClusterResourceModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 include cluster resource mode based on context it is used
func (m V1IncludeClusterResourceMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
