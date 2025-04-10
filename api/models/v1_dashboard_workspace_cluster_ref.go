// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1DashboardWorkspaceClusterRef Workspace cluster reference
//
// swagger:model v1DashboardWorkspaceClusterRef
type V1DashboardWorkspaceClusterRef struct {

	// name
	Name string `json:"name,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 dashboard workspace cluster ref
func (m *V1DashboardWorkspaceClusterRef) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 dashboard workspace cluster ref based on context it is used
func (m *V1DashboardWorkspaceClusterRef) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DashboardWorkspaceClusterRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DashboardWorkspaceClusterRef) UnmarshalBinary(b []byte) error {
	var res V1DashboardWorkspaceClusterRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
