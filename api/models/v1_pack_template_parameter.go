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

// V1PackTemplateParameter Pack template parameter
//
// swagger:model v1PackTemplateParameter
type V1PackTemplateParameter struct {

	// Pack template parameter description
	Description string `json:"description,omitempty"`

	// Pack template parameter display name
	DisplayName string `json:"displayName,omitempty"`

	// Pack template parameter format
	Format string `json:"format,omitempty"`

	// Pack template parameter hidden flag, if true then the parameter is hidden in the UI
	Hidden bool `json:"hidden,omitempty"`

	// Pack template parameter list options as string array
	ListOptions []string `json:"listOptions"`

	// Pack template parameter name
	Name string `json:"name,omitempty"`

	// Pack template parameter optional flag, if true then the parameter value is not mandatory
	Optional bool `json:"optional,omitempty"`

	// Pack template parameter options array
	Options map[string]V1PackTemplateParameterOption `json:"options,omitempty"`

	// Pack template parameter readonly flag, if true then the parameter value can't be overridden
	ReadOnly bool `json:"readOnly,omitempty"`

	// Pack template parameter regex, if set then parameter value must match with specified regex
	Regex string `json:"regex,omitempty"`

	// Pack template parameter target key which is mapped to the key defined in the pack values
	TargetKey string `json:"targetKey,omitempty"`

	// Pack template parameter data type
	Type string `json:"type,omitempty"`

	// Pack template parameter value
	Value string `json:"value,omitempty"`
}

// Validate validates this v1 pack template parameter
func (m *V1PackTemplateParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackTemplateParameter) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	for k := range m.Options {

		if err := validate.Required("options"+"."+k, "body", m.Options[k]); err != nil {
			return err
		}
		if val, ok := m.Options[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("options" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 pack template parameter based on the context it is used
func (m *V1PackTemplateParameter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PackTemplateParameter) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Options {

		if val, ok := m.Options[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PackTemplateParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackTemplateParameter) UnmarshalBinary(b []byte) error {
	var res V1PackTemplateParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
