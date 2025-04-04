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

// V1LoadBalancerSpec LoadBalancerSpec defines an Azure load balancer.
//
// swagger:model v1LoadBalancerSpec
type V1LoadBalancerSpec struct {

	// api server l b static IP
	APIServerLBStaticIP string `json:"apiServerLBStaticIP,omitempty"`

	// ip allocation method
	// Enum: ["Static","Dynamic"]
	IPAllocationMethod *string `json:"ipAllocationMethod,omitempty"`

	// private DNS name
	PrivateDNSName string `json:"privateDNSName,omitempty"`

	// Load Balancer type
	// Enum: ["Internal","Public"]
	Type *string `json:"type,omitempty"`
}

// Validate validates this v1 load balancer spec
func (m *V1LoadBalancerSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAllocationMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1LoadBalancerSpecTypeIPAllocationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Static","Dynamic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1LoadBalancerSpecTypeIPAllocationMethodPropEnum = append(v1LoadBalancerSpecTypeIPAllocationMethodPropEnum, v)
	}
}

const (

	// V1LoadBalancerSpecIPAllocationMethodStatic captures enum value "Static"
	V1LoadBalancerSpecIPAllocationMethodStatic string = "Static"

	// V1LoadBalancerSpecIPAllocationMethodDynamic captures enum value "Dynamic"
	V1LoadBalancerSpecIPAllocationMethodDynamic string = "Dynamic"
)

// prop value enum
func (m *V1LoadBalancerSpec) validateIPAllocationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1LoadBalancerSpecTypeIPAllocationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1LoadBalancerSpec) validateIPAllocationMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAllocationMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPAllocationMethodEnum("ipAllocationMethod", "body", *m.IPAllocationMethod); err != nil {
		return err
	}

	return nil
}

var v1LoadBalancerSpecTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Internal","Public"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1LoadBalancerSpecTypeTypePropEnum = append(v1LoadBalancerSpecTypeTypePropEnum, v)
	}
}

const (

	// V1LoadBalancerSpecTypeInternal captures enum value "Internal"
	V1LoadBalancerSpecTypeInternal string = "Internal"

	// V1LoadBalancerSpecTypePublic captures enum value "Public"
	V1LoadBalancerSpecTypePublic string = "Public"
)

// prop value enum
func (m *V1LoadBalancerSpec) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1LoadBalancerSpecTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1LoadBalancerSpec) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 load balancer spec based on context it is used
func (m *V1LoadBalancerSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1LoadBalancerSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LoadBalancerSpec) UnmarshalBinary(b []byte) error {
	var res V1LoadBalancerSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
