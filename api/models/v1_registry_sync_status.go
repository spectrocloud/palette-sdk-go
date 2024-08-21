// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1RegistrySyncStatus Status of the registry sync
//
// swagger:model v1RegistrySyncStatus
type V1RegistrySyncStatus struct {

	// last run time
	// Format: date-time
	LastRunTime V1Time `json:"lastRunTime,omitempty"`

	// last synced time
	// Format: date-time
	LastSyncedTime V1Time `json:"lastSyncedTime,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this v1 registry sync status
func (m *V1RegistrySyncStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastRunTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSyncedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RegistrySyncStatus) validateLastRunTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastRunTime) { // not required
		return nil
	}

	if err := m.LastRunTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastRunTime")
		}
		return err
	}

	return nil
}

func (m *V1RegistrySyncStatus) validateLastSyncedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastSyncedTime) { // not required
		return nil
	}

	if err := m.LastSyncedTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastSyncedTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RegistrySyncStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RegistrySyncStatus) UnmarshalBinary(b []byte) error {
	var res V1RegistrySyncStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}