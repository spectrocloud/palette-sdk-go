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
	"github.com/go-openapi/validate"
)

// V1OpenStackAzs List of OpenStack azs
//
// swagger:model v1OpenStackAzs
type V1OpenStackAzs struct {

	// azs
	// Required: true
	// Unique: true
	Azs []*V1OpenStackAz `json:"azs"`
}

// Validate validates this v1 open stack azs
func (m *V1OpenStackAzs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OpenStackAzs) validateAzs(formats strfmt.Registry) error {

	if err := validate.Required("azs", "body", m.Azs); err != nil {
		return err
	}

	if err := validate.UniqueItems("azs", "body", m.Azs); err != nil {
		return err
	}

	for i := 0; i < len(m.Azs); i++ {
		if swag.IsZero(m.Azs[i]) { // not required
			continue
		}

		if m.Azs[i] != nil {
			if err := m.Azs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 open stack azs based on the context it is used
func (m *V1OpenStackAzs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAzs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OpenStackAzs) contextValidateAzs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Azs); i++ {

		if m.Azs[i] != nil {

			if swag.IsZero(m.Azs[i]) { // not required
				return nil
			}

			if err := m.Azs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OpenStackAzs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OpenStackAzs) UnmarshalBinary(b []byte) error {
	var res V1OpenStackAzs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
