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

// V1TencentAccounts v1 tencent accounts
//
// swagger:model v1TencentAccounts
type V1TencentAccounts struct {

	// items
	// Required: true
	// Unique: true
	Items []*V1TencentAccount `json:"items"`

	// listmeta
	Listmeta *V1ListMetaData `json:"listmeta,omitempty"`
}

// Validate validates this v1 tencent accounts
func (m *V1TencentAccounts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListmeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TencentAccounts) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	if err := validate.UniqueItems("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1TencentAccounts) validateListmeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Listmeta) { // not required
		return nil
	}

	if m.Listmeta != nil {
		if err := m.Listmeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listmeta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TencentAccounts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TencentAccounts) UnmarshalBinary(b []byte) error {
	var res V1TencentAccounts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
