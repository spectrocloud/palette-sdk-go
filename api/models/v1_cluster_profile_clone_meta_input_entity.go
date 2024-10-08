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

// V1ClusterProfileCloneMetaInputEntity Cluster profile clone metadata
//
// swagger:model v1ClusterProfileCloneMetaInputEntity
type V1ClusterProfileCloneMetaInputEntity struct {

	// Cloned cluster profile name
	// Required: true
	Name *string `json:"name"`

	// target
	Target *V1ClusterProfileCloneTarget `json:"target,omitempty"`

	// Cloned cluster profile version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 cluster profile clone meta input entity
func (m *V1ClusterProfileCloneMetaInputEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterProfileCloneMetaInputEntity) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterProfileCloneMetaInputEntity) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterProfileCloneMetaInputEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterProfileCloneMetaInputEntity) UnmarshalBinary(b []byte) error {
	var res V1ClusterProfileCloneMetaInputEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
