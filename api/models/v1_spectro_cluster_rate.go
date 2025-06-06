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
)

// V1SpectroClusterRate Cluster estimated rate information
//
// swagger:model v1SpectroClusterRate
type V1SpectroClusterRate struct {

	// machine pools
	MachinePools []*V1MachinePoolRate `json:"machinePools"`

	// name
	Name string `json:"name,omitempty"`

	// rate
	Rate *V1TotalClusterRate `json:"rate,omitempty"`

	// resource metadata
	ResourceMetadata *V1CloudResourceMetadata `json:"resourceMetadata,omitempty"`
}

// Validate validates this v1 spectro cluster rate
func (m *V1SpectroClusterRate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachinePools(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterRate) validateMachinePools(formats strfmt.Registry) error {
	if swag.IsZero(m.MachinePools) { // not required
		return nil
	}

	for i := 0; i < len(m.MachinePools); i++ {
		if swag.IsZero(m.MachinePools[i]) { // not required
			continue
		}

		if m.MachinePools[i] != nil {
			if err := m.MachinePools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machinePools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machinePools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterRate) validateRate(formats strfmt.Registry) error {
	if swag.IsZero(m.Rate) { // not required
		return nil
	}

	if m.Rate != nil {
		if err := m.Rate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rate")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterRate) validateResourceMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceMetadata) { // not required
		return nil
	}

	if m.ResourceMetadata != nil {
		if err := m.ResourceMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceMetadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 spectro cluster rate based on the context it is used
func (m *V1SpectroClusterRate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMachinePools(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterRate) contextValidateMachinePools(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MachinePools); i++ {

		if m.MachinePools[i] != nil {

			if swag.IsZero(m.MachinePools[i]) { // not required
				return nil
			}

			if err := m.MachinePools[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machinePools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machinePools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterRate) contextValidateRate(ctx context.Context, formats strfmt.Registry) error {

	if m.Rate != nil {

		if swag.IsZero(m.Rate) { // not required
			return nil
		}

		if err := m.Rate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rate")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterRate) contextValidateResourceMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceMetadata != nil {

		if swag.IsZero(m.ResourceMetadata) { // not required
			return nil
		}

		if err := m.ResourceMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceMetadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterRate) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
