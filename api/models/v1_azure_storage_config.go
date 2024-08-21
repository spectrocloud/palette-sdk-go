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

// V1AzureStorageConfig Azure storage config object
//
// swagger:model v1AzureStorageConfig
type V1AzureStorageConfig struct {

	// Azure container name
	// Required: true
	ContainerName *string `json:"containerName"`

	// Azure cloud account credentials
	// Required: true
	Credentials *V1AzureAccountEntitySpec `json:"credentials"`

	// Azure resource group name, to which the storage account is mapped
	// Required: true
	ResourceGroup *string `json:"resourceGroup"`

	// Azure sku
	Sku string `json:"sku,omitempty"`

	// Azure storage name
	// Required: true
	StorageName *string `json:"storageName"`
}

// Validate validates this v1 azure storage config
func (m *V1AzureStorageConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AzureStorageConfig) validateContainerName(formats strfmt.Registry) error {

	if err := validate.Required("containerName", "body", m.ContainerName); err != nil {
		return err
	}

	return nil
}

func (m *V1AzureStorageConfig) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("credentials", "body", m.Credentials); err != nil {
		return err
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *V1AzureStorageConfig) validateResourceGroup(formats strfmt.Registry) error {

	if err := validate.Required("resourceGroup", "body", m.ResourceGroup); err != nil {
		return err
	}

	return nil
}

func (m *V1AzureStorageConfig) validateStorageName(formats strfmt.Registry) error {

	if err := validate.Required("storageName", "body", m.StorageName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AzureStorageConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AzureStorageConfig) UnmarshalBinary(b []byte) error {
	var res V1AzureStorageConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}