// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1NotificationSource Describes origin info for the notification
//
// swagger:model v1NotificationSource
type V1NotificationSource struct {

	// Describes component where notification originated
	Component string `json:"component,omitempty"`
}

// Validate validates this v1 notification source
func (m *V1NotificationSource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NotificationSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NotificationSource) UnmarshalBinary(b []byte) error {
	var res V1NotificationSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}