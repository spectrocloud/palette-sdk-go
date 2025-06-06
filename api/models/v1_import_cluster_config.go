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

// V1ImportClusterConfig v1 import cluster config
//
// swagger:model v1ImportClusterConfig
type V1ImportClusterConfig struct {

	// If the importMode is empty then cluster is imported with full permission mode. By default importMode is empty and cluster is imported in full permission mode.
	// Enum: ["read-only"]
	ImportMode string `json:"importMode,omitempty"`

	// Cluster proxy settings
	Proxy *V1ClusterProxySpec `json:"proxy,omitempty"`
}

// Validate validates this v1 import cluster config
func (m *V1ImportClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImportMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1ImportClusterConfigTypeImportModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["read-only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ImportClusterConfigTypeImportModePropEnum = append(v1ImportClusterConfigTypeImportModePropEnum, v)
	}
}

const (

	// V1ImportClusterConfigImportModeReadDashOnly captures enum value "read-only"
	V1ImportClusterConfigImportModeReadDashOnly string = "read-only"
)

// prop value enum
func (m *V1ImportClusterConfig) validateImportModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1ImportClusterConfigTypeImportModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1ImportClusterConfig) validateImportMode(formats strfmt.Registry) error {
	if swag.IsZero(m.ImportMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateImportModeEnum("importMode", "body", m.ImportMode); err != nil {
		return err
	}

	return nil
}

func (m *V1ImportClusterConfig) validateProxy(formats strfmt.Registry) error {
	if swag.IsZero(m.Proxy) { // not required
		return nil
	}

	if m.Proxy != nil {
		if err := m.Proxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 import cluster config based on the context it is used
func (m *V1ImportClusterConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ImportClusterConfig) contextValidateProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.Proxy != nil {

		if swag.IsZero(m.Proxy) { // not required
			return nil
		}

		if err := m.Proxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ImportClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImportClusterConfig) UnmarshalBinary(b []byte) error {
	var res V1ImportClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
