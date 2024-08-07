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

// V1OpenStackRegions List of OpenStack regions and domains
//
// swagger:model v1OpenStackRegions
type V1OpenStackRegions struct {

	// List of OpenStack domains
	// Required: true
	// Unique: true
	Domains []*V1OpenStackDomain `json:"domains"`

	// List of OpenStack regions
	// Required: true
	// Unique: true
	Regions []*V1OpenStackRegion `json:"regions"`
}

// Validate validates this v1 open stack regions
func (m *V1OpenStackRegions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OpenStackRegions) validateDomains(formats strfmt.Registry) error {

	if err := validate.Required("domains", "body", m.Domains); err != nil {
		return err
	}

	if err := validate.UniqueItems("domains", "body", m.Domains); err != nil {
		return err
	}

	for i := 0; i < len(m.Domains); i++ {
		if swag.IsZero(m.Domains[i]) { // not required
			continue
		}

		if m.Domains[i] != nil {
			if err := m.Domains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1OpenStackRegions) validateRegions(formats strfmt.Registry) error {

	if err := validate.Required("regions", "body", m.Regions); err != nil {
		return err
	}

	if err := validate.UniqueItems("regions", "body", m.Regions); err != nil {
		return err
	}

	for i := 0; i < len(m.Regions); i++ {
		if swag.IsZero(m.Regions[i]) { // not required
			continue
		}

		if m.Regions[i] != nil {
			if err := m.Regions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OpenStackRegions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OpenStackRegions) UnmarshalBinary(b []byte) error {
	var res V1OpenStackRegions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
