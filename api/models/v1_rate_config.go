// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1RateConfig Rate config
//
// swagger:model v1RateConfig
type V1RateConfig struct {

	// aws
	Aws *V1PublicCloudRateConfig `json:"aws,omitempty"`

	// azure
	Azure *V1PublicCloudRateConfig `json:"azure,omitempty"`

	// custom
	// Unique: true
	Custom []*V1CustomCloudRateConfig `json:"custom"`

	// edge
	Edge *V1PrivateCloudRateConfig `json:"edge,omitempty"`

	// edge native
	EdgeNative *V1PrivateCloudRateConfig `json:"edgeNative,omitempty"`

	// gcp
	Gcp *V1PublicCloudRateConfig `json:"gcp,omitempty"`

	// generic
	Generic *V1PrivateCloudRateConfig `json:"generic,omitempty"`

	// libvirt
	Libvirt *V1PrivateCloudRateConfig `json:"libvirt,omitempty"`

	// maas
	Maas *V1PrivateCloudRateConfig `json:"maas,omitempty"`

	// openstack
	Openstack *V1PrivateCloudRateConfig `json:"openstack,omitempty"`

	// vsphere
	Vsphere *V1PrivateCloudRateConfig `json:"vsphere,omitempty"`
}

// Validate validates this v1 rate config
func (m *V1RateConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeNative(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLibvirt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RateConfig) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *V1RateConfig) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *V1RateConfig) validateCustom(formats strfmt.Registry) error {

	if swag.IsZero(m.Custom) { // not required
		return nil
	}

	if err := validate.UniqueItems("custom", "body", m.Custom); err != nil {
		return err
	}

	for i := 0; i < len(m.Custom); i++ {
		if swag.IsZero(m.Custom[i]) { // not required
			continue
		}

		if m.Custom[i] != nil {
			if err := m.Custom[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1RateConfig) validateEdge(formats strfmt.Registry) error {

	if swag.IsZero(m.Edge) { // not required
		return nil
	}

	if m.Edge != nil {
		if err := m.Edge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge")
			}
			return err
		}
	}

	return nil
}

func (m *V1RateConfig) validateEdgeNative(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeNative) { // not required
		return nil
	}

	if m.EdgeNative != nil {
		if err := m.EdgeNative.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeNative")
			}
			return err
		}
	}

	return nil
}

func (m *V1RateConfig) validateGcp(formats strfmt.Registry) error {

	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *V1RateConfig) validateGeneric(formats strfmt.Registry) error {

	if swag.IsZero(m.Generic) { // not required
		return nil
	}

	if m.Generic != nil {
		if err := m.Generic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generic")
			}
			return err
		}
	}

	return nil
}

func (m *V1RateConfig) validateLibvirt(formats strfmt.Registry) error {

	if swag.IsZero(m.Libvirt) { // not required
		return nil
	}

	if m.Libvirt != nil {
		if err := m.Libvirt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("libvirt")
			}
			return err
		}
	}

	return nil
}

func (m *V1RateConfig) validateMaas(formats strfmt.Registry) error {

	if swag.IsZero(m.Maas) { // not required
		return nil
	}

	if m.Maas != nil {
		if err := m.Maas.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maas")
			}
			return err
		}
	}

	return nil
}

func (m *V1RateConfig) validateOpenstack(formats strfmt.Registry) error {

	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	if m.Openstack != nil {
		if err := m.Openstack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

func (m *V1RateConfig) validateVsphere(formats strfmt.Registry) error {

	if swag.IsZero(m.Vsphere) { // not required
		return nil
	}

	if m.Vsphere != nil {
		if err := m.Vsphere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphere")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RateConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RateConfig) UnmarshalBinary(b []byte) error {
	var res V1RateConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
