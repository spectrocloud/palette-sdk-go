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

// V1UserRolesEntity v1 user roles entity
//
// swagger:model v1UserRolesEntity
type V1UserRolesEntity struct {

	// inherited roles
	InheritedRoles []*V1UIDSummary `json:"inheritedRoles"`

	// roles
	Roles []*V1UIDSummary `json:"roles"`
}

// Validate validates this v1 user roles entity
func (m *V1UserRolesEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInheritedRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UserRolesEntity) validateInheritedRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.InheritedRoles) { // not required
		return nil
	}

	for i := 0; i < len(m.InheritedRoles); i++ {
		if swag.IsZero(m.InheritedRoles[i]) { // not required
			continue
		}

		if m.InheritedRoles[i] != nil {
			if err := m.InheritedRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inheritedRoles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inheritedRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1UserRolesEntity) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 user roles entity based on the context it is used
func (m *V1UserRolesEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInheritedRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UserRolesEntity) contextValidateInheritedRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InheritedRoles); i++ {

		if m.InheritedRoles[i] != nil {

			if swag.IsZero(m.InheritedRoles[i]) { // not required
				return nil
			}

			if err := m.InheritedRoles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inheritedRoles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inheritedRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1UserRolesEntity) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Roles); i++ {

		if m.Roles[i] != nil {

			if swag.IsZero(m.Roles[i]) { // not required
				return nil
			}

			if err := m.Roles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1UserRolesEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UserRolesEntity) UnmarshalBinary(b []byte) error {
	var res V1UserRolesEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
