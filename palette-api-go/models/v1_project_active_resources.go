// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ProjectActiveResources Active project resources
//
// swagger:model v1ProjectActiveResources
type V1ProjectActiveResources struct {

	// app deployments
	AppDeployments *V1ProjectActiveAppDeployments `json:"appDeployments,omitempty"`

	// clusters
	Clusters *V1ProjectActiveClusters `json:"clusters,omitempty"`

	// virtual clusters
	VirtualClusters *V1ProjectActiveClusters `json:"virtualClusters,omitempty"`
}

// Validate validates this v1 project active resources
func (m *V1ProjectActiveResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppDeployments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProjectActiveResources) validateAppDeployments(formats strfmt.Registry) error {

	if swag.IsZero(m.AppDeployments) { // not required
		return nil
	}

	if m.AppDeployments != nil {
		if err := m.AppDeployments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appDeployments")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectActiveResources) validateClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	if m.Clusters != nil {
		if err := m.Clusters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters")
			}
			return err
		}
	}

	return nil
}

func (m *V1ProjectActiveResources) validateVirtualClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualClusters) { // not required
		return nil
	}

	if m.VirtualClusters != nil {
		if err := m.VirtualClusters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualClusters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectActiveResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectActiveResources) UnmarshalBinary(b []byte) error {
	var res V1ProjectActiveResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
