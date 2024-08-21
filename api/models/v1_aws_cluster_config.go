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

// V1AwsClusterConfig Cluster level configuration for aws cloud and applicable for all the machine pools
//
// swagger:model v1AwsClusterConfig
type V1AwsClusterConfig struct {

	// Create bastion node option we have earlier supported creation of bastion by default capa seems to favour session manager against bastion node https://github.com/kubernetes-sigs/cluster-api-provider-aws/issues/947
	BastionDisabled bool `json:"bastionDisabled,omitempty"`

	// ControlPlaneLoadBalancer specifies how API server elb will be configured, this field is optional, not provided, "", default => "Internet-facing" "Internet-facing" => "Internet-facing" "internal" => "internal" For spectro saas setup we require to talk to the apiserver from our cluster so ControlPlaneLoadBalancer should be "", not provided or "Internet-facing"
	ControlPlaneLoadBalancer string `json:"controlPlaneLoadBalancer,omitempty"`

	// region
	// Required: true
	Region *string `json:"region"`

	// ssh key name
	SSHKeyName string `json:"sshKeyName,omitempty"`

	// VPC Id to deploy cluster into should have one public and one private subnet for the the cluster creation, this field is optional, If VPC Id is not provided a fully managed VPC will be created
	VpcID string `json:"vpcId,omitempty"`
}

// Validate validates this v1 aws cluster config
func (m *V1AwsClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsClusterConfig) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsClusterConfig) UnmarshalBinary(b []byte) error {
	var res V1AwsClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}