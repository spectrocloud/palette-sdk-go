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

// V1MachineCertificate K8 Certificates for control plane nodes
//
// swagger:model v1MachineCertificate
type V1MachineCertificate struct {

	// Applicable certificate authorities
	CertificateAuthorities []*V1CertificateAuthority `json:"certificateAuthorities"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 machine certificate
func (m *V1MachineCertificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateAuthorities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineCertificate) validateCertificateAuthorities(formats strfmt.Registry) error {

	if swag.IsZero(m.CertificateAuthorities) { // not required
		return nil
	}

	for i := 0; i < len(m.CertificateAuthorities); i++ {
		if swag.IsZero(m.CertificateAuthorities[i]) { // not required
			continue
		}

		if m.CertificateAuthorities[i] != nil {
			if err := m.CertificateAuthorities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificateAuthorities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineCertificate) UnmarshalBinary(b []byte) error {
	var res V1MachineCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
