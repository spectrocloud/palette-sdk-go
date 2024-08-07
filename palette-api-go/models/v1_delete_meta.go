// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1DeleteMeta Properties to send back after deletion operation
//
// swagger:model v1DeleteMeta
type V1DeleteMeta struct {

	// count
	Count int64 `json:"count,omitempty"`

	// items
	Items map[string]string `json:"items,omitempty"`
}

// Validate validates this v1 delete meta
func (m *V1DeleteMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DeleteMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DeleteMeta) UnmarshalBinary(b []byte) error {
	var res V1DeleteMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
