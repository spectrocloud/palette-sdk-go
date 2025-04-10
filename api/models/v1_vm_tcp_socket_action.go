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

// V1VMTCPSocketAction TCPSocketAction describes an action based on opening a socket
//
// swagger:model v1VmTcpSocketAction
type V1VMTCPSocketAction struct {

	// Optional: Host name to connect to, defaults to the pod IP.
	Host string `json:"host,omitempty"`

	// Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	// Required: true
	Port *string `json:"port"`
}

// Validate validates this v1 Vm Tcp socket action
func (m *V1VMTCPSocketAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMTCPSocketAction) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 Vm Tcp socket action based on context it is used
func (m *V1VMTCPSocketAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMTCPSocketAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMTCPSocketAction) UnmarshalBinary(b []byte) error {
	var res V1VMTCPSocketAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
