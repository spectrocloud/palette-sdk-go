// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1BulkDeleteFailure v1 bulk delete failure
//
// swagger:model v1BulkDeleteFailure
type V1BulkDeleteFailure struct {

	// err msg
	ErrMsg string `json:"errMsg,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 bulk delete failure
func (m *V1BulkDeleteFailure) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1BulkDeleteFailure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BulkDeleteFailure) UnmarshalBinary(b []byte) error {
	var res V1BulkDeleteFailure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}