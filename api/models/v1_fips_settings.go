// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1FipsSettings FIPS configuration
//
// swagger:model v1FipsSettings
type V1FipsSettings struct {

	// fips cluster feature config
	FipsClusterFeatureConfig *V1NonFipsConfig `json:"fipsClusterFeatureConfig,omitempty"`

	// fips cluster import config
	FipsClusterImportConfig *V1NonFipsConfig `json:"fipsClusterImportConfig,omitempty"`

	// fips pack config
	FipsPackConfig *V1NonFipsConfig `json:"fipsPackConfig,omitempty"`
}

// Validate validates this v1 fips settings
func (m *V1FipsSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFipsClusterFeatureConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFipsClusterImportConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFipsPackConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1FipsSettings) validateFipsClusterFeatureConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.FipsClusterFeatureConfig) { // not required
		return nil
	}

	if m.FipsClusterFeatureConfig != nil {
		if err := m.FipsClusterFeatureConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fipsClusterFeatureConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1FipsSettings) validateFipsClusterImportConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.FipsClusterImportConfig) { // not required
		return nil
	}

	if m.FipsClusterImportConfig != nil {
		if err := m.FipsClusterImportConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fipsClusterImportConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1FipsSettings) validateFipsPackConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.FipsPackConfig) { // not required
		return nil
	}

	if m.FipsPackConfig != nil {
		if err := m.FipsPackConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fipsPackConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1FipsSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FipsSettings) UnmarshalBinary(b []byte) error {
	var res V1FipsSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
