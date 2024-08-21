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

// V1AppDeploymentProfileEntity Application deployment profile request payload
//
// swagger:model v1AppDeploymentProfileEntity
type V1AppDeploymentProfileEntity struct {

	// Application deployment profile uid
	// Required: true
	AppProfileUID *string `json:"appProfileUid"`
}

// Validate validates this v1 app deployment profile entity
func (m *V1AppDeploymentProfileEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppProfileUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppDeploymentProfileEntity) validateAppProfileUID(formats strfmt.Registry) error {

	if err := validate.Required("appProfileUid", "body", m.AppProfileUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppDeploymentProfileEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppDeploymentProfileEntity) UnmarshalBinary(b []byte) error {
	var res V1AppDeploymentProfileEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}