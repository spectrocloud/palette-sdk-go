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

// V1VMDataVolumeSourcePVC DataVolumeSourcePVC provides the parameters to create a Data Volume from an existing PVC
//
// swagger:model v1VmDataVolumeSourcePVC
type V1VMDataVolumeSourcePVC struct {

	// The name of the source PVC
	// Required: true
	Name *string `json:"name"`

	// The namespace of the source PVC
	// Required: true
	Namespace *string `json:"namespace"`
}

// Validate validates this v1 Vm data volume source p v c
func (m *V1VMDataVolumeSourcePVC) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMDataVolumeSourcePVC) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1VMDataVolumeSourcePVC) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 Vm data volume source p v c based on context it is used
func (m *V1VMDataVolumeSourcePVC) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMDataVolumeSourcePVC) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMDataVolumeSourcePVC) UnmarshalBinary(b []byte) error {
	var res V1VMDataVolumeSourcePVC
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
