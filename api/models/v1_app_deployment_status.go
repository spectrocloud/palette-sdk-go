// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AppDeploymentStatus Application deployment status
//
// swagger:model v1AppDeploymentStatus
type V1AppDeploymentStatus struct {

	// Application deployment tiers
	AppTiers []*V1ClusterPackStatus `json:"appTiers"`

	// lifecycle status
	LifecycleStatus *V1LifecycleStatus `json:"lifecycleStatus,omitempty"`

	// Application deployment state [ "Pending", "Deploying", "Deployed", "Updating" ]
	State string `json:"state,omitempty"`
}

// Validate validates this v1 app deployment status
func (m *V1AppDeploymentStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppTiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifecycleStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppDeploymentStatus) validateAppTiers(formats strfmt.Registry) error {
	if swag.IsZero(m.AppTiers) { // not required
		return nil
	}

	for i := 0; i < len(m.AppTiers); i++ {
		if swag.IsZero(m.AppTiers[i]) { // not required
			continue
		}

		if m.AppTiers[i] != nil {
			if err := m.AppTiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appTiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appTiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AppDeploymentStatus) validateLifecycleStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.LifecycleStatus) { // not required
		return nil
	}

	if m.LifecycleStatus != nil {
		if err := m.LifecycleStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lifecycleStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lifecycleStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 app deployment status based on the context it is used
func (m *V1AppDeploymentStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppTiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLifecycleStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppDeploymentStatus) contextValidateAppTiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppTiers); i++ {

		if m.AppTiers[i] != nil {

			if swag.IsZero(m.AppTiers[i]) { // not required
				return nil
			}

			if err := m.AppTiers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appTiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appTiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AppDeploymentStatus) contextValidateLifecycleStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.LifecycleStatus != nil {

		if swag.IsZero(m.LifecycleStatus) { // not required
			return nil
		}

		if err := m.LifecycleStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lifecycleStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lifecycleStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppDeploymentStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppDeploymentStatus) UnmarshalBinary(b []byte) error {
	var res V1AppDeploymentStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
