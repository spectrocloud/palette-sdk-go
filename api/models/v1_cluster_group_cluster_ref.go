// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterGroupClusterRef Cluster group cluster reference
//
// swagger:model v1ClusterGroupClusterRef
type V1ClusterGroupClusterRef struct {

	// cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// cluster Uid
	ClusterUID string `json:"clusterUid,omitempty"`
}

// Validate validates this v1 cluster group cluster ref
func (m *V1ClusterGroupClusterRef) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 cluster group cluster ref based on context it is used
func (m *V1ClusterGroupClusterRef) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterGroupClusterRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterGroupClusterRef) UnmarshalBinary(b []byte) error {
	var res V1ClusterGroupClusterRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
