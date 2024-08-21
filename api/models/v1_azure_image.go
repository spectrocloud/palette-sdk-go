// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AzureImage Refers to Azure Shared Gallery image
//
// swagger:model v1AzureImage
type V1AzureImage struct {

	// gallery
	Gallery string `json:"gallery,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// resource group
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// subscription ID
	SubscriptionID string `json:"subscriptionID,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 azure image
func (m *V1AzureImage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AzureImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AzureImage) UnmarshalBinary(b []byte) error {
	var res V1AzureImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}