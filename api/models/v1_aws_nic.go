// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AwsNic AWS network interface
//
// swagger:model v1AwsNic
type V1AwsNic struct {

	// index
	Index int8 `json:"index,omitempty"`

	// private i ps
	PrivateIPs []string `json:"privateIPs"`

	// public Ip
	PublicIP string `json:"publicIp,omitempty"`
}

// Validate validates this v1 aws nic
func (m *V1AwsNic) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 aws nic based on context it is used
func (m *V1AwsNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsNic) UnmarshalBinary(b []byte) error {
	var res V1AwsNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
