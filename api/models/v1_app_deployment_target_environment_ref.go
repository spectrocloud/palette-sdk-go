// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AppDeploymentTargetEnvironmentRef Application deployment target environment reference
//
// swagger:model v1AppDeploymentTargetEnvironmentRef
type V1AppDeploymentTargetEnvironmentRef struct {

	// Application deployment target resource name
	Name string `json:"name,omitempty"`

	// Application deployment target resource type [ "nestedCluster", "clusterGroup" ]
	Type string `json:"type,omitempty"`

	// Application deployment target resource uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 app deployment target environment ref
func (m *V1AppDeploymentTargetEnvironmentRef) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AppDeploymentTargetEnvironmentRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppDeploymentTargetEnvironmentRef) UnmarshalBinary(b []byte) error {
	var res V1AppDeploymentTargetEnvironmentRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}