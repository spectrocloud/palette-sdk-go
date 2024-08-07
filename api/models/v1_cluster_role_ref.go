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

// V1ClusterRoleRef Cluster role ref
//
// swagger:model v1ClusterRoleRef
type V1ClusterRoleRef struct {

	// kind
	// Enum: [Role ClusterRole]
	Kind string `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 cluster role ref
func (m *V1ClusterRoleRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1ClusterRoleRefTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Role","ClusterRole"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ClusterRoleRefTypeKindPropEnum = append(v1ClusterRoleRefTypeKindPropEnum, v)
	}
}

const (

	// V1ClusterRoleRefKindRole captures enum value "Role"
	V1ClusterRoleRefKindRole string = "Role"

	// V1ClusterRoleRefKindClusterRole captures enum value "ClusterRole"
	V1ClusterRoleRefKindClusterRole string = "ClusterRole"
)

// prop value enum
func (m *V1ClusterRoleRef) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1ClusterRoleRefTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1ClusterRoleRef) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterRoleRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterRoleRef) UnmarshalBinary(b []byte) error {
	var res V1ClusterRoleRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
