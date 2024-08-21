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

// V1CertificateAuthority Certificate Authority
//
// swagger:model v1CertificateAuthority
type V1CertificateAuthority struct {

	// certificates
	Certificates []*V1Certificate `json:"certificates"`

	// Certificate expiry time
	// Format: date-time
	Expiry V1Time `json:"expiry,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 certificate authority
func (m *V1CertificateAuthority) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CertificateAuthority) validateCertificates(formats strfmt.Registry) error {

	if swag.IsZero(m.Certificates) { // not required
		return nil
	}

	for i := 0; i < len(m.Certificates); i++ {
		if swag.IsZero(m.Certificates[i]) { // not required
			continue
		}

		if m.Certificates[i] != nil {
			if err := m.Certificates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CertificateAuthority) validateExpiry(formats strfmt.Registry) error {

	if swag.IsZero(m.Expiry) { // not required
		return nil
	}

	if err := m.Expiry.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("expiry")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CertificateAuthority) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CertificateAuthority) UnmarshalBinary(b []byte) error {
	var res V1CertificateAuthority
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}