// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterHelmRegistry Cluster helm registry information
//
// swagger:model v1ClusterHelmRegistry
type V1ClusterHelmRegistry struct {

	// name
	Name string `json:"name,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 cluster helm registry
func (m *V1ClusterHelmRegistry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 cluster helm registry based on context it is used
func (m *V1ClusterHelmRegistry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterHelmRegistry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterHelmRegistry) UnmarshalBinary(b []byte) error {
	var res V1ClusterHelmRegistry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
