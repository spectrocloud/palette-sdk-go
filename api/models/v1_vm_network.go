// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1VMNetwork Network represents a network type and a resource that should be connected to the vm.
//
// swagger:model v1VmNetwork
type V1VMNetwork struct {

	// multus
	Multus *V1VMMultusNetwork `json:"multus,omitempty"`

	// Network name. Must be a DNS_LABEL and unique within the vm. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// Required: true
	Name *string `json:"name"`

	// pod
	Pod *V1VMPodNetwork `json:"pod,omitempty"`
}

// Validate validates this v1 Vm network
func (m *V1VMNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMultus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMNetwork) validateMultus(formats strfmt.Registry) error {
	if swag.IsZero(m.Multus) { // not required
		return nil
	}

	if m.Multus != nil {
		if err := m.Multus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("multus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("multus")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMNetwork) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1VMNetwork) validatePod(formats strfmt.Registry) error {
	if swag.IsZero(m.Pod) { // not required
		return nil
	}

	if m.Pod != nil {
		if err := m.Pod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pod")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 Vm network based on the context it is used
func (m *V1VMNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMultus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMNetwork) contextValidateMultus(ctx context.Context, formats strfmt.Registry) error {

	if m.Multus != nil {

		if swag.IsZero(m.Multus) { // not required
			return nil
		}

		if err := m.Multus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("multus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("multus")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMNetwork) contextValidatePod(ctx context.Context, formats strfmt.Registry) error {

	if m.Pod != nil {

		if swag.IsZero(m.Pod) { // not required
			return nil
		}

		if err := m.Pod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMNetwork) UnmarshalBinary(b []byte) error {
	var res V1VMNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
