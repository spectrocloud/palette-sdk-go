// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AwsHybridConfig AwsHybridConfig specifies the AWS Hybrid configuration for the cluster
//
// swagger:model v1AwsHybridConfig
type V1AwsHybridConfig struct {

	// AWS VPC CIDR is the CIDR of the AWS/EKS cluster's VPC
	AwsVpcCidr string `json:"awsVpcCidr,omitempty"`

	// IamRolesAnywhere specifies the IAM Roles Anywhere configuration for the AWS/EKS cluster
	IamRolesAnywhere *V1IamRolesAnywhere `json:"iamRolesAnywhere,omitempty"`

	// RemoteNodeCIDRs specifies the Node CIDRs of all remote nodes
	RemoteNodeCidrs []string `json:"remoteNodeCidrs"`

	// RemotePodCIDRs specifies the Pod CIDRs of all remote pods
	RemotePodCidrs []string `json:"remotePodCidrs"`

	// SystemsManager specifies the Systems Manager configuration for the AWS/EKS cluster
	SystemsManager *V1SystemsManager `json:"systemsManager,omitempty"`
}

// Validate validates this v1 aws hybrid config
func (m *V1AwsHybridConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIamRolesAnywhere(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemsManager(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsHybridConfig) validateIamRolesAnywhere(formats strfmt.Registry) error {

	if swag.IsZero(m.IamRolesAnywhere) { // not required
		return nil
	}

	if m.IamRolesAnywhere != nil {
		if err := m.IamRolesAnywhere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iamRolesAnywhere")
			}
			return err
		}
	}

	return nil
}

func (m *V1AwsHybridConfig) validateSystemsManager(formats strfmt.Registry) error {

	if swag.IsZero(m.SystemsManager) { // not required
		return nil
	}

	if m.SystemsManager != nil {
		if err := m.SystemsManager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systemsManager")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsHybridConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsHybridConfig) UnmarshalBinary(b []byte) error {
	var res V1AwsHybridConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}