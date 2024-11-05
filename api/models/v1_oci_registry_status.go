// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OciRegistryStatus Status of the oci registry
//
// swagger:model v1OciRegistryStatus
type V1OciRegistryStatus struct {

	// sync status
	SyncStatus *V1RegistrySyncStatus `json:"syncStatus,omitempty"`
}

// Validate validates this v1 oci registry status
func (m *V1OciRegistryStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OciRegistryStatus) validateSyncStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.SyncStatus) { // not required
		return nil
	}

	if m.SyncStatus != nil {
		if err := m.SyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syncStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OciRegistryStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OciRegistryStatus) UnmarshalBinary(b []byte) error {
	var res V1OciRegistryStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
