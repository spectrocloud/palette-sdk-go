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

// V1WorkspaceBackupStatus Workspace backup status
//
// swagger:model v1WorkspaceBackupStatus
type V1WorkspaceBackupStatus struct {

	// workspace backup statuses
	WorkspaceBackupStatuses []*V1WorkspaceBackupStatusMeta `json:"workspaceBackupStatuses"`
}

// Validate validates this v1 workspace backup status
func (m *V1WorkspaceBackupStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkspaceBackupStatuses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceBackupStatus) validateWorkspaceBackupStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkspaceBackupStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.WorkspaceBackupStatuses); i++ {
		if swag.IsZero(m.WorkspaceBackupStatuses[i]) { // not required
			continue
		}

		if m.WorkspaceBackupStatuses[i] != nil {
			if err := m.WorkspaceBackupStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspaceBackupStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workspaceBackupStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 workspace backup status based on the context it is used
func (m *V1WorkspaceBackupStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkspaceBackupStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceBackupStatus) contextValidateWorkspaceBackupStatuses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WorkspaceBackupStatuses); i++ {

		if m.WorkspaceBackupStatuses[i] != nil {

			if swag.IsZero(m.WorkspaceBackupStatuses[i]) { // not required
				return nil
			}

			if err := m.WorkspaceBackupStatuses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspaceBackupStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workspaceBackupStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspaceBackupStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceBackupStatus) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceBackupStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
