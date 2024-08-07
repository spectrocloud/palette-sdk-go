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

// V1VMPort Port represents a port to expose from the virtual machine. Default protocol TCP. The port field is mandatory
//
// swagger:model v1VmPort
type V1VMPort struct {

	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.
	Name string `json:"name,omitempty"`

	// Number of port to expose for the virtual machine. This must be a valid port number, 0 < x < 65536.
	// Required: true
	Port *int32 `json:"port"`

	// Protocol for port. Must be UDP or TCP. Defaults to "TCP".
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this v1 Vm port
func (m *V1VMPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMPort) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMPort) UnmarshalBinary(b []byte) error {
	var res V1VMPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
