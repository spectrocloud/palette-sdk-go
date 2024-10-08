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

// V1EdgeMachinePoolHost EdgeHost of Edge clusters
//
// swagger:model v1EdgeMachinePoolHost
type V1EdgeMachinePoolHost struct {

	// HostAddress is a FQDN or IP address of the Host
	// Required: true
	HostAddress *string `json:"hostAddress"`

	// HostIdentity is the identity to access the edge host
	HostIdentity *V1EdgeMachinePoolHostIdentity `json:"hostIdentity,omitempty"`

	// HostName is the name of the EdgeHost
	HostName string `json:"hostName,omitempty"`

	// HostUid is the ID of the EdgeHost
	// Required: true
	HostUID *string `json:"hostUid"`
}

// Validate validates this v1 edge machine pool host
func (m *V1EdgeMachinePoolHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EdgeMachinePoolHost) validateHostAddress(formats strfmt.Registry) error {

	if err := validate.Required("hostAddress", "body", m.HostAddress); err != nil {
		return err
	}

	return nil
}

func (m *V1EdgeMachinePoolHost) validateHostIdentity(formats strfmt.Registry) error {

	if swag.IsZero(m.HostIdentity) { // not required
		return nil
	}

	if m.HostIdentity != nil {
		if err := m.HostIdentity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostIdentity")
			}
			return err
		}
	}

	return nil
}

func (m *V1EdgeMachinePoolHost) validateHostUID(formats strfmt.Registry) error {

	if err := validate.Required("hostUid", "body", m.HostUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EdgeMachinePoolHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EdgeMachinePoolHost) UnmarshalBinary(b []byte) error {
	var res V1EdgeMachinePoolHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
