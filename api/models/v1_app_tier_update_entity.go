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

// V1AppTierUpdateEntity Application tier update request payload
//
// swagger:model v1AppTierUpdateEntity
type V1AppTierUpdateEntity struct {

	// Application tier container registry uid
	ContainerRegistryUID string `json:"containerRegistryUid,omitempty"`

	// Application tier installation order
	InstallOrder int32 `json:"installOrder,omitempty"`

	// Application tier manifests
	Manifests []*V1ManifestRefUpdateEntity `json:"manifests"`

	// Application tier name
	Name string `json:"name,omitempty"`

	// Application tier properties
	Properties []*V1AppTierPropertyEntity `json:"properties"`

	// Application tier configuration values in yaml format
	Values string `json:"values,omitempty"`

	// Application tier version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 app tier update entity
func (m *V1AppTierUpdateEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManifests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppTierUpdateEntity) validateManifests(formats strfmt.Registry) error {
	if swag.IsZero(m.Manifests) { // not required
		return nil
	}

	for i := 0; i < len(m.Manifests); i++ {
		if swag.IsZero(m.Manifests[i]) { // not required
			continue
		}

		if m.Manifests[i] != nil {
			if err := m.Manifests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("manifests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("manifests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AppTierUpdateEntity) validateProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	for i := 0; i < len(m.Properties); i++ {
		if swag.IsZero(m.Properties[i]) { // not required
			continue
		}

		if m.Properties[i] != nil {
			if err := m.Properties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 app tier update entity based on the context it is used
func (m *V1AppTierUpdateEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateManifests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppTierUpdateEntity) contextValidateManifests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Manifests); i++ {

		if m.Manifests[i] != nil {

			if swag.IsZero(m.Manifests[i]) { // not required
				return nil
			}

			if err := m.Manifests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("manifests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("manifests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AppTierUpdateEntity) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Properties); i++ {

		if m.Properties[i] != nil {

			if swag.IsZero(m.Properties[i]) { // not required
				return nil
			}

			if err := m.Properties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppTierUpdateEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppTierUpdateEntity) UnmarshalBinary(b []byte) error {
	var res V1AppTierUpdateEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
