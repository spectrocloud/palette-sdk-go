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

// V1ClusterVirtualPacksValues Virtual cluster packs values
//
// swagger:model v1ClusterVirtualPacksValues
type V1ClusterVirtualPacksValues struct {

	// packs
	Packs []*V1ClusterVirtualPacksValue `json:"packs"`
}

// Validate validates this v1 cluster virtual packs values
func (m *V1ClusterVirtualPacksValues) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePacks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterVirtualPacksValues) validatePacks(formats strfmt.Registry) error {

	if swag.IsZero(m.Packs) { // not required
		return nil
	}

	for i := 0; i < len(m.Packs); i++ {
		if swag.IsZero(m.Packs[i]) { // not required
			continue
		}

		if m.Packs[i] != nil {
			if err := m.Packs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterVirtualPacksValues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterVirtualPacksValues) UnmarshalBinary(b []byte) error {
	var res V1ClusterVirtualPacksValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
