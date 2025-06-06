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

// V1PackMetadataSpec Pack metadata spec
//
// swagger:model v1PackMetadataSpec
type V1PackMetadataSpec struct {

	// Pack add-on sub type such as monitoring, db etc
	AddonSubType string `json:"addonSubType,omitempty"`

	// Pack add-on type such as logging, monitoring, security etc
	AddonType string `json:"addonType,omitempty"`

	// Pack supported cloud types
	CloudTypes []string `json:"cloudTypes"`

	// Pack display name
	DisplayName string `json:"displayName,omitempty"`

	// Pack group
	Group string `json:"group,omitempty"`

	// layer
	Layer V1PackLayer `json:"layer,omitempty"`

	// Pack name
	Name string `json:"name,omitempty"`

	// Pack registries array
	Registries []*V1RegistryPackMetadata `json:"registries"`

	// type
	Type *V1PackType `json:"type,omitempty"`
}

// Validate validates this v1 pack metadata spec
func (m *V1PackMetadataSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackMetadataSpec) validateLayer(formats strfmt.Registry) error {
	if swag.IsZero(m.Layer) { // not required
		return nil
	}

	if err := m.Layer.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("layer")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("layer")
		}
		return err
	}

	return nil
}

func (m *V1PackMetadataSpec) validateRegistries(formats strfmt.Registry) error {
	if swag.IsZero(m.Registries) { // not required
		return nil
	}

	for i := 0; i < len(m.Registries); i++ {
		if swag.IsZero(m.Registries[i]) { // not required
			continue
		}

		if m.Registries[i] != nil {
			if err := m.Registries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("registries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PackMetadataSpec) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 pack metadata spec based on the context it is used
func (m *V1PackMetadataSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLayer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackMetadataSpec) contextValidateLayer(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Layer) { // not required
		return nil
	}

	if err := m.Layer.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("layer")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("layer")
		}
		return err
	}

	return nil
}

func (m *V1PackMetadataSpec) contextValidateRegistries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Registries); i++ {

		if m.Registries[i] != nil {

			if swag.IsZero(m.Registries[i]) { // not required
				return nil
			}

			if err := m.Registries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("registries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PackMetadataSpec) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PackMetadataSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackMetadataSpec) UnmarshalBinary(b []byte) error {
	var res V1PackMetadataSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
