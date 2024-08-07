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

// V1TencentAvailabilityZones List of Tencent Availability zones
//
// swagger:model v1TencentAvailabilityZones
type V1TencentAvailabilityZones struct {

	// zones
	// Required: true
	Zones []*V1TencentAvailabilityZone `json:"zones"`
}

// Validate validates this v1 tencent availability zones
func (m *V1TencentAvailabilityZones) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateZones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TencentAvailabilityZones) validateZones(formats strfmt.Registry) error {

	if err := validate.Required("zones", "body", m.Zones); err != nil {
		return err
	}

	for i := 0; i < len(m.Zones); i++ {
		if swag.IsZero(m.Zones[i]) { // not required
			continue
		}

		if m.Zones[i] != nil {
			if err := m.Zones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TencentAvailabilityZones) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TencentAvailabilityZones) UnmarshalBinary(b []byte) error {
	var res V1TencentAvailabilityZones
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
