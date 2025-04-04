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

// V1AzurePrivateDNSZones List of Azure storage accounts
//
// swagger:model v1AzurePrivateDnsZones
type V1AzurePrivateDNSZones struct {

	// private Dns zones
	PrivateDNSZones []*V1AzurePrivateDNSZone `json:"privateDnsZones"`
}

// Validate validates this v1 azure private Dns zones
func (m *V1AzurePrivateDNSZones) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateDNSZones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AzurePrivateDNSZones) validatePrivateDNSZones(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateDNSZones) { // not required
		return nil
	}

	for i := 0; i < len(m.PrivateDNSZones); i++ {
		if swag.IsZero(m.PrivateDNSZones[i]) { // not required
			continue
		}

		if m.PrivateDNSZones[i] != nil {
			if err := m.PrivateDNSZones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privateDnsZones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privateDnsZones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 azure private Dns zones based on the context it is used
func (m *V1AzurePrivateDNSZones) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrivateDNSZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AzurePrivateDNSZones) contextValidatePrivateDNSZones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PrivateDNSZones); i++ {

		if m.PrivateDNSZones[i] != nil {

			if swag.IsZero(m.PrivateDNSZones[i]) { // not required
				return nil
			}

			if err := m.PrivateDNSZones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privateDnsZones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privateDnsZones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AzurePrivateDNSZones) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AzurePrivateDNSZones) UnmarshalBinary(b []byte) error {
	var res V1AzurePrivateDNSZones
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
