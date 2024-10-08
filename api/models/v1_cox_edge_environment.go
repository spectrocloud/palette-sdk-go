// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CoxEdgeEnvironment CoxEdge environment entity
//
// swagger:model v1CoxEdgeEnvironment
type V1CoxEdgeEnvironment struct {

	// CoxEdge environment id
	ID string `json:"id,omitempty"`

	// CoxEdge environment state
	IsDeleted bool `json:"isDeleted,omitempty"`

	// CoxEdge environment name
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 cox edge environment
func (m *V1CoxEdgeEnvironment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CoxEdgeEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CoxEdgeEnvironment) UnmarshalBinary(b []byte) error {
	var res V1CoxEdgeEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
