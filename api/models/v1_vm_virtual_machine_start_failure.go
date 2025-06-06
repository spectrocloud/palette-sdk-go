// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMVirtualMachineStartFailure VirtualMachineStartFailure tracks VMIs which failed to transition successfully to running using the VM status
//
// swagger:model v1VmVirtualMachineStartFailure
type V1VMVirtualMachineStartFailure struct {

	// consecutive fail count
	ConsecutiveFailCount int32 `json:"consecutiveFailCount,omitempty"`

	// last failed VM i UID
	LastFailedVMIUID string `json:"lastFailedVMIUID,omitempty"`

	// retry after timestamp
	// Format: date-time
	RetryAfterTimestamp V1Time `json:"retryAfterTimestamp,omitempty"`
}

// Validate validates this v1 Vm virtual machine start failure
func (m *V1VMVirtualMachineStartFailure) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRetryAfterTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMVirtualMachineStartFailure) validateRetryAfterTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.RetryAfterTimestamp) { // not required
		return nil
	}

	if err := m.RetryAfterTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("retryAfterTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("retryAfterTimestamp")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 Vm virtual machine start failure based on the context it is used
func (m *V1VMVirtualMachineStartFailure) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRetryAfterTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMVirtualMachineStartFailure) contextValidateRetryAfterTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.RetryAfterTimestamp) { // not required
		return nil
	}

	if err := m.RetryAfterTimestamp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("retryAfterTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("retryAfterTimestamp")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMVirtualMachineStartFailure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMVirtualMachineStartFailure) UnmarshalBinary(b []byte) error {
	var res V1VMVirtualMachineStartFailure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
