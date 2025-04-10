// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ProjectAlertComponent Project alert component
//
// swagger:model v1ProjectAlertComponent
type V1ProjectAlertComponent struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// supported channels
	SupportedChannels []string `json:"supportedChannels"`
}

// Validate validates this v1 project alert component
func (m *V1ProjectAlertComponent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 project alert component based on context it is used
func (m *V1ProjectAlertComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectAlertComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectAlertComponent) UnmarshalBinary(b []byte) error {
	var res V1ProjectAlertComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
