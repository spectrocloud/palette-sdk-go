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

// V1CoxEdgeMachinePoolCloudConfigEntity v1 cox edge machine pool cloud config entity
//
// swagger:model v1CoxEdgeMachinePoolCloudConfigEntity
type V1CoxEdgeMachinePoolCloudConfigEntity struct {

	// deployments
	Deployments []*V1CoxEdgeDeployment `json:"deployments"`

	// Array of coxedge load persistent storages
	// Unique: true
	PersistentStorages []*V1CoxEdgeLoadPersistentStorage `json:"persistentStorages"`

	// security group rules
	SecurityGroupRules []*V1CoxEdgeSecurityGroupRule `json:"securityGroupRules"`

	// spec
	Spec string `json:"spec,omitempty"`
}

// Validate validates this v1 cox edge machine pool cloud config entity
func (m *V1CoxEdgeMachinePoolCloudConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersistentStorages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CoxEdgeMachinePoolCloudConfigEntity) validateDeployments(formats strfmt.Registry) error {

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

func (m *V1CoxEdgeMachinePoolCloudConfigEntity) validatePersistentStorages(formats strfmt.Registry) error {

	if swag.IsZero(m.PersistentStorages) { // not required
		return nil
	}

	if err := validate.UniqueItems("persistentStorages", "body", m.PersistentStorages); err != nil {
		return err
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

func (m *V1CoxEdgeMachinePoolCloudConfigEntity) validateSecurityGroupRules(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroupRules) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityGroupRules); i++ {
		if swag.IsZero(m.SecurityGroupRules[i]) { // not required
			continue
		}

		if m.SecurityGroupRules[i] != nil {
			if err := m.SecurityGroupRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityGroupRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CoxEdgeMachinePoolCloudConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CoxEdgeMachinePoolCloudConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1CoxEdgeMachinePoolCloudConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}