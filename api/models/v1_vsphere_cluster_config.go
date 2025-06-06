// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1VsphereClusterConfig v1 vsphere cluster config
//
// swagger:model v1VsphereClusterConfig
type V1VsphereClusterConfig struct {

	// The optional control plane endpoint, which can be an IP or FQDN
	ControlPlaneEndpoint *V1ControlPlaneEndPoint `json:"controlPlaneEndpoint,omitempty"`

	// NTPServers is a list of NTP servers to use instead of the machine image's default NTP server list.
	NtpServers []string `json:"ntpServers"`

	// Placement configuration Placement config in ClusterConfig serve as default values for each MachinePool
	// Required: true
	Placement *V1VspherePlacementConfig `json:"placement"`

	// SSHKeys specifies a list of ssh authorized keys for the 'spectro' user
	SSHKeys []string `json:"sshKeys"`

	// whether this cluster should use dhcp or static IP, if false, use DHCP if this is set, then all machinepools should have staticIP with provided IPPool adding this as an additional standalone flag without relating to placement.Nework main reason is to enable more validation for placement.Network.StaticIP which should come together with valid Network.IPPool and Network.Name
	StaticIP bool `json:"staticIp,omitempty"`
}

// Validate validates this v1 vsphere cluster config
func (m *V1VsphereClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControlPlaneEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlacement(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VsphereClusterConfig) validateControlPlaneEndpoint(formats strfmt.Registry) error {
	if swag.IsZero(m.ControlPlaneEndpoint) { // not required
		return nil
	}

	if m.ControlPlaneEndpoint != nil {
		if err := m.ControlPlaneEndpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controlPlaneEndpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controlPlaneEndpoint")
			}
			return err
		}
	}

	return nil
}

func (m *V1VsphereClusterConfig) validatePlacement(formats strfmt.Registry) error {

	if err := validate.Required("placement", "body", m.Placement); err != nil {
		return err
	}

	if m.Placement != nil {
		if err := m.Placement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("placement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("placement")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 vsphere cluster config based on the context it is used
func (m *V1VsphereClusterConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateControlPlaneEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlacement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VsphereClusterConfig) contextValidateControlPlaneEndpoint(ctx context.Context, formats strfmt.Registry) error {

	if m.ControlPlaneEndpoint != nil {

		if swag.IsZero(m.ControlPlaneEndpoint) { // not required
			return nil
		}

		if err := m.ControlPlaneEndpoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controlPlaneEndpoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controlPlaneEndpoint")
			}
			return err
		}
	}

	return nil
}

func (m *V1VsphereClusterConfig) contextValidatePlacement(ctx context.Context, formats strfmt.Registry) error {

	if m.Placement != nil {

		if err := m.Placement.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("placement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("placement")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VsphereClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VsphereClusterConfig) UnmarshalBinary(b []byte) error {
	var res V1VsphereClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
