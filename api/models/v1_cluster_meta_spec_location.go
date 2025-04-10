// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterMetaSpecLocation Cluster location information
//
// swagger:model v1ClusterMetaSpecLocation
type V1ClusterMetaSpecLocation struct {

	// coordinates
	Coordinates []float64 `json:"coordinates"`

	// country code
	CountryCode string `json:"countryCode,omitempty"`

	// country name
	CountryName string `json:"countryName,omitempty"`

	// region code
	RegionCode string `json:"regionCode,omitempty"`

	// region name
	RegionName string `json:"regionName,omitempty"`
}

// Validate validates this v1 cluster meta spec location
func (m *V1ClusterMetaSpecLocation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 cluster meta spec location based on context it is used
func (m *V1ClusterMetaSpecLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterMetaSpecLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterMetaSpecLocation) UnmarshalBinary(b []byte) error {
	var res V1ClusterMetaSpecLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
