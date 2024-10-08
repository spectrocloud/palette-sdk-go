// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CoxEdgePort CoxEdge network port
//
// swagger:model v1CoxEdgePort
type V1CoxEdgePort struct {

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// public port
	PublicPort string `json:"publicPort,omitempty"`

	// public port desc
	PublicPortDesc string `json:"publicPortDesc,omitempty"`
}

// Validate validates this v1 cox edge port
func (m *V1CoxEdgePort) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CoxEdgePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CoxEdgePort) UnmarshalBinary(b []byte) error {
	var res V1CoxEdgePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
