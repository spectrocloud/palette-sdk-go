// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1CloudMachineStatus cloud machine status
//
// swagger:model v1CloudMachineStatus
type V1CloudMachineStatus struct {

	// health
	Health *V1MachineHealth `json:"health,omitempty"`

	// instance state
	// Enum: [Pending Provisioning Provisioned Running Deleting Deleted Failed Unknown]
	InstanceState string `json:"instanceState,omitempty"`

	// maintenance status
	MaintenanceStatus *V1MachineMaintenanceStatus `json:"maintenanceStatus,omitempty"`
}

// Validate validates this v1 cloud machine status
func (m *V1CloudMachineStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenanceStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CloudMachineStatus) validateHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

var v1CloudMachineStatusTypeInstanceStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Provisioning","Provisioned","Running","Deleting","Deleted","Failed","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1CloudMachineStatusTypeInstanceStatePropEnum = append(v1CloudMachineStatusTypeInstanceStatePropEnum, v)
	}
}

const (

	// V1CloudMachineStatusInstanceStatePending captures enum value "Pending"
	V1CloudMachineStatusInstanceStatePending string = "Pending"

	// V1CloudMachineStatusInstanceStateProvisioning captures enum value "Provisioning"
	V1CloudMachineStatusInstanceStateProvisioning string = "Provisioning"

	// V1CloudMachineStatusInstanceStateProvisioned captures enum value "Provisioned"
	V1CloudMachineStatusInstanceStateProvisioned string = "Provisioned"

	// V1CloudMachineStatusInstanceStateRunning captures enum value "Running"
	V1CloudMachineStatusInstanceStateRunning string = "Running"

	// V1CloudMachineStatusInstanceStateDeleting captures enum value "Deleting"
	V1CloudMachineStatusInstanceStateDeleting string = "Deleting"

	// V1CloudMachineStatusInstanceStateDeleted captures enum value "Deleted"
	V1CloudMachineStatusInstanceStateDeleted string = "Deleted"

	// V1CloudMachineStatusInstanceStateFailed captures enum value "Failed"
	V1CloudMachineStatusInstanceStateFailed string = "Failed"

	// V1CloudMachineStatusInstanceStateUnknown captures enum value "Unknown"
	V1CloudMachineStatusInstanceStateUnknown string = "Unknown"
)

// prop value enum
func (m *V1CloudMachineStatus) validateInstanceStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1CloudMachineStatusTypeInstanceStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1CloudMachineStatus) validateInstanceState(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceState) { // not required
		return nil
	}

	// value enum
	if err := m.validateInstanceStateEnum("instanceState", "body", m.InstanceState); err != nil {
		return err
	}

	return nil
}

func (m *V1CloudMachineStatus) validateMaintenanceStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.MaintenanceStatus) { // not required
		return nil
	}

	if m.MaintenanceStatus != nil {
		if err := m.MaintenanceStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenanceStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CloudMachineStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CloudMachineStatus) UnmarshalBinary(b []byte) error {
	var res V1CloudMachineStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
