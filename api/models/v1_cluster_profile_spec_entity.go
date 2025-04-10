// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterProfileSpecEntity Cluster profile update spec
//
// swagger:model v1ClusterProfileSpecEntity
type V1ClusterProfileSpecEntity struct {

	// Cluster profile version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 cluster profile spec entity
func (m *V1ClusterProfileSpecEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 cluster profile spec entity based on context it is used
func (m *V1ClusterProfileSpecEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterProfileSpecEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterProfileSpecEntity) UnmarshalBinary(b []byte) error {
	var res V1ClusterProfileSpecEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
