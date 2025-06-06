// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterMetaStatusUpdates Cluster meta updates information
//
// swagger:model v1ClusterMetaStatusUpdates
type V1ClusterMetaStatusUpdates struct {

	// is updates pending
	IsUpdatesPending bool `json:"isUpdatesPending"`
}

// Validate validates this v1 cluster meta status updates
func (m *V1ClusterMetaStatusUpdates) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 cluster meta status updates based on context it is used
func (m *V1ClusterMetaStatusUpdates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterMetaStatusUpdates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterMetaStatusUpdates) UnmarshalBinary(b []byte) error {
	var res V1ClusterMetaStatusUpdates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
