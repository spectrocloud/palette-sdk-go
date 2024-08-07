// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AppDeploymentClusterHealth Application deployment cluster health status
//
// swagger:model v1AppDeploymentClusterHealth
type V1AppDeploymentClusterHealth struct {

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 app deployment cluster health
func (m *V1AppDeploymentClusterHealth) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AppDeploymentClusterHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppDeploymentClusterHealth) UnmarshalBinary(b []byte) error {
	var res V1AppDeploymentClusterHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
