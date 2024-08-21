// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SpectroClusterProfilesPacksManifests v1 spectro cluster profiles packs manifests
//
// swagger:model v1SpectroClusterProfilesPacksManifests
type V1SpectroClusterProfilesPacksManifests struct {

	// profiles
	// Required: true
	Profiles []*V1ClusterProfilePacksManifests `json:"profiles"`
}

// Validate validates this v1 spectro cluster profiles packs manifests
func (m *V1SpectroClusterProfilesPacksManifests) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterProfilesPacksManifests) validateProfiles(formats strfmt.Registry) error {

	if err := validate.Required("profiles", "body", m.Profiles); err != nil {
		return err
	}

	for i := 0; i < len(m.Profiles); i++ {
		if swag.IsZero(m.Profiles[i]) { // not required
			continue
		}

		if m.Profiles[i] != nil {
			if err := m.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterProfilesPacksManifests) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterProfilesPacksManifests) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterProfilesPacksManifests
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}