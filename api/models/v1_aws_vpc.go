// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1AwsVpc A virtual network dedicated to a AWS account
//
// swagger:model v1AwsVpc
type V1AwsVpc struct {

	// cidr block
	CidrBlock string `json:"cidrBlock,omitempty"`

	// Name of the virtual network
	Name string `json:"name,omitempty"`

	// List of subnets associated to a AWS VPC
	Subnets []*V1AwsSubnet `json:"subnets"`

	// Id of the virtual network
	// Required: true
	VpcID *string `json:"vpcId"`
}

// Validate validates this v1 aws vpc
func (m *V1AwsVpc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubnets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsVpc) validateSubnets(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnets) { // not required
		return nil
	}

	for i := 0; i < len(m.Subnets); i++ {
		if swag.IsZero(m.Subnets[i]) { // not required
			continue
		}

		if m.Subnets[i] != nil {
			if err := m.Subnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AwsVpc) validateVpcID(formats strfmt.Registry) error {

	if err := validate.Required("vpcId", "body", m.VpcID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsVpc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsVpc) UnmarshalBinary(b []byte) error {
	var res V1AwsVpc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}