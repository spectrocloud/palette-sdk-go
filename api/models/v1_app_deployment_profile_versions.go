// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AppDeploymentProfileVersions Application deployment profile versions
//
// swagger:model v1AppDeploymentProfileVersions
type V1AppDeploymentProfileVersions struct {

	// Application deployment profile available versions
	AvailableVersions []*V1AppDeploymentProfileVersion `json:"availableVersions"`

	// Application deployment profile latest versions
	LatestVersions []*V1AppDeploymentProfileVersion `json:"latestVersions"`

	// metadata
	Metadata *V1AppDeploymentProfileMeta `json:"metadata,omitempty"`
}

// Validate validates this v1 app deployment profile versions
func (m *V1AppDeploymentProfileVersions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppDeploymentProfileVersions) validateAvailableVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailableVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailableVersions); i++ {
		if swag.IsZero(m.AvailableVersions[i]) { // not required
			continue
		}

		if m.AvailableVersions[i] != nil {
			if err := m.AvailableVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availableVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AppDeploymentProfileVersions) validateLatestVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.LatestVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.LatestVersions); i++ {
		if swag.IsZero(m.LatestVersions[i]) { // not required
			continue
		}

		if m.LatestVersions[i] != nil {
			if err := m.LatestVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("latestVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AppDeploymentProfileVersions) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppDeploymentProfileVersions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppDeploymentProfileVersions) UnmarshalBinary(b []byte) error {
	var res V1AppDeploymentProfileVersions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}