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

// V1OsType v1 os type
//
// swagger:model v1OsType
type V1OsType string

const (

	// V1OsTypeLinux captures enum value "Linux"
	V1OsTypeLinux V1OsType = "Linux"

	// V1OsTypeWindows captures enum value "Windows"
	V1OsTypeWindows V1OsType = "Windows"
)

// for schema
var v1OsTypeEnum []interface{}

func init() {
	var res []V1OsType
	if err := json.Unmarshal([]byte(`["Linux","Windows"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1OsTypeEnum = append(v1OsTypeEnum, v)
	}
}

func (m V1OsType) validateV1OsTypeEnum(path, location string, value V1OsType) error {
	if err := validate.EnumCase(path, location, value, v1OsTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 os type
func (m V1OsType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1OsTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}