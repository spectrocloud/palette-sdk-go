// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SystemReverseProxy system config reverse proxy
//
// swagger:model v1SystemReverseProxy
type V1SystemReverseProxy struct {

	// ca cert
	CaCert string `json:"caCert,omitempty"`

	// client cert
	ClientCert string `json:"clientCert,omitempty"`

	// client key
	ClientKey string `json:"clientKey,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// protocol
	// Enum: [http https]
	Protocol string `json:"protocol,omitempty"`

	// server
	Server string `json:"server,omitempty"`

	// v host port
	VHostPort int64 `json:"vHostPort,omitempty"`
}

// Validate validates this v1 system reverse proxy
func (m *V1SystemReverseProxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1SystemReverseProxyTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","https"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SystemReverseProxyTypeProtocolPropEnum = append(v1SystemReverseProxyTypeProtocolPropEnum, v)
	}
}

const (

	// V1SystemReverseProxyProtocolHTTP captures enum value "http"
	V1SystemReverseProxyProtocolHTTP string = "http"

	// V1SystemReverseProxyProtocolHTTPS captures enum value "https"
	V1SystemReverseProxyProtocolHTTPS string = "https"
)

// prop value enum
func (m *V1SystemReverseProxy) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1SystemReverseProxyTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1SystemReverseProxy) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SystemReverseProxy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SystemReverseProxy) UnmarshalBinary(b []byte) error {
	var res V1SystemReverseProxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
