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

// V1DashboardWorkspaceStatus Workspace status
//
// swagger:model v1DashboardWorkspaceStatus
type V1DashboardWorkspaceStatus struct {

	// namespaces
	// Unique: true
	Namespaces []*V1DashboardWorkspaceNamespaceAllocation `json:"namespaces"`

	// total
	Total *V1DashboardWorkspaceAllocation `json:"total,omitempty"`
}

// Validate validates this v1 dashboard workspace status
func (m *V1DashboardWorkspaceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DashboardWorkspaceStatus) validateNamespaces(formats strfmt.Registry) error {

	if swag.IsZero(m.Namespaces) { // not required
		return nil
	}

	if err := validate.UniqueItems("namespaces", "body", m.Namespaces); err != nil {
		return err
	}

	for i := 0; i < len(m.Namespaces); i++ {
		if swag.IsZero(m.Namespaces[i]) { // not required
			continue
		}

		if m.Namespaces[i] != nil {
			if err := m.Namespaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("namespaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1DashboardWorkspaceStatus) validateTotal(formats strfmt.Registry) error {

	if swag.IsZero(m.Total) { // not required
		return nil
	}

	if m.Total != nil {
		if err := m.Total.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1DashboardWorkspaceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DashboardWorkspaceStatus) UnmarshalBinary(b []byte) error {
	var res V1DashboardWorkspaceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
