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

// V1EksFargateProfiles Fargate profiles
//
// swagger:model v1EksFargateProfiles
type V1EksFargateProfiles struct {

	// fargate profiles
	FargateProfiles []*V1FargateProfile `json:"fargateProfiles"`
}

// Validate validates this v1 eks fargate profiles
func (m *V1EksFargateProfiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFargateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EksFargateProfiles) validateFargateProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.FargateProfiles) { // not required
		return nil
	}

	for i := 0; i < len(m.FargateProfiles); i++ {
		if swag.IsZero(m.FargateProfiles[i]) { // not required
			continue
		}

		if m.FargateProfiles[i] != nil {
			if err := m.FargateProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fargateProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EksFargateProfiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EksFargateProfiles) UnmarshalBinary(b []byte) error {
	var res V1EksFargateProfiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
