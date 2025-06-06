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

// V1VMPersistentVolumeClaimSpec PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes
//
// swagger:model v1VmPersistentVolumeClaimSpec
type V1VMPersistentVolumeClaimSpec struct {

	// AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	AccessModes []string `json:"accessModes"`

	// data source
	DataSource *V1VMTypedLocalObjectReference `json:"dataSource,omitempty"`

	// data source ref
	DataSourceRef *V1VMTypedLocalObjectReference `json:"dataSourceRef,omitempty"`

	// resources
	Resources *V1VMCoreResourceRequirements `json:"resources,omitempty"`

	// selector
	Selector *V1VMLabelSelector `json:"selector,omitempty"`

	// Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	StorageClassName string `json:"storageClassName,omitempty"`

	// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
	VolumeMode string `json:"volumeMode,omitempty"`

	// VolumeName is the binding reference to the PersistentVolume backing this claim.
	VolumeName string `json:"volumeName,omitempty"`
}

// Validate validates this v1 Vm persistent volume claim spec
func (m *V1VMPersistentVolumeClaimSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataSourceRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMPersistentVolumeClaimSpec) validateDataSource(formats strfmt.Registry) error {
	if swag.IsZero(m.DataSource) { // not required
		return nil
	}

	if m.DataSource != nil {
		if err := m.DataSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataSource")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMPersistentVolumeClaimSpec) validateDataSourceRef(formats strfmt.Registry) error {
	if swag.IsZero(m.DataSourceRef) { // not required
		return nil
	}

	if m.DataSourceRef != nil {
		if err := m.DataSourceRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataSourceRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataSourceRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMPersistentVolumeClaimSpec) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMPersistentVolumeClaimSpec) validateSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.Selector) { // not required
		return nil
	}

	if m.Selector != nil {
		if err := m.Selector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 Vm persistent volume claim spec based on the context it is used
func (m *V1VMPersistentVolumeClaimSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMPersistentVolumeClaimSpec) contextValidateDataSource(ctx context.Context, formats strfmt.Registry) error {

	if m.DataSource != nil {

		if swag.IsZero(m.DataSource) { // not required
			return nil
		}

		if err := m.DataSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataSource")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMPersistentVolumeClaimSpec) contextValidateDataSourceRef(ctx context.Context, formats strfmt.Registry) error {

	if m.DataSourceRef != nil {

		if swag.IsZero(m.DataSourceRef) { // not required
			return nil
		}

		if err := m.DataSourceRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataSourceRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataSourceRef")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMPersistentVolumeClaimSpec) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if m.Resources != nil {

		if swag.IsZero(m.Resources) { // not required
			return nil
		}

		if err := m.Resources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMPersistentVolumeClaimSpec) contextValidateSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.Selector != nil {

		if swag.IsZero(m.Selector) { // not required
			return nil
		}

		if err := m.Selector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMPersistentVolumeClaimSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMPersistentVolumeClaimSpec) UnmarshalBinary(b []byte) error {
	var res V1VMPersistentVolumeClaimSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
