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

// V1ClusterRepaveSource v1 cluster repave source
//
// swagger:model v1ClusterRepaveSource
type V1ClusterRepaveSource string

func NewV1ClusterRepaveSource(value V1ClusterRepaveSource) *V1ClusterRepaveSource {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1ClusterRepaveSource.
func (m V1ClusterRepaveSource) Pointer() *V1ClusterRepaveSource {
	return &m
}

const (

	// V1ClusterRepaveSourceUser captures enum value "user"
	V1ClusterRepaveSourceUser V1ClusterRepaveSource = "user"

	// V1ClusterRepaveSourceHubble captures enum value "hubble"
	V1ClusterRepaveSourceHubble V1ClusterRepaveSource = "hubble"

	// V1ClusterRepaveSourcePalette captures enum value "palette"
	V1ClusterRepaveSourcePalette V1ClusterRepaveSource = "palette"

	// V1ClusterRepaveSourceStylus captures enum value "stylus"
	V1ClusterRepaveSourceStylus V1ClusterRepaveSource = "stylus"
)

// for schema
var v1ClusterRepaveSourceEnum []interface{}

func init() {
	var res []V1ClusterRepaveSource
	if err := json.Unmarshal([]byte(`["user","hubble","palette","stylus"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ClusterRepaveSourceEnum = append(v1ClusterRepaveSourceEnum, v)
	}
}

func (m V1ClusterRepaveSource) validateV1ClusterRepaveSourceEnum(path, location string, value V1ClusterRepaveSource) error {
	if err := validate.EnumCase(path, location, value, v1ClusterRepaveSourceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 cluster repave source
func (m V1ClusterRepaveSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ClusterRepaveSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 cluster repave source based on context it is used
func (m V1ClusterRepaveSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
