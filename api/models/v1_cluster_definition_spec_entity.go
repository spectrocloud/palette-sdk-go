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

// V1ClusterDefinitionSpecEntity Cluster definition spec entity
//
// swagger:model v1ClusterDefinitionSpecEntity
type V1ClusterDefinitionSpecEntity struct {

	// cloud type
	// Required: true
	CloudType *string `json:"cloudType"`

	// Cluster definition profiles
	// Required: true
	// Unique: true
	Profiles []*V1ClusterDefinitionProfileEntity `json:"profiles"`
}

// Validate validates this v1 cluster definition spec entity
func (m *V1ClusterDefinitionSpecEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterDefinitionSpecEntity) validateCloudType(formats strfmt.Registry) error {

	if err := validate.Required("cloudType", "body", m.CloudType); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterDefinitionSpecEntity) validateProfiles(formats strfmt.Registry) error {

	if err := validate.Required("profiles", "body", m.Profiles); err != nil {
		return err
	}

	if err := validate.UniqueItems("profiles", "body", m.Profiles); err != nil {
		return err
	}

	for i := 0; i < len(m.Profiles); i++ {
		if swag.IsZero(m.Profiles[i]) { // not required
			continue
		}

		if m.Profiles[i] != nil {
			if err := m.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("profiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 cluster definition spec entity based on the context it is used
func (m *V1ClusterDefinitionSpecEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProfiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterDefinitionSpecEntity) contextValidateProfiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Profiles); i++ {

		if m.Profiles[i] != nil {

			if swag.IsZero(m.Profiles[i]) { // not required
				return nil
			}

			if err := m.Profiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("profiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterDefinitionSpecEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterDefinitionSpecEntity) UnmarshalBinary(b []byte) error {
	var res V1ClusterDefinitionSpecEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
