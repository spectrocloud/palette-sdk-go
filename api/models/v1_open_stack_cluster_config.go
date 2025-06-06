// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OpenStackClusterConfig Cluster level configuration for OpenStack cloud and applicable for all the machine pools
//
// swagger:model v1OpenStackClusterConfig
type V1OpenStackClusterConfig struct {

	// Create bastion node option we have earlier supported creation of bastion by default
	BastionDisabled bool `json:"bastionDisabled,omitempty"`

	// DNSNameservers is the list of nameservers for OpenStack Subnet being created. Set this value when you need create a new network/subnet while the access through DNS is required.
	DNSNameservers []string `json:"dnsNameservers"`

	// domain
	Domain *V1OpenStackResource `json:"domain,omitempty"`

	// For static placement
	Network *V1OpenStackResource `json:"network,omitempty"`

	// For dynamic provision NodeCIDR is the OpenStack Subnet to be created. Cluster actuator will create a network, a subnet with NodeCIDR, and a router connected to this subnet. If you leave this empty, no network will be created.
	NodeCidr string `json:"nodeCidr,omitempty"`

	// project
	Project *V1OpenStackResource `json:"project,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// ssh key name
	SSHKeyName string `json:"sshKeyName,omitempty"`

	// subnet
	Subnet *V1OpenStackResource `json:"subnet,omitempty"`
}

// Validate validates this v1 open stack cluster config
func (m *V1OpenStackClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OpenStackClusterConfig) validateDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if m.Domain != nil {
		if err := m.Domain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

func (m *V1OpenStackClusterConfig) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *V1OpenStackClusterConfig) validateProject(formats strfmt.Registry) error {
	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *V1OpenStackClusterConfig) validateSubnet(formats strfmt.Registry) error {
	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {
		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 open stack cluster config based on the context it is used
func (m *V1OpenStackClusterConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDomain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OpenStackClusterConfig) contextValidateDomain(ctx context.Context, formats strfmt.Registry) error {

	if m.Domain != nil {

		if swag.IsZero(m.Domain) { // not required
			return nil
		}

		if err := m.Domain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

func (m *V1OpenStackClusterConfig) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.Network != nil {

		if swag.IsZero(m.Network) { // not required
			return nil
		}

		if err := m.Network.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *V1OpenStackClusterConfig) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	if m.Project != nil {

		if swag.IsZero(m.Project) { // not required
			return nil
		}

		if err := m.Project.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *V1OpenStackClusterConfig) contextValidateSubnet(ctx context.Context, formats strfmt.Registry) error {

	if m.Subnet != nil {

		if swag.IsZero(m.Subnet) { // not required
			return nil
		}

		if err := m.Subnet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OpenStackClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OpenStackClusterConfig) UnmarshalBinary(b []byte) error {
	var res V1OpenStackClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
