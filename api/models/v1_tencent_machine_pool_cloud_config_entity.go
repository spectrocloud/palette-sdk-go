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

// V1TencentMachinePoolCloudConfigEntity v1 tencent machine pool cloud config entity
//
// swagger:model v1TencentMachinePoolCloudConfigEntity
type V1TencentMachinePoolCloudConfigEntity struct {

	// azs
	Azs []string `json:"azs"`

	// instance type
	InstanceType string `json:"instanceType,omitempty"`

	// rootDeviceSize in GBs
	// Maximum: 2000
	// Minimum: 1
	RootDeviceSize int64 `json:"rootDeviceSize,omitempty"`

	// AZ to subnet mapping filled by ally from hubble SubnetIDs ["ap-guangzhou-6"] = "subnet-079b6061" This field is optional If we don't provide a subnetId then by default the first subnet from the AZ will be picked up for deployment
	SubnetIds map[string]string `json:"subnetIds,omitempty"`
}

// Validate validates this v1 tencent machine pool cloud config entity
func (m *V1TencentMachinePoolCloudConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRootDeviceSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TencentMachinePoolCloudConfigEntity) validateRootDeviceSize(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V1TencentMachinePoolCloudConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TencentMachinePoolCloudConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1TencentMachinePoolCloudConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
