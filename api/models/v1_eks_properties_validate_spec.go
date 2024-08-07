// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EksPropertiesValidateSpec Eks properties validate spec
//
// swagger:model V1EksPropertiesValidateSpec
type V1EksPropertiesValidateSpec struct {

	// cloud account Uid
	CloudAccountUID string `json:"cloudAccountUid,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// ssh key name
	SSHKeyName string `json:"sshKeyName,omitempty"`

	// subnets
	Subnets []string `json:"subnets"`

	// vpc Id
	VpcID string `json:"vpcId,omitempty"`
}

// Validate validates this v1 eks properties validate spec
func (m *V1EksPropertiesValidateSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1EksPropertiesValidateSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EksPropertiesValidateSpec) UnmarshalBinary(b []byte) error {
	var res V1EksPropertiesValidateSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
