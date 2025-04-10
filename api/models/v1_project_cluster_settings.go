// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ProjectClusterSettings v1 project cluster settings
//
// swagger:model v1ProjectClusterSettings
type V1ProjectClusterSettings struct {

	// nodes auto remediation setting
	NodesAutoRemediationSetting *V1NodesAutoRemediationSettings `json:"nodesAutoRemediationSetting,omitempty"`

	// tenant cluster settings
	TenantClusterSettings *V1TenantClusterSettings `json:"tenantClusterSettings,omitempty"`
}

// Validate validates this v1 project cluster settings
func (m *V1ProjectClusterSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodesAutoRemediationSetting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantClusterSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectClusterSettings) validateNodesAutoRemediationSetting(formats strfmt.Registry) error {
	if swag.IsZero(m.NodesAutoRemediationSetting) { // not required
		return nil
	}

	if m.NodesAutoRemediationSetting != nil {
		if err := m.NodesAutoRemediationSetting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodesAutoRemediationSetting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodesAutoRemediationSetting")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectClusterSettings) validateTenantClusterSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.TenantClusterSettings) { // not required
		return nil
	}

	if m.TenantClusterSettings != nil {
		if err := m.TenantClusterSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenantClusterSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenantClusterSettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 project cluster settings based on the context it is used
func (m *V1ProjectClusterSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodesAutoRemediationSetting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenantClusterSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectClusterSettings) contextValidateNodesAutoRemediationSetting(ctx context.Context, formats strfmt.Registry) error {

	if m.NodesAutoRemediationSetting != nil {

		if swag.IsZero(m.NodesAutoRemediationSetting) { // not required
			return nil
		}

		if err := m.NodesAutoRemediationSetting.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodesAutoRemediationSetting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodesAutoRemediationSetting")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectClusterSettings) contextValidateTenantClusterSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.TenantClusterSettings != nil {

		if swag.IsZero(m.TenantClusterSettings) { // not required
			return nil
		}

		if err := m.TenantClusterSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenantClusterSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenantClusterSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectClusterSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectClusterSettings) UnmarshalBinary(b []byte) error {
	var res V1ProjectClusterSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
