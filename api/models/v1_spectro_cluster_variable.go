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

// V1SpectroClusterVariable Variable with value which will be used within the packs of cluster profile
//
// swagger:model v1SpectroClusterVariable
type V1SpectroClusterVariable struct {

	// Variable name
	// Required: true
	Name *string `json:"name"`

	// Actual value of the variable to be used within the cluster
	Value string `json:"value,omitempty"`
}

// Validate validates this v1 spectro cluster variable
func (m *V1SpectroClusterVariable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterVariable) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 spectro cluster variable based on context it is used
func (m *V1SpectroClusterVariable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterVariable) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
