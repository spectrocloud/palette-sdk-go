// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1TenantEmailPatch Tenant EmailId
//
// swagger:model v1TenantEmailPatch
type V1TenantEmailPatch struct {

	// email Id
	EmailID string `json:"emailId,omitempty"`
}

// Validate validates this v1 tenant email patch
func (m *V1TenantEmailPatch) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1TenantEmailPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TenantEmailPatch) UnmarshalBinary(b []byte) error {
	var res V1TenantEmailPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}