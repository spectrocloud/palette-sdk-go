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

// V1VMDHCPOptions Extra DHCP options to use in the interface.
//
// swagger:model v1VmDHCPOptions
type V1VMDHCPOptions struct {

	// If specified will pass option 67 to interface's DHCP server
	BootFileName string `json:"bootFileName,omitempty"`

	// If specified will pass the configured NTP server to the VM via DHCP option 042.
	NtpServers []string `json:"ntpServers"`

	// If specified will pass extra DHCP options for private use, range: 224-254
	PrivateOptions []*V1VMDHCPPrivateOptions `json:"privateOptions"`

	// If specified will pass option 66 to interface's DHCP server
	TftpServerName string `json:"tftpServerName,omitempty"`
}

// Validate validates this v1 Vm d h c p options
func (m *V1VMDHCPOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMDHCPOptions) validatePrivateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivateOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.PrivateOptions); i++ {
		if swag.IsZero(m.PrivateOptions[i]) { // not required
			continue
		}

		if m.PrivateOptions[i] != nil {
			if err := m.PrivateOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privateOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMDHCPOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMDHCPOptions) UnmarshalBinary(b []byte) error {
	var res V1VMDHCPOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}