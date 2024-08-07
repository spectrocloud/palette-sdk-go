// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterLocation Cluster location information
//
// swagger:model v1ClusterLocation
type V1ClusterLocation struct {

	// country code for cluster location
	CountryCode string `json:"countryCode,omitempty"`

	// country name for cluster location
	CountryName string `json:"countryName,omitempty"`

	// geo loc
	GeoLoc *V1GeolocationLatlong `json:"geoLoc,omitempty"`

	// region code for cluster location
	RegionCode string `json:"regionCode,omitempty"`

	// region name for cluster location
	RegionName string `json:"regionName,omitempty"`
}

// Validate validates this v1 cluster location
func (m *V1ClusterLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeoLoc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterLocation) validateGeoLoc(formats strfmt.Registry) error {

	if swag.IsZero(m.GeoLoc) { // not required
		return nil
	}

	if m.GeoLoc != nil {
		if err := m.GeoLoc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geoLoc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterLocation) UnmarshalBinary(b []byte) error {
	var res V1ClusterLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
