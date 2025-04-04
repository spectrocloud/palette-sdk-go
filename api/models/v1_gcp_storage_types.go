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

// V1GcpStorageTypes List of GCP storage types
//
// swagger:model v1GcpStorageTypes
type V1GcpStorageTypes struct {

	// storage types
	StorageTypes []*V1StorageType `json:"storageTypes"`
}

// Validate validates this v1 gcp storage types
func (m *V1GcpStorageTypes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorageTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GcpStorageTypes) validateStorageTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageTypes); i++ {
		if swag.IsZero(m.StorageTypes[i]) { // not required
			continue
		}

		if m.StorageTypes[i] != nil {
			if err := m.StorageTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storageTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storageTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 gcp storage types based on the context it is used
func (m *V1GcpStorageTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GcpStorageTypes) contextValidateStorageTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StorageTypes); i++ {

		if m.StorageTypes[i] != nil {

			if swag.IsZero(m.StorageTypes[i]) { // not required
				return nil
			}

			if err := m.StorageTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storageTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storageTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GcpStorageTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GcpStorageTypes) UnmarshalBinary(b []byte) error {
	var res V1GcpStorageTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
