// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterConfigResponse v1 cluster config response
//
// swagger:model v1ClusterConfigResponse
type V1ClusterConfigResponse struct {

	// HostClusterConfig defines the configuration entity of host clusters config entity
	HostClusterConfig *V1HostClusterConfigResponse `json:"hostClusterConfig,omitempty"`
}

// Validate validates this v1 cluster config response
func (m *V1ClusterConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterConfigResponse) validateHostClusterConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.HostClusterConfig) { // not required
		return nil
	}

	if m.HostClusterConfig != nil {
		if err := m.HostClusterConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostClusterConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterConfigResponse) UnmarshalBinary(b []byte) error {
	var res V1ClusterConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
