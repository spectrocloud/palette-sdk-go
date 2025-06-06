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

// V1SpcPatchTimeEntity v1 spc patch time entity
//
// swagger:model v1SpcPatchTimeEntity
type V1SpcPatchTimeEntity struct {

	// cluster hash
	ClusterHash string `json:"clusterHash,omitempty"`

	// patch time
	// Format: date-time
	PatchTime V1Time `json:"patchTime,omitempty"`
}

// Validate validates this v1 spc patch time entity
func (m *V1SpcPatchTimeEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePatchTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpcPatchTimeEntity) validatePatchTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PatchTime) { // not required
		return nil
	}

	if err := m.PatchTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("patchTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("patchTime")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 spc patch time entity based on the context it is used
func (m *V1SpcPatchTimeEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePatchTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpcPatchTimeEntity) contextValidatePatchTime(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.PatchTime) { // not required
		return nil
	}

	if err := m.PatchTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("patchTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("patchTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpcPatchTimeEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpcPatchTimeEntity) UnmarshalBinary(b []byte) error {
	var res V1SpcPatchTimeEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
