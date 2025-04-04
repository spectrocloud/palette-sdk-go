// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1EdgeHostNetwork Network defines the network configuration for a virtual machine
//
// swagger:model v1EdgeHostNetwork
type V1EdgeHostNetwork struct {

	// NetworkName of the network where this machine will be connected
	// Required: true
	NetworkName *string `json:"networkName"`

	// NetworkType  specifies the type of network
	// Required: true
	// Enum: ["default","bridge"]
	NetworkType *string `json:"networkType"`
}

// Validate validates this v1 edge host network
func (m *V1EdgeHostNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeHostNetwork) validateNetworkName(formats strfmt.Registry) error {

	if err := validate.Required("networkName", "body", m.NetworkName); err != nil {
		return err
	}

	return nil
}

var v1EdgeHostNetworkTypeNetworkTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","bridge"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1EdgeHostNetworkTypeNetworkTypePropEnum = append(v1EdgeHostNetworkTypeNetworkTypePropEnum, v)
	}
}

const (

	// V1EdgeHostNetworkNetworkTypeDefault captures enum value "default"
	V1EdgeHostNetworkNetworkTypeDefault string = "default"

	// V1EdgeHostNetworkNetworkTypeBridge captures enum value "bridge"
	V1EdgeHostNetworkNetworkTypeBridge string = "bridge"
)

// prop value enum
func (m *V1EdgeHostNetwork) validateNetworkTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1EdgeHostNetworkTypeNetworkTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1EdgeHostNetwork) validateNetworkType(formats strfmt.Registry) error {

	if err := validate.Required("networkType", "body", m.NetworkType); err != nil {
		return err
	}

	// value enum
	if err := m.validateNetworkTypeEnum("networkType", "body", *m.NetworkType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 edge host network based on context it is used
func (m *V1EdgeHostNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeHostNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeHostNetwork) UnmarshalBinary(b []byte) error {
	var res V1EdgeHostNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
