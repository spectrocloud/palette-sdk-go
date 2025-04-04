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

// V1PackRegistryImportEntity Pack registry import entity
//
// swagger:model v1PackRegistryImportEntity
type V1PackRegistryImportEntity struct {

	// matching registries
	MatchingRegistries []*V1PackRegistryMetadata `json:"matchingRegistries"`

	// metadata
	Metadata *V1PackRegistryMetadata `json:"metadata,omitempty"`
}

// Validate validates this v1 pack registry import entity
func (m *V1PackRegistryImportEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchingRegistries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackRegistryImportEntity) validateMatchingRegistries(formats strfmt.Registry) error {
	if swag.IsZero(m.MatchingRegistries) { // not required
		return nil
	}

	for i := 0; i < len(m.MatchingRegistries); i++ {
		if swag.IsZero(m.MatchingRegistries[i]) { // not required
			continue
		}

		if m.MatchingRegistries[i] != nil {
			if err := m.MatchingRegistries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchingRegistries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matchingRegistries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PackRegistryImportEntity) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 pack registry import entity based on the context it is used
func (m *V1PackRegistryImportEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatchingRegistries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackRegistryImportEntity) contextValidateMatchingRegistries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MatchingRegistries); i++ {

		if m.MatchingRegistries[i] != nil {

			if swag.IsZero(m.MatchingRegistries[i]) { // not required
				return nil
			}

			if err := m.MatchingRegistries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchingRegistries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matchingRegistries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PackRegistryImportEntity) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PackRegistryImportEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackRegistryImportEntity) UnmarshalBinary(b []byte) error {
	var res V1PackRegistryImportEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
