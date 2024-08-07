// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1RegistryAuth Auth credentials of the registry
//
// swagger:model v1RegistryAuth
type V1RegistryAuth struct {

	// password
	// Format: password
	Password strfmt.Password `json:"password,omitempty"`

	// tls
	TLS *V1TLSConfiguration `json:"tls,omitempty"`

	// token
	// Format: password
	Token strfmt.Password `json:"token,omitempty"`

	// type
	// Enum: [noAuth basic token]
	Type string `json:"type,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this v1 registry auth
func (m *V1RegistryAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RegistryAuth) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1RegistryAuth) validateTLS(formats strfmt.Registry) error {

	if swag.IsZero(m.TLS) { // not required
		return nil
	}

	if m.TLS != nil {
		if err := m.TLS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

func (m *V1RegistryAuth) validateToken(formats strfmt.Registry) error {

	if swag.IsZero(m.Token) { // not required
		return nil
	}

	if err := validate.FormatOf("token", "body", "password", m.Token.String(), formats); err != nil {
		return err
	}

	return nil
}

var v1RegistryAuthTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["noAuth","basic","token"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1RegistryAuthTypeTypePropEnum = append(v1RegistryAuthTypeTypePropEnum, v)
	}
}

const (

	// V1RegistryAuthTypeNoAuth captures enum value "noAuth"
	V1RegistryAuthTypeNoAuth string = "noAuth"

	// V1RegistryAuthTypeBasic captures enum value "basic"
	V1RegistryAuthTypeBasic string = "basic"

	// V1RegistryAuthTypeToken captures enum value "token"
	V1RegistryAuthTypeToken string = "token"
)

// prop value enum
func (m *V1RegistryAuth) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1RegistryAuthTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1RegistryAuth) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RegistryAuth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RegistryAuth) UnmarshalBinary(b []byte) error {
	var res V1RegistryAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
