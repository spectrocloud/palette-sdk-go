// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CoxEdgeDeployment v1 cox edge deployment
//
// swagger:model v1CoxEdgeDeployment
type V1CoxEdgeDeployment struct {

	// cpu utilization
	CPUUtilization int32 `json:"cpuUtilization,omitempty"`

	// enable auto scaling
	EnableAutoScaling bool `json:"enableAutoScaling,omitempty"`

	// instances per pop
	InstancesPerPop int32 `json:"instancesPerPop,omitempty"`

	// max instances per pop
	MaxInstancesPerPop int32 `json:"maxInstancesPerPop,omitempty"`

	// min instances per pop
	MinInstancesPerPop int32 `json:"minInstancesPerPop,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pops
	Pops []string `json:"pops"`
}

// Validate validates this v1 cox edge deployment
func (m *V1CoxEdgeDeployment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CoxEdgeDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CoxEdgeDeployment) UnmarshalBinary(b []byte) error {
	var res V1CoxEdgeDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
