// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1NonFipsConfig Non-FIPS configuration
//
// swagger:model v1NonFipsConfig
type V1NonFipsConfig struct {

	// enable or disable the non FIPS complaint
	// Enum: ["nonFipsEnabled","nonFipsDisabled"]
	Mode *string `json:"mode,omitempty"`
}

// Validate validates this v1 non fips config
func (m *V1NonFipsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1NonFipsConfigTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nonFipsEnabled","nonFipsDisabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1NonFipsConfigTypeModePropEnum = append(v1NonFipsConfigTypeModePropEnum, v)
	}
}

const (

	// V1NonFipsConfigModeNonFipsEnabled captures enum value "nonFipsEnabled"
	V1NonFipsConfigModeNonFipsEnabled string = "nonFipsEnabled"

	// V1NonFipsConfigModeNonFipsDisabled captures enum value "nonFipsDisabled"
	V1NonFipsConfigModeNonFipsDisabled string = "nonFipsDisabled"
)

// prop value enum
func (m *V1NonFipsConfig) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1NonFipsConfigTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1NonFipsConfig) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 non fips config based on context it is used
func (m *V1NonFipsConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NonFipsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NonFipsConfig) UnmarshalBinary(b []byte) error {
	var res V1NonFipsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
