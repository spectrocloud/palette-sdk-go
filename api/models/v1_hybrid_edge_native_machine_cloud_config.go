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

// V1HybridEdgeNativeMachineCloudConfig HybridEdgeNativeMachineCloudConfig defines hybrid Edge-Native cluster's machine configurations
//
// swagger:model v1HybridEdgeNativeMachineCloudConfig
type V1HybridEdgeNativeMachineCloudConfig struct {

	// Architecture type of the edge hosts
	// Required: true
	ArchType *V1ArchType `json:"archType"`

	// Edge host's configuration
	// Required: true
	EdgeHosts []*V1EdgeNativeHybridMachinePoolHost `json:"edgeHosts"`

	// Hybrid cluster reference
	HybridCluster *V1HybridCluster `json:"hybridCluster,omitempty"`
}

// Validate validates this v1 hybrid edge native machine cloud config
func (m *V1HybridEdgeNativeMachineCloudConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHybridCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1HybridEdgeNativeMachineCloudConfig) validateArchType(formats strfmt.Registry) error {

	if err := validate.Required("archType", "body", m.ArchType); err != nil {
		return err
	}

	if err := validate.Required("archType", "body", m.ArchType); err != nil {
		return err
	}

	if m.ArchType != nil {
		if err := m.ArchType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archType")
			}
			return err
		}
	}

	return nil
}

func (m *V1HybridEdgeNativeMachineCloudConfig) validateEdgeHosts(formats strfmt.Registry) error {

	if err := validate.Required("edgeHosts", "body", m.EdgeHosts); err != nil {
		return err
	}

	for i := 0; i < len(m.EdgeHosts); i++ {
		if swag.IsZero(m.EdgeHosts[i]) { // not required
			continue
		}

		if m.EdgeHosts[i] != nil {
			if err := m.EdgeHosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeHosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edgeHosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1HybridEdgeNativeMachineCloudConfig) validateHybridCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.HybridCluster) { // not required
		return nil
	}

	if m.HybridCluster != nil {
		if err := m.HybridCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hybridCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hybridCluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 hybrid edge native machine cloud config based on the context it is used
func (m *V1HybridEdgeNativeMachineCloudConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArchType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHybridCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1HybridEdgeNativeMachineCloudConfig) contextValidateArchType(ctx context.Context, formats strfmt.Registry) error {

	if m.ArchType != nil {

		if err := m.ArchType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archType")
			}
			return err
		}
	}

	return nil
}

func (m *V1HybridEdgeNativeMachineCloudConfig) contextValidateEdgeHosts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EdgeHosts); i++ {

		if m.EdgeHosts[i] != nil {

			if swag.IsZero(m.EdgeHosts[i]) { // not required
				return nil
			}

			if err := m.EdgeHosts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeHosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edgeHosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1HybridEdgeNativeMachineCloudConfig) contextValidateHybridCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.HybridCluster != nil {

		if swag.IsZero(m.HybridCluster) { // not required
			return nil
		}

		if err := m.HybridCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hybridCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hybridCluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1HybridEdgeNativeMachineCloudConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1HybridEdgeNativeMachineCloudConfig) UnmarshalBinary(b []byte) error {
	var res V1HybridEdgeNativeMachineCloudConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
