// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ComplianceScanDriverSpec Compliance Scan driver spec
//
// swagger:model v1ComplianceScanDriverSpec
type V1ComplianceScanDriverSpec struct {

	// config
	Config *V1ComplianceScanConfig `json:"config,omitempty"`

	// is cluster config
	IsClusterConfig bool `json:"isClusterConfig,omitempty"`
}

// Validate validates this v1 compliance scan driver spec
func (m *V1ComplianceScanDriverSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ComplianceScanDriverSpec) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ComplianceScanDriverSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ComplianceScanDriverSpec) UnmarshalBinary(b []byte) error {
	var res V1ComplianceScanDriverSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
