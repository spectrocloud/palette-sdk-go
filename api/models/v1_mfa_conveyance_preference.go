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

// V1MfaConveyancePreference v1 mfa conveyance preference
//
// swagger:model v1MfaConveyancePreference
type V1MfaConveyancePreference string

const (

	// V1MfaConveyancePreferenceNone captures enum value "none"
	V1MfaConveyancePreferenceNone V1MfaConveyancePreference = "none"

	// V1MfaConveyancePreferenceIndirect captures enum value "indirect"
	V1MfaConveyancePreferenceIndirect V1MfaConveyancePreference = "indirect"

	// V1MfaConveyancePreferenceDirect captures enum value "direct"
	V1MfaConveyancePreferenceDirect V1MfaConveyancePreference = "direct"

	// V1MfaConveyancePreferenceEnterprise captures enum value "enterprise"
	V1MfaConveyancePreferenceEnterprise V1MfaConveyancePreference = "enterprise"
)

// for schema
var v1MfaConveyancePreferenceEnum []interface{}

func init() {
	var res []V1MfaConveyancePreference
	if err := json.Unmarshal([]byte(`["none","indirect","direct","enterprise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1MfaConveyancePreferenceEnum = append(v1MfaConveyancePreferenceEnum, v)
	}
}

func (m V1MfaConveyancePreference) validateV1MfaConveyancePreferenceEnum(path, location string, value V1MfaConveyancePreference) error {
	if err := validate.EnumCase(path, location, value, v1MfaConveyancePreferenceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 mfa conveyance preference
func (m V1MfaConveyancePreference) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1MfaConveyancePreferenceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
