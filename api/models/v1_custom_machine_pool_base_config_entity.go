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

// V1CustomMachinePoolBaseConfigEntity Machine pool configuration for the custom cluster
//
// swagger:model v1CustomMachinePoolBaseConfigEntity
type V1CustomMachinePoolBaseConfigEntity struct {

	// Additional labels to be part of the machine pool
	AdditionalLabels map[string]string `json:"additionalLabels,omitempty"`

	// Whether this pool is for control plane
	IsControlPlane bool `json:"isControlPlane"`

	// control plane or worker taints
	// Unique: true
	Taints []*V1Taint `json:"taints"`

	// If IsControlPlane==true && useControlPlaneAsWorker==true, then will remove control plane taint this will not be used for worker pools
	UseControlPlaneAsWorker bool `json:"useControlPlaneAsWorker"`
}

// Validate validates this v1 custom machine pool base config entity
func (m *V1CustomMachinePoolBaseConfigEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CustomMachinePoolBaseConfigEntity) validateTaints(formats strfmt.Registry) error {
	if swag.IsZero(m.Taints) { // not required
		return nil
	}

	if err := validate.UniqueItems("taints", "body", m.Taints); err != nil {
		return err
	}

	for i := 0; i < len(m.Taints); i++ {
		if swag.IsZero(m.Taints[i]) { // not required
			continue
		}

		if m.Taints[i] != nil {
			if err := m.Taints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 custom machine pool base config entity based on the context it is used
func (m *V1CustomMachinePoolBaseConfigEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CustomMachinePoolBaseConfigEntity) contextValidateTaints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Taints); i++ {

		if m.Taints[i] != nil {

			if swag.IsZero(m.Taints[i]) { // not required
				return nil
			}

			if err := m.Taints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CustomMachinePoolBaseConfigEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CustomMachinePoolBaseConfigEntity) UnmarshalBinary(b []byte) error {
	var res V1CustomMachinePoolBaseConfigEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
