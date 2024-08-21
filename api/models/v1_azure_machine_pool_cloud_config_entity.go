// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AzureMachinePoolCloudConfigEntity v1 azure machine pool cloud config entity
//
// swagger:model v1AzureMachinePoolCloudConfigEntity
type V1AzureMachinePoolCloudConfigEntity struct {

	// azs
	Azs []string `json:"azs"`

	// Instance type stands for VMSize in Azure
	InstanceType string `json:"instanceType,omitempty"`

	// whether this pool is for system node Pool
	IsSystemNodePool bool `json:"isSystemNodePool,omitempty"`

	// os disk
	OsDisk *V1AzureOSDisk `json:"osDisk,omitempty"`
}

// Validate validates this v1 azure machine pool cloud config entity
func (m *V1AzureMachinePoolCloudConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOsDisk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AzureMachinePoolCloudConfigEntity) validateOsDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.OsDisk) { // not required
		return nil
	}

	if m.OsDisk != nil {
		if err := m.OsDisk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osDisk")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AzureMachinePoolCloudConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AzureMachinePoolCloudConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1AzureMachinePoolCloudConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}