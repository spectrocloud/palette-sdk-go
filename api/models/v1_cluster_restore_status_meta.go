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

// V1ClusterRestoreStatusMeta Cluster Restore Status Meta
//
// swagger:model v1ClusterRestoreStatusMeta
type V1ClusterRestoreStatusMeta struct {

	// actor
	Actor *V1ClusterFeatureActor `json:"actor,omitempty"`

	// backup name
	BackupName string `json:"backupName,omitempty"`

	// backup request Uid
	BackupRequestUID string `json:"backupRequestUid,omitempty"`

	// restore request Uid
	RestoreRequestUID string `json:"restoreRequestUid,omitempty"`

	// restore status meta
	RestoreStatusMeta *V1RestoreStatusMeta `json:"restoreStatusMeta,omitempty"`

	// source cluster ref
	SourceClusterRef *V1ResourceReference `json:"sourceClusterRef,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 cluster restore status meta
func (m *V1ClusterRestoreStatusMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoreStatusMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceClusterRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterRestoreStatusMeta) validateActor(formats strfmt.Registry) error {
	if swag.IsZero(m.Actor) { // not required
		return nil
	}

	if m.Actor != nil {
		if err := m.Actor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actor")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterRestoreStatusMeta) validateRestoreStatusMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.RestoreStatusMeta) { // not required
		return nil
	}

	if m.RestoreStatusMeta != nil {
		if err := m.RestoreStatusMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoreStatusMeta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoreStatusMeta")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterRestoreStatusMeta) validateSourceClusterRef(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceClusterRef) { // not required
		return nil
	}

	if m.SourceClusterRef != nil {
		if err := m.SourceClusterRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceClusterRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourceClusterRef")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 cluster restore status meta based on the context it is used
func (m *V1ClusterRestoreStatusMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestoreStatusMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceClusterRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterRestoreStatusMeta) contextValidateActor(ctx context.Context, formats strfmt.Registry) error {

	if m.Actor != nil {

		if swag.IsZero(m.Actor) { // not required
			return nil
		}

		if err := m.Actor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actor")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterRestoreStatusMeta) contextValidateRestoreStatusMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.RestoreStatusMeta != nil {

		if swag.IsZero(m.RestoreStatusMeta) { // not required
			return nil
		}

		if err := m.RestoreStatusMeta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoreStatusMeta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoreStatusMeta")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterRestoreStatusMeta) contextValidateSourceClusterRef(ctx context.Context, formats strfmt.Registry) error {

	if m.SourceClusterRef != nil {

		if swag.IsZero(m.SourceClusterRef) { // not required
			return nil
		}

		if err := m.SourceClusterRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceClusterRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourceClusterRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterRestoreStatusMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterRestoreStatusMeta) UnmarshalBinary(b []byte) error {
	var res V1ClusterRestoreStatusMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
