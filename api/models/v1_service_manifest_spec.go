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

// V1ServiceManifestSpec v1 service manifest spec
//
// swagger:model v1ServiceManifestSpec
type V1ServiceManifestSpec struct {

	// manifests
	// Unique: true
	Manifests []*V1GitRepoFileContent `json:"manifests"`

	// name
	Name string `json:"name,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 service manifest spec
func (m *V1ServiceManifestSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManifests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ServiceManifestSpec) validateManifests(formats strfmt.Registry) error {

	if swag.IsZero(m.Manifests) { // not required
		return nil
	}

	if err := validate.UniqueItems("manifests", "body", m.Manifests); err != nil {
		return err
	}

	for i := 0; i < len(m.Manifests); i++ {
		if swag.IsZero(m.Manifests[i]) { // not required
			continue
		}

		if m.Manifests[i] != nil {
			if err := m.Manifests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("manifests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ServiceManifestSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ServiceManifestSpec) UnmarshalBinary(b []byte) error {
	var res V1ServiceManifestSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
