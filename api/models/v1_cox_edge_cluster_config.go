// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1CoxEdgeClusterConfig Cluster level configuration for coxedge cloud and applicable for all the machine pools
//
// swagger:model v1CoxEdgeClusterConfig
type V1CoxEdgeClusterConfig struct {

	// cox edge load balancer config
	// Required: true
	CoxEdgeLoadBalancerConfig *V1CoxEdgeLoadBalancerConfig `json:"coxEdgeLoadBalancerConfig"`

	// cox edge worker load balancer config
	CoxEdgeWorkerLoadBalancerConfig *V1CoxEdgeLoadBalancerConfig `json:"coxEdgeWorkerLoadBalancerConfig,omitempty"`

	// environment
	Environment string `json:"environment,omitempty"`

	// organization Id
	OrganizationID string `json:"organizationId,omitempty"`

	// CoxEdge ssh authorized keys
	// Required: true
	SSHAuthorizedKeys []string `json:"sshAuthorizedKeys"`
}

// Validate validates this v1 cox edge cluster config
func (m *V1CoxEdgeClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoxEdgeLoadBalancerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoxEdgeWorkerLoadBalancerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHAuthorizedKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CoxEdgeClusterConfig) validateCoxEdgeLoadBalancerConfig(formats strfmt.Registry) error {

	if err := validate.Required("coxEdgeLoadBalancerConfig", "body", m.CoxEdgeLoadBalancerConfig); err != nil {
		return err
	}

	if m.CoxEdgeLoadBalancerConfig != nil {
		if err := m.CoxEdgeLoadBalancerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coxEdgeLoadBalancerConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1CoxEdgeClusterConfig) validateCoxEdgeWorkerLoadBalancerConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.CoxEdgeWorkerLoadBalancerConfig) { // not required
		return nil
	}

	if m.CoxEdgeWorkerLoadBalancerConfig != nil {
		if err := m.CoxEdgeWorkerLoadBalancerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coxEdgeWorkerLoadBalancerConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1CoxEdgeClusterConfig) validateSSHAuthorizedKeys(formats strfmt.Registry) error {

	if err := validate.Required("sshAuthorizedKeys", "body", m.SSHAuthorizedKeys); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CoxEdgeClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CoxEdgeClusterConfig) UnmarshalBinary(b []byte) error {
	var res V1CoxEdgeClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
