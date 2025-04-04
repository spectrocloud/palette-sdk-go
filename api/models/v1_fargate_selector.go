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

// V1FargateSelector FargateSelector specifies a selector for pods that should run on this fargate pool
//
// swagger:model v1FargateSelector
type V1FargateSelector struct {

	// Labels specifies which pod labels this selector should match.
	Labels map[string]string `json:"labels,omitempty"`

	// Namespace specifies which namespace this selector should match.
	// Required: true
	Namespace *string `json:"namespace"`
}

// Validate validates this v1 fargate selector
func (m *V1FargateSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1FargateSelector) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 fargate selector based on context it is used
func (m *V1FargateSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1FargateSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FargateSelector) UnmarshalBinary(b []byte) error {
	var res V1FargateSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
