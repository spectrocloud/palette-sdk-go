// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterFeatureActor Compliance Scan actor
//
// swagger:model v1ClusterFeatureActor
type V1ClusterFeatureActor struct {

	// actor type
	ActorType string `json:"actorType,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 cluster feature actor
func (m *V1ClusterFeatureActor) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 cluster feature actor based on context it is used
func (m *V1ClusterFeatureActor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterFeatureActor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterFeatureActor) UnmarshalBinary(b []byte) error {
	var res V1ClusterFeatureActor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
