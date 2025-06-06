// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ControlPlaneEndPoint v1 control plane end point
//
// swagger:model v1ControlPlaneEndPoint
type V1ControlPlaneEndPoint struct {

	// DDNSSearchDomain is the search domain used for resolving IP addresses when the EndpointType is DDNS. This search domain is appended to the generated Hostname to obtain the complete DNS name for the endpoint. If Host is already a DDNS FQDN, DDNSSearchDomain is not required
	DdnsSearchDomain string `json:"ddnsSearchDomain,omitempty"`

	// IP or FQDN(External/DDNS)
	Host string `json:"host,omitempty"`

	// VIP or External
	// Enum: ["VIP","External","DDNS"]
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 control plane end point
func (m *V1ControlPlaneEndPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1ControlPlaneEndPointTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VIP","External","DDNS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ControlPlaneEndPointTypeTypePropEnum = append(v1ControlPlaneEndPointTypeTypePropEnum, v)
	}
}

const (

	// V1ControlPlaneEndPointTypeVIP captures enum value "VIP"
	V1ControlPlaneEndPointTypeVIP string = "VIP"

	// V1ControlPlaneEndPointTypeExternal captures enum value "External"
	V1ControlPlaneEndPointTypeExternal string = "External"

	// V1ControlPlaneEndPointTypeDDNS captures enum value "DDNS"
	V1ControlPlaneEndPointTypeDDNS string = "DDNS"
)

// prop value enum
func (m *V1ControlPlaneEndPoint) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1ControlPlaneEndPointTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1ControlPlaneEndPoint) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 control plane end point based on context it is used
func (m *V1ControlPlaneEndPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ControlPlaneEndPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ControlPlaneEndPoint) UnmarshalBinary(b []byte) error {
	var res V1ControlPlaneEndPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
