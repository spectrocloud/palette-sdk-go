// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroClusterNotifications Spectro cluster notifications
//
// swagger:model v1SpectroClusterNotifications
type V1SpectroClusterNotifications struct {

	// is available
	IsAvailable bool `json:"isAvailable"`
}

// Validate validates this v1 spectro cluster notifications
func (m *V1SpectroClusterNotifications) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterNotifications) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterNotifications) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterNotifications
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}