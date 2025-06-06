// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UserKubectlSession v1 user kubectl session
//
// swagger:model v1UserKubectlSession
type V1UserKubectlSession struct {

	// cluster Uid
	ClusterUID string `json:"clusterUid,omitempty"`

	// creation time
	CreationTime string `json:"creationTime,omitempty"`

	// is active
	IsActive bool `json:"isActive,omitempty"`

	// pod Ip
	PodIP string `json:"podIp,omitempty"`

	// pod name
	PodName string `json:"podName,omitempty"`

	// port
	Port string `json:"port,omitempty"`

	// project Uid
	ProjectUID string `json:"projectUid,omitempty"`

	// session Uid
	SessionUID string `json:"sessionUid,omitempty"`

	// shelly cluster
	ShellyCluster string `json:"shellyCluster,omitempty"`

	// tenant cluster endpoint
	TenantClusterEndpoint string `json:"tenantClusterEndpoint,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`

	// user Uid
	UserUID string `json:"userUid,omitempty"`
}

// Validate validates this v1 user kubectl session
func (m *V1UserKubectlSession) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 user kubectl session based on context it is used
func (m *V1UserKubectlSession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UserKubectlSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UserKubectlSession) UnmarshalBinary(b []byte) error {
	var res V1UserKubectlSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
