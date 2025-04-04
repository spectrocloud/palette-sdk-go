// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UserInfo User basic information
//
// swagger:model v1UserInfo
type V1UserInfo struct {

	// Organization name
	OrgName string `json:"orgName,omitempty"`

	// tenant Uid
	TenantUID string `json:"tenantUid,omitempty"`

	// user Uid
	UserUID string `json:"userUid,omitempty"`
}

// Validate validates this v1 user info
func (m *V1UserInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 user info based on context it is used
func (m *V1UserInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UserInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UserInfo) UnmarshalBinary(b []byte) error {
	var res V1UserInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
