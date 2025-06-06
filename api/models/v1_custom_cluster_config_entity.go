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

// V1CustomClusterConfigEntity v1 custom cluster config entity
//
// swagger:model v1CustomClusterConfigEntity
type V1CustomClusterConfigEntity struct {

	// location
	Location *V1ClusterLocation `json:"location,omitempty"`

	// machine management config
	MachineManagementConfig *V1MachineManagementConfig `json:"machineManagementConfig,omitempty"`

	// resources
	Resources *V1ClusterResourcesEntity `json:"resources,omitempty"`
}

// Validate validates this v1 custom cluster config entity
func (m *V1CustomClusterConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineManagementConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CustomClusterConfigEntity) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *V1CustomClusterConfigEntity) validateMachineManagementConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.MachineManagementConfig) { // not required
		return nil
	}

	if m.MachineManagementConfig != nil {
		if err := m.MachineManagementConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineManagementConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineManagementConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1CustomClusterConfigEntity) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 custom cluster config entity based on the context it is used
func (m *V1CustomClusterConfigEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachineManagementConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CustomClusterConfigEntity) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {

		if swag.IsZero(m.Location) { // not required
			return nil
		}

		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *V1CustomClusterConfigEntity) contextValidateMachineManagementConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.MachineManagementConfig != nil {

		if swag.IsZero(m.MachineManagementConfig) { // not required
			return nil
		}

		if err := m.MachineManagementConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineManagementConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineManagementConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1CustomClusterConfigEntity) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if m.Resources != nil {

		if swag.IsZero(m.Resources) { // not required
			return nil
		}

		if err := m.Resources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CustomClusterConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CustomClusterConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1CustomClusterConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
