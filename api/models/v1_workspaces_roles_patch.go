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

// V1WorkspacesRolesPatch v1 workspaces roles patch
//
// swagger:model v1WorkspacesRolesPatch
type V1WorkspacesRolesPatch struct {

	// workspaces
	Workspaces []*V1WorkspaceRolesPatch `json:"workspaces"`
}

// Validate validates this v1 workspaces roles patch
func (m *V1WorkspacesRolesPatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkspaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspacesRolesPatch) validateWorkspaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Workspaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Workspaces); i++ {
		if swag.IsZero(m.Workspaces[i]) { // not required
			continue
		}

		if m.Workspaces[i] != nil {
			if err := m.Workspaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workspaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 workspaces roles patch based on the context it is used
func (m *V1WorkspacesRolesPatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkspaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspacesRolesPatch) contextValidateWorkspaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Workspaces); i++ {

		if m.Workspaces[i] != nil {

			if swag.IsZero(m.Workspaces[i]) { // not required
				return nil
			}

			if err := m.Workspaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workspaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspacesRolesPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspacesRolesPatch) UnmarshalBinary(b []byte) error {
	var res V1WorkspacesRolesPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
