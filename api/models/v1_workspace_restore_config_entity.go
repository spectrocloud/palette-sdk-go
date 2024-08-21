// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1WorkspaceRestoreConfigEntity Cluster restore config
//
// swagger:model v1WorkspaceRestoreConfigEntity
type V1WorkspaceRestoreConfigEntity struct {

	// backup request Uid
	// Required: true
	BackupRequestUID *string `json:"backupRequestUid"`

	// restore configs
	// Unique: true
	RestoreConfigs []*V1WorkspaceRestoreConfig `json:"restoreConfigs"`
}

// Validate validates this v1 workspace restore config entity
func (m *V1WorkspaceRestoreConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupRequestUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoreConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceRestoreConfigEntity) validateBackupRequestUID(formats strfmt.Registry) error {

	if err := validate.Required("backupRequestUid", "body", m.BackupRequestUID); err != nil {
		return err
	}

	return nil
}

func (m *V1WorkspaceRestoreConfigEntity) validateRestoreConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.RestoreConfigs) { // not required
		return nil
	}

	if err := validate.UniqueItems("restoreConfigs", "body", m.RestoreConfigs); err != nil {
		return err
	}

	for i := 0; i < len(m.RestoreConfigs); i++ {
		if swag.IsZero(m.RestoreConfigs[i]) { // not required
			continue
		}

		if m.RestoreConfigs[i] != nil {
			if err := m.RestoreConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restoreConfigs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspaceRestoreConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceRestoreConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceRestoreConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}