// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GcpNetworks List of GCP networks
//
// swagger:model v1GcpNetworks
type V1GcpNetworks struct {

	// networks
	Networks []*V1GcpNetwork `json:"networks"`
}

// Validate validates this v1 gcp networks
func (m *V1GcpNetworks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GcpNetworks) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 gcp networks based on the context it is used
func (m *V1GcpNetworks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GcpNetworks) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Networks); i++ {

		if m.Networks[i] != nil {

			if swag.IsZero(m.Networks[i]) { // not required
				return nil
			}

			if err := m.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GcpNetworks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GcpNetworks) UnmarshalBinary(b []byte) error {
	var res V1GcpNetworks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
