// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1EksMachineCloudConfigEntity v1 eks machine cloud config entity
//
// swagger:model v1EksMachineCloudConfigEntity
type V1EksMachineCloudConfigEntity struct {

	// aws launch template
	AwsLaunchTemplate *V1AwsLaunchTemplate `json:"awsLaunchTemplate,omitempty"`

	// azs
	Azs []string `json:"azs"`

	// EC2 instance capacity type
	// Enum: [on-demand spot]
	CapacityType *string `json:"capacityType,omitempty"`

	// flag to know if aws launch template is enabled
	EnableAwsLaunchTemplate bool `json:"enableAwsLaunchTemplate,omitempty"`

	// instance type
	InstanceType string `json:"instanceType,omitempty"`

	// rootDeviceSize in GBs
	// Maximum: 2000
	// Minimum: 1
	RootDeviceSize int64 `json:"rootDeviceSize,omitempty"`

	// SpotMarketOptions allows users to configure instances to be run using AWS Spot instances.
	SpotMarketOptions *V1SpotMarketOptions `json:"spotMarketOptions,omitempty"`

	// subnets
	Subnets []*V1EksSubnetEntity `json:"subnets"`
}

// Validate validates this v1 eks machine cloud config entity
func (m *V1EksMachineCloudConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsLaunchTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootDeviceSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpotMarketOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EksMachineCloudConfigEntity) validateAwsLaunchTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsLaunchTemplate) { // not required
		return nil
	}

	if m.AwsLaunchTemplate != nil {
		if err := m.AwsLaunchTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsLaunchTemplate")
			}
			return err
		}
	}

	return nil
}

var v1EksMachineCloudConfigEntityTypeCapacityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on-demand","spot"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1EksMachineCloudConfigEntityTypeCapacityTypePropEnum = append(v1EksMachineCloudConfigEntityTypeCapacityTypePropEnum, v)
	}
}

const (

	// V1EksMachineCloudConfigEntityCapacityTypeOnDemand captures enum value "on-demand"
	V1EksMachineCloudConfigEntityCapacityTypeOnDemand string = "on-demand"

	// V1EksMachineCloudConfigEntityCapacityTypeSpot captures enum value "spot"
	V1EksMachineCloudConfigEntityCapacityTypeSpot string = "spot"
)

// prop value enum
func (m *V1EksMachineCloudConfigEntity) validateCapacityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1EksMachineCloudConfigEntityTypeCapacityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1EksMachineCloudConfigEntity) validateCapacityType(formats strfmt.Registry) error {

	if swag.IsZero(m.CapacityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCapacityTypeEnum("capacityType", "body", *m.CapacityType); err != nil {
		return err
	}

	return nil
}

func (m *V1EksMachineCloudConfigEntity) validateRootDeviceSize(formats strfmt.Registry) error {

	if swag.IsZero(m.RootDeviceSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("rootDeviceSize", "body", int64(m.RootDeviceSize), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rootDeviceSize", "body", int64(m.RootDeviceSize), 2000, false); err != nil {
		return err
	}

	return nil
}

func (m *V1EksMachineCloudConfigEntity) validateSpotMarketOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.SpotMarketOptions) { // not required
		return nil
	}

	if m.SpotMarketOptions != nil {
		if err := m.SpotMarketOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spotMarketOptions")
			}
			return err
		}
	}

	return nil
}

func (m *V1EksMachineCloudConfigEntity) validateSubnets(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V1EksMachineCloudConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EksMachineCloudConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1EksMachineCloudConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
