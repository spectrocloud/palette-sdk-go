// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ObjectScopeEntity Object scope identity meta
//
// swagger:model v1ObjectScopeEntity
type V1ObjectScopeEntity struct {

	// name
	Name string `json:"name,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 object scope entity
func (m *V1ObjectScopeEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ObjectScopeEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ObjectScopeEntity) UnmarshalBinary(b []byte) error {
	var res V1ObjectScopeEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}