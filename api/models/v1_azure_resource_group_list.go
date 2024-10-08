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

// V1AzureResourceGroupList List of Azure resource group
//
// swagger:model v1AzureResourceGroupList
type V1AzureResourceGroupList struct {

	// resource group list
	ResourceGroupList []*V1ResourceGroup `json:"resourceGroupList"`
}

// Validate validates this v1 azure resource group list
func (m *V1AzureResourceGroupList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceGroupList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AzureResourceGroupList) validateResourceGroupList(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroupList) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceGroupList); i++ {
		if swag.IsZero(m.ResourceGroupList[i]) { // not required
			continue
		}

		if m.ResourceGroupList[i] != nil {
			if err := m.ResourceGroupList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceGroupList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AzureResourceGroupList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AzureResourceGroupList) UnmarshalBinary(b []byte) error {
	var res V1AzureResourceGroupList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
