// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1VMHTTPGetAction HTTPGetAction describes an action based on HTTP Get requests.
//
// swagger:model v1VmHttpGetAction
type V1VMHTTPGetAction struct {

	// Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
	Host string `json:"host,omitempty"`

	// Custom headers to set in the request. HTTP allows repeated headers.
	HTTPHeaders []*V1VMHTTPHeader `json:"httpHeaders"`

	// Path to access on the HTTP server.
	Path string `json:"path,omitempty"`

	// Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	// Required: true
	Port *string `json:"port"`

	// Scheme to use for connecting to the host. Defaults to HTTP.
	Scheme string `json:"scheme,omitempty"`
}

// Validate validates this v1 Vm Http get action
func (m *V1VMHTTPGetAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMHTTPGetAction) validateHTTPHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPHeaders) { // not required
		return nil
	}

	for i := 0; i < len(m.HTTPHeaders); i++ {
		if swag.IsZero(m.HTTPHeaders[i]) { // not required
			continue
		}

		if m.HTTPHeaders[i] != nil {
			if err := m.HTTPHeaders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("httpHeaders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("httpHeaders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VMHTTPGetAction) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 Vm Http get action based on the context it is used
func (m *V1VMHTTPGetAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHTTPHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMHTTPGetAction) contextValidateHTTPHeaders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HTTPHeaders); i++ {

		if m.HTTPHeaders[i] != nil {

			if swag.IsZero(m.HTTPHeaders[i]) { // not required
				return nil
			}

			if err := m.HTTPHeaders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("httpHeaders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("httpHeaders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMHTTPGetAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMHTTPGetAction) UnmarshalBinary(b []byte) error {
	var res V1VMHTTPGetAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
