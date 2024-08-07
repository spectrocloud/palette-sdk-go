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

// V1CoxEdgeMachineSpec CoxEdge cloud VM definition spec
//
// swagger:model v1CoxEdgeMachineSpec
type V1CoxEdgeMachineSpec struct {

	// add anycast Ip address
	AddAnycastIPAddress bool `json:"addAnycastIpAddress,omitempty"`

	// deployments
	Deployments []*V1CoxEdgeDeployment `json:"deployments"`

	// image
	Image string `json:"image,omitempty"`

	// persistent storages
	PersistentStorages []*V1CoxEdgeLoadPersistentStorage `json:"persistentStorages"`

	// ports
	Ports []*V1CoxEdgePort `json:"ports"`

	// provider Id
	ProviderID string `json:"providerId,omitempty"`

	// specs
	Specs string `json:"specs,omitempty"`

	// ssh authorized keys
	SSHAuthorizedKeys []string `json:"sshAuthorizedKeys"`
}

// Validate validates this v1 cox edge machine spec
func (m *V1CoxEdgeMachineSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersistentStorages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CoxEdgeMachineSpec) validateDeployments(formats strfmt.Registry) error {

	if swag.IsZero(m.Deployments) { // not required
		return nil
	}

	for i := 0; i < len(m.Deployments); i++ {
		if swag.IsZero(m.Deployments[i]) { // not required
			continue
		}

		if m.Deployments[i] != nil {
			if err := m.Deployments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CoxEdgeMachineSpec) validatePersistentStorages(formats strfmt.Registry) error {

	if swag.IsZero(m.PersistentStorages) { // not required
		return nil
	}

	for i := 0; i < len(m.PersistentStorages); i++ {
		if swag.IsZero(m.PersistentStorages[i]) { // not required
			continue
		}

		if m.PersistentStorages[i] != nil {
			if err := m.PersistentStorages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("persistentStorages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CoxEdgeMachineSpec) validatePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CoxEdgeMachineSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CoxEdgeMachineSpec) UnmarshalBinary(b []byte) error {
	var res V1CoxEdgeMachineSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
