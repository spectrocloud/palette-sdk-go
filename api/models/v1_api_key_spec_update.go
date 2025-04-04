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

// V1APIKeySpecUpdate API key update request specification
//
// swagger:model v1ApiKeySpecUpdate
type V1APIKeySpecUpdate struct {

	// API key expiry date
	// Format: date-time
	Expiry V1Time `json:"expiry,omitempty"`
}

// Validate validates this v1 Api key spec update
func (m *V1APIKeySpecUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1APIKeySpecUpdate) validateExpiry(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiry) { // not required
		return nil
	}

	if err := m.Expiry.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("expiry")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("expiry")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 Api key spec update based on the context it is used
func (m *V1APIKeySpecUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExpiry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1APIKeySpecUpdate) contextValidateExpiry(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Expiry) { // not required
		return nil
	}

	if err := m.Expiry.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("expiry")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("expiry")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1APIKeySpecUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1APIKeySpecUpdate) UnmarshalBinary(b []byte) error {
	var res V1APIKeySpecUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
