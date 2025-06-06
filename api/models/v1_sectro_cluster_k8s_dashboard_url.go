// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SectroClusterK8sDashboardURL Service version information
//
// swagger:model v1SectroClusterK8sDashboardUrl
type V1SectroClusterK8sDashboardURL struct {

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this v1 sectro cluster k8s dashboard Url
func (m *V1SectroClusterK8sDashboardURL) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 sectro cluster k8s dashboard Url based on context it is used
func (m *V1SectroClusterK8sDashboardURL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SectroClusterK8sDashboardURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SectroClusterK8sDashboardURL) UnmarshalBinary(b []byte) error {
	var res V1SectroClusterK8sDashboardURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
