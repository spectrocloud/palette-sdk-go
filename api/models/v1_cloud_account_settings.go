// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CloudAccountSettings Cloud account settings
//
// swagger:model v1CloudAccountSettings
type V1CloudAccountSettings struct {

	// Will disable certain properties request to cloud and the input is collected directly from the user
	DisablePropertiesRequest bool `json:"disablePropertiesRequest"`
}

// Validate validates this v1 cloud account settings
func (m *V1CloudAccountSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 cloud account settings based on context it is used
func (m *V1CloudAccountSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CloudAccountSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CloudAccountSettings) UnmarshalBinary(b []byte) error {
	var res V1CloudAccountSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
