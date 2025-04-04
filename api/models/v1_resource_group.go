// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ResourceGroup Azure resource Group is a container that holds related resources for an Azure solution
//
// swagger:model v1ResourceGroup
type V1ResourceGroup struct {

	// The ID of the resource group
	ID string `json:"id,omitempty"`

	// The location of the resource group. It cannot be changed after the resource group has been created
	Location string `json:"location,omitempty"`

	// The type of the resource group
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 resource group
func (m *V1ResourceGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 resource group based on context it is used
func (m *V1ResourceGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ResourceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ResourceGroup) UnmarshalBinary(b []byte) error {
	var res V1ResourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
