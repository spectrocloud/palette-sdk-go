// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1IPPoolEntity IP Pool entity definition
//
// swagger:model v1IpPoolEntity
type V1IPPoolEntity struct {

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// spec
	Spec *V1IPPoolEntitySpec `json:"spec,omitempty"`

	// status
	Status *V1IPPoolStatus `json:"status,omitempty"`
}

// Validate validates this v1 Ip pool entity
func (m *V1IPPoolEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IPPoolEntity) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1IPPoolEntity) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

func (m *V1IPPoolEntity) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IPPoolEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IPPoolEntity) UnmarshalBinary(b []byte) error {
	var res V1IPPoolEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1IPPoolEntitySpec v1 IP pool entity spec
//
// swagger:model V1IPPoolEntitySpec
type V1IPPoolEntitySpec struct {

	// pool
	Pool *V1Pool `json:"pool,omitempty"`

	// priavet gateway Uid
	PriavetGatewayUID string `json:"priavetGatewayUid,omitempty"`

	// if true, restricts this IP pool to be used by single cluster at any time
	RestrictToSingleCluster bool `json:"restrictToSingleCluster"`
}

// Validate validates this v1 IP pool entity spec
func (m *V1IPPoolEntitySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IPPoolEntitySpec) validatePool(formats strfmt.Registry) error {

	if swag.IsZero(m.Pool) { // not required
		return nil
	}

	if m.Pool != nil {
		if err := m.Pool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "pool")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IPPoolEntitySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IPPoolEntitySpec) UnmarshalBinary(b []byte) error {
	var res V1IPPoolEntitySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}