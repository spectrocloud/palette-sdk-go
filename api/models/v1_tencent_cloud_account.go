// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1TencentCloudAccount v1 tencent cloud account
//
// swagger:model v1TencentCloudAccount
type V1TencentCloudAccount struct {

	// Tencent api secretID
	// Required: true
	SecretID *string `json:"secretId"`

	// Tencent api secret key
	// Required: true
	SecretKey *string `json:"secretKey"`
}

// Validate validates this v1 tencent cloud account
func (m *V1TencentCloudAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecretID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TencentCloudAccount) validateSecretID(formats strfmt.Registry) error {

	if err := validate.Required("secretId", "body", m.SecretID); err != nil {
		return err
	}

	return nil
}

func (m *V1TencentCloudAccount) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("secretKey", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TencentCloudAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TencentCloudAccount) UnmarshalBinary(b []byte) error {
	var res V1TencentCloudAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}