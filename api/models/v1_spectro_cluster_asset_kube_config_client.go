// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroClusterAssetKubeConfigClient Cluster asset Kube Config Client
//
// swagger:model v1SpectroClusterAssetKubeConfigClient
type V1SpectroClusterAssetKubeConfigClient struct {

	// kubeconfigclient
	Kubeconfigclient string `json:"kubeconfigclient,omitempty"`
}

// Validate validates this v1 spectro cluster asset kube config client
func (m *V1SpectroClusterAssetKubeConfigClient) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 spectro cluster asset kube config client based on context it is used
func (m *V1SpectroClusterAssetKubeConfigClient) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterAssetKubeConfigClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterAssetKubeConfigClient) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterAssetKubeConfigClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
