// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1APIKeyStatus API key status
//
// swagger:model v1ApiKeyStatus
type V1APIKeyStatus struct {

	// API key active state
	IsActive bool `json:"isActive,omitempty"`
}

// Validate validates this v1 Api key status
func (m *V1APIKeyStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 Api key status based on context it is used
func (m *V1APIKeyStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1APIKeyStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1APIKeyStatus) UnmarshalBinary(b []byte) error {
	var res V1APIKeyStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
