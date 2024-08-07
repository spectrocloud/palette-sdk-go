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

// V1EdgeHostsMeta v1 edge hosts meta
//
// swagger:model v1EdgeHostsMeta
type V1EdgeHostsMeta struct {

	// edge hosts
	EdgeHosts []*V1EdgeHostMeta `json:"edgeHosts"`
}

// Validate validates this v1 edge hosts meta
func (m *V1EdgeHostsMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeHosts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeHostsMeta) validateEdgeHosts(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeHosts) { // not required
		return nil
	}

	for i := 0; i < len(m.EdgeHosts); i++ {
		if swag.IsZero(m.EdgeHosts[i]) { // not required
			continue
		}

		if m.EdgeHosts[i] != nil {
			if err := m.EdgeHosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeHosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeHostsMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeHostsMeta) UnmarshalBinary(b []byte) error {
	var res V1EdgeHostsMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
