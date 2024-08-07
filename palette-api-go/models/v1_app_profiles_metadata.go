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

// V1AppProfilesMetadata v1 app profiles metadata
//
// swagger:model v1AppProfilesMetadata
type V1AppProfilesMetadata struct {

	// app profiles
	// Unique: true
	AppProfiles []*V1AppProfileMetadata `json:"appProfiles"`
}

// Validate validates this v1 app profiles metadata
func (m *V1AppProfilesMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppProfilesMetadata) validateAppProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.AppProfiles) { // not required
		return nil
	}

	if err := validate.UniqueItems("appProfiles", "body", m.AppProfiles); err != nil {
		return err
	}

	for i := 0; i < len(m.AppProfiles); i++ {
		if swag.IsZero(m.AppProfiles[i]) { // not required
			continue
		}

		if m.AppProfiles[i] != nil {
			if err := m.AppProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppProfilesMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppProfilesMetadata) UnmarshalBinary(b []byte) error {
	var res V1AppProfilesMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
