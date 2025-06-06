// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1PackParamsEntity Pack params request payload
//
// swagger:model v1PackParamsEntity
type V1PackParamsEntity struct {

	// references
	// Unique: true
	References []string `json:"references"`
}

// Validate validates this v1 pack params entity
func (m *V1PackParamsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReferences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackParamsEntity) validateReferences(formats strfmt.Registry) error {
	if swag.IsZero(m.References) { // not required
		return nil
	}

	if err := validate.UniqueItems("references", "body", m.References); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 pack params entity based on context it is used
func (m *V1PackParamsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PackParamsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackParamsEntity) UnmarshalBinary(b []byte) error {
	var res V1PackParamsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
