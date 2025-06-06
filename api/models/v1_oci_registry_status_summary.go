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

// V1OciRegistryStatusSummary OCI registry status summary
//
// swagger:model v1OciRegistryStatusSummary
type V1OciRegistryStatusSummary struct {

	// sync
	Sync *V1RegistrySyncStatus `json:"sync,omitempty"`
}

// Validate validates this v1 oci registry status summary
func (m *V1OciRegistryStatusSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSync(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OciRegistryStatusSummary) validateSync(formats strfmt.Registry) error {
	if swag.IsZero(m.Sync) { // not required
		return nil
	}

	if m.Sync != nil {
		if err := m.Sync.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 oci registry status summary based on the context it is used
func (m *V1OciRegistryStatusSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSync(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OciRegistryStatusSummary) contextValidateSync(ctx context.Context, formats strfmt.Registry) error {

	if m.Sync != nil {

		if swag.IsZero(m.Sync) { // not required
			return nil
		}

		if err := m.Sync.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OciRegistryStatusSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OciRegistryStatusSummary) UnmarshalBinary(b []byte) error {
	var res V1OciRegistryStatusSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
