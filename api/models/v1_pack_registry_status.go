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

// V1PackRegistryStatus Status of the pack registry
//
// swagger:model v1PackRegistryStatus
type V1PackRegistryStatus struct {

	// pack sync status
	PackSyncStatus *V1RegistrySyncStatus `json:"packSyncStatus,omitempty"`
}

// Validate validates this v1 pack registry status
func (m *V1PackRegistryStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackSyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackRegistryStatus) validatePackSyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.PackSyncStatus) { // not required
		return nil
	}

	if m.PackSyncStatus != nil {
		if err := m.PackSyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packSyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packSyncStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 pack registry status based on the context it is used
func (m *V1PackRegistryStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePackSyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackRegistryStatus) contextValidatePackSyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.PackSyncStatus != nil {

		if swag.IsZero(m.PackSyncStatus) { // not required
			return nil
		}

		if err := m.PackSyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packSyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packSyncStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PackRegistryStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackRegistryStatus) UnmarshalBinary(b []byte) error {
	var res V1PackRegistryStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
