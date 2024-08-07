// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TenantClusterSettings v1 tenant cluster settings
//
// swagger:model v1TenantClusterSettings
type V1TenantClusterSettings struct {

	// nodes auto remediation setting
	NodesAutoRemediationSetting *V1NodesAutoRemediationSettings `json:"nodesAutoRemediationSetting,omitempty"`
}

// Validate validates this v1 tenant cluster settings
func (m *V1TenantClusterSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodesAutoRemediationSetting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TenantClusterSettings) validateNodesAutoRemediationSetting(formats strfmt.Registry) error {

	if swag.IsZero(m.NodesAutoRemediationSetting) { // not required
		return nil
	}

	if m.NodesAutoRemediationSetting != nil {
		if err := m.NodesAutoRemediationSetting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodesAutoRemediationSetting")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TenantClusterSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TenantClusterSettings) UnmarshalBinary(b []byte) error {
	var res V1TenantClusterSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
