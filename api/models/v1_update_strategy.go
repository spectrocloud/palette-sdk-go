// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1UpdateStrategy UpdatesStrategy will be used to translate to RollingUpdateStrategy of a MachineDeployment We'll start with default values for the translation, can expose more details later Following is details of parameters translated from the type ScaleOut => maxSurge=1, maxUnavailable=0 ScaleIn => maxSurge=0, maxUnavailable=1
//
// swagger:model v1UpdateStrategy
type V1UpdateStrategy struct {

	// update strategy, either ScaleOut or ScaleIn if empty, will default to RollingUpdateScaleOut
	// Enum: ["RollingUpdateScaleOut","RollingUpdateScaleIn"]
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 update strategy
func (m *V1UpdateStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1UpdateStrategyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RollingUpdateScaleOut","RollingUpdateScaleIn"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1UpdateStrategyTypeTypePropEnum = append(v1UpdateStrategyTypeTypePropEnum, v)
	}
}

const (

	// V1UpdateStrategyTypeRollingUpdateScaleOut captures enum value "RollingUpdateScaleOut"
	V1UpdateStrategyTypeRollingUpdateScaleOut string = "RollingUpdateScaleOut"

	// V1UpdateStrategyTypeRollingUpdateScaleIn captures enum value "RollingUpdateScaleIn"
	V1UpdateStrategyTypeRollingUpdateScaleIn string = "RollingUpdateScaleIn"
)

// prop value enum
func (m *V1UpdateStrategy) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1UpdateStrategyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1UpdateStrategy) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 update strategy based on context it is used
func (m *V1UpdateStrategy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UpdateStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UpdateStrategy) UnmarshalBinary(b []byte) error {
	var res V1UpdateStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
