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

// V1RoleSpec Role specifications
//
// swagger:model v1RoleSpec
type V1RoleSpec struct {

	// permissions
	// Unique: true
	Permissions []string `json:"permissions"`

	// scope
	Scope V1Scope `json:"scope,omitempty"`

	// type
	// Enum: [system user]
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 role spec
func (m *V1RoleSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
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

func (m *V1RoleSpec) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if err := validate.UniqueItems("permissions", "body", m.Permissions); err != nil {
		return err
	}

	return nil
}

func (m *V1RoleSpec) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if err := m.Scope.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		}
		return err
	}

	return nil
}

var v1RoleSpecTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["system","user"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1RoleSpecTypeTypePropEnum = append(v1RoleSpecTypeTypePropEnum, v)
	}
}

const (

	// V1RoleSpecTypeSystem captures enum value "system"
	V1RoleSpecTypeSystem string = "system"

	// V1RoleSpecTypeUser captures enum value "user"
	V1RoleSpecTypeUser string = "user"
)

// prop value enum
func (m *V1RoleSpec) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1RoleSpecTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1RoleSpec) validateType(formats strfmt.Registry) error {

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
func (m *V1RoleSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RoleSpec) UnmarshalBinary(b []byte) error {
	var res V1RoleSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
