// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1AuthTokenSettings System auth token settings
//
// swagger:model v1AuthTokenSettings
type V1AuthTokenSettings struct {

	// Auth token expiry time in minutes
	// Maximum: 1440
	// Minimum: 15
	ExpiryTimeMinutes int32 `json:"expiryTimeMinutes"`
}

// Validate validates this v1 auth token settings
func (m *V1AuthTokenSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiryTimeMinutes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AuthTokenSettings) validateExpiryTimeMinutes(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiryTimeMinutes) { // not required
		return nil
	}

	if err := validate.MinimumInt("expiryTimeMinutes", "body", int64(m.ExpiryTimeMinutes), 15, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("expiryTimeMinutes", "body", int64(m.ExpiryTimeMinutes), 1440, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AuthTokenSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AuthTokenSettings) UnmarshalBinary(b []byte) error {
	var res V1AuthTokenSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}