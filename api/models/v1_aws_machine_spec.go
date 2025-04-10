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

// V1AwsMachineSpec AWS cloud VM definition spec
//
// swagger:model v1AwsMachineSpec
type V1AwsMachineSpec struct {

	// Additional Security groups
	AdditionalSecurityGroups []*V1AwsResourceReference `json:"additionalSecurityGroups"`

	// ami
	// Required: true
	Ami *string `json:"ami"`

	// az
	Az string `json:"az,omitempty"`

	// dns name
	DNSName string `json:"dnsName,omitempty"`

	// iam profile
	IamProfile string `json:"iamProfile,omitempty"`

	// instance type
	// Required: true
	InstanceType *string `json:"instanceType"`

	// nics
	Nics []*V1AwsNic `json:"nics"`

	// phase
	Phase string `json:"phase,omitempty"`

	// ssh key name
	SSHKeyName string `json:"sshKeyName,omitempty"`

	// subnet Id
	SubnetID string `json:"subnetId,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// vpc Id
	// Required: true
	VpcID *string `json:"vpcId"`
}

// Validate validates this v1 aws machine spec
func (m *V1AwsMachineSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalSecurityGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
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

func (m *V1AwsMachineSpec) validateAdditionalSecurityGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalSecurityGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalSecurityGroups); i++ {
		if swag.IsZero(m.AdditionalSecurityGroups[i]) { // not required
			continue
		}

		if m.AdditionalSecurityGroups[i] != nil {
			if err := m.AdditionalSecurityGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalSecurityGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalSecurityGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AwsMachineSpec) validateAmi(formats strfmt.Registry) error {

	if err := validate.Required("ami", "body", m.Ami); err != nil {
		return err
	}

	return nil
}

func (m *V1AwsMachineSpec) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instanceType", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

func (m *V1AwsMachineSpec) validateNics(formats strfmt.Registry) error {
	if swag.IsZero(m.Nics) { // not required
		return nil
	}

	for i := 0; i < len(m.Nics); i++ {
		if swag.IsZero(m.Nics[i]) { // not required
			continue
		}

		if m.Nics[i] != nil {
			if err := m.Nics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AwsMachineSpec) validateVpcID(formats strfmt.Registry) error {

	if err := validate.Required("vpcId", "body", m.VpcID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 aws machine spec based on the context it is used
func (m *V1AwsMachineSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalSecurityGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AwsMachineSpec) contextValidateAdditionalSecurityGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalSecurityGroups); i++ {

		if m.AdditionalSecurityGroups[i] != nil {

			if swag.IsZero(m.AdditionalSecurityGroups[i]) { // not required
				return nil
			}

			if err := m.AdditionalSecurityGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalSecurityGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalSecurityGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AwsMachineSpec) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nics); i++ {

		if m.Nics[i] != nil {

			if swag.IsZero(m.Nics[i]) { // not required
				return nil
			}

			if err := m.Nics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AwsMachineSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AwsMachineSpec) UnmarshalBinary(b []byte) error {
	var res V1AwsMachineSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
