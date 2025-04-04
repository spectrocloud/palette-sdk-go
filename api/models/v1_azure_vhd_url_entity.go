// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AzureVhdURLEntity Azure vhd url entity
//
// swagger:model v1AzureVhdUrlEntity
type V1AzureVhdURLEntity struct {

	// The name of the resource
	Name string `json:"name,omitempty"`

	// The url of the Azure Vhd
	URL string `json:"url,omitempty"`
}

// Validate validates this v1 azure vhd Url entity
func (m *V1AzureVhdURLEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 azure vhd Url entity based on context it is used
func (m *V1AzureVhdURLEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AzureVhdURLEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AzureVhdURLEntity) UnmarshalBinary(b []byte) error {
	var res V1AzureVhdURLEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
