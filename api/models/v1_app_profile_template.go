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

// V1AppProfileTemplate Application profile template information
//
// swagger:model v1AppProfileTemplate
type V1AppProfileTemplate struct {

	// Application profile tiers
	// Unique: true
	AppTiers []*V1AppTierRef `json:"appTiers"`

	// Application profile registries reference
	RegistryRefs []*V1ObjectReference `json:"registryRefs"`
}

// Validate validates this v1 app profile template
func (m *V1AppProfileTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppTiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryRefs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppProfileTemplate) validateAppTiers(formats strfmt.Registry) error {

	if swag.IsZero(m.AppTiers) { // not required
		return nil
	}

	if err := validate.UniqueItems("appTiers", "body", m.AppTiers); err != nil {
		return err
	}

	for i := 0; i < len(m.AppTiers); i++ {
		if swag.IsZero(m.AppTiers[i]) { // not required
			continue
		}

		if m.AppTiers[i] != nil {
			if err := m.AppTiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appTiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AppProfileTemplate) validateRegistryRefs(formats strfmt.Registry) error {

	if swag.IsZero(m.RegistryRefs) { // not required
		return nil
	}

	for i := 0; i < len(m.RegistryRefs); i++ {
		if swag.IsZero(m.RegistryRefs[i]) { // not required
			continue
		}

		if m.RegistryRefs[i] != nil {
			if err := m.RegistryRefs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registryRefs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppProfileTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppProfileTemplate) UnmarshalBinary(b []byte) error {
	var res V1AppProfileTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}