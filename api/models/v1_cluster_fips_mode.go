// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1ClusterFipsMode v1 cluster fips mode
//
// swagger:model v1ClusterFipsMode
type V1ClusterFipsMode string

const (

	// V1ClusterFipsModeFull captures enum value "full"
	V1ClusterFipsModeFull V1ClusterFipsMode = "full"

	// V1ClusterFipsModeNone captures enum value "none"
	V1ClusterFipsModeNone V1ClusterFipsMode = "none"

	// V1ClusterFipsModePartial captures enum value "partial"
	V1ClusterFipsModePartial V1ClusterFipsMode = "partial"

	// V1ClusterFipsModeUnknown captures enum value "unknown"
	V1ClusterFipsModeUnknown V1ClusterFipsMode = "unknown"
)

// for schema
var v1ClusterFipsModeEnum []interface{}

func init() {
	var res []V1ClusterFipsMode
	if err := json.Unmarshal([]byte(`["full","none","partial","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ClusterFipsModeEnum = append(v1ClusterFipsModeEnum, v)
	}
}

func (m V1ClusterFipsMode) validateV1ClusterFipsModeEnum(path, location string, value V1ClusterFipsMode) error {
	if err := validate.EnumCase(path, location, value, v1ClusterFipsModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 cluster fips mode
func (m V1ClusterFipsMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ClusterFipsModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
