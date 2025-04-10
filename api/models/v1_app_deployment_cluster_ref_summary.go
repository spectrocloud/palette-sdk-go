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

// V1AppDeploymentClusterRefSummary Application deployment cluster reference
//
// swagger:model v1AppDeploymentClusterRefSummary
type V1AppDeploymentClusterRefSummary struct {

	// Application deployment source cluster type[ "virtualCluster", "hostCluster" ]
	// Enum: ["virtual","host"]
	DeploymentClusterType string `json:"deploymentClusterType,omitempty"`

	// Application deployment source cluster name
	Name string `json:"name,omitempty"`

	// Application deployment source cluster uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 app deployment cluster ref summary
func (m *V1AppDeploymentClusterRefSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentClusterType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1AppDeploymentClusterRefSummaryTypeDeploymentClusterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual","host"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1AppDeploymentClusterRefSummaryTypeDeploymentClusterTypePropEnum = append(v1AppDeploymentClusterRefSummaryTypeDeploymentClusterTypePropEnum, v)
	}
}

const (

	// V1AppDeploymentClusterRefSummaryDeploymentClusterTypeVirtual captures enum value "virtual"
	V1AppDeploymentClusterRefSummaryDeploymentClusterTypeVirtual string = "virtual"

	// V1AppDeploymentClusterRefSummaryDeploymentClusterTypeHost captures enum value "host"
	V1AppDeploymentClusterRefSummaryDeploymentClusterTypeHost string = "host"
)

// prop value enum
func (m *V1AppDeploymentClusterRefSummary) validateDeploymentClusterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1AppDeploymentClusterRefSummaryTypeDeploymentClusterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1AppDeploymentClusterRefSummary) validateDeploymentClusterType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentClusterType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeploymentClusterTypeEnum("deploymentClusterType", "body", m.DeploymentClusterType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 app deployment cluster ref summary based on context it is used
func (m *V1AppDeploymentClusterRefSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AppDeploymentClusterRefSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppDeploymentClusterRefSummary) UnmarshalBinary(b []byte) error {
	var res V1AppDeploymentClusterRefSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
