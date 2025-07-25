// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EcrAuthorizationRequestEntity Ecr registry credentials entity
//
// swagger:model v1EcrAuthorizationRequestEntity
type V1EcrAuthorizationRequestEntity struct {

	// aws cloud account
	AwsCloudAccount *V1AwsCloudAccount `json:"awsCloudAccount,omitempty"`

	// Endpoint url to make the request
	Endpoint string `json:"endpoint,omitempty"`

	// If it is public or private
	IsPrivate bool `json:"isPrivate,omitempty"`

	// Name of the region
	Region string `json:"region,omitempty"`
}

// Validate validates this v1 ecr authorization request entity
func (m *V1EcrAuthorizationRequestEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsCloudAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EcrAuthorizationRequestEntity) validateAwsCloudAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsCloudAccount) { // not required
		return nil
	}

	if m.AwsCloudAccount != nil {
		if err := m.AwsCloudAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsCloudAccount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EcrAuthorizationRequestEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EcrAuthorizationRequestEntity) UnmarshalBinary(b []byte) error {
	var res V1EcrAuthorizationRequestEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
