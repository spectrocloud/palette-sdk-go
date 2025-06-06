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

// V1HTTPPatch v1 Http patch
//
// swagger:model v1HttpPatch
type V1HTTPPatch struct {

	// A path to the pointer from which reference will be taken
	From string `json:"from,omitempty"`

	// The operation to be performed
	// Required: true
	// Enum: ["add","remove","replace","move","copy"]
	Op *string `json:"op"`

	// A path to the pointer on which operation will be done
	// Required: true
	Path *string `json:"path"`

	// The value to be used within the operations.
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this v1 Http patch
func (m *V1HTTPPatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1HttpPatchTypeOpPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["add","remove","replace","move","copy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1HttpPatchTypeOpPropEnum = append(v1HttpPatchTypeOpPropEnum, v)
	}
}

const (

	// V1HTTPPatchOpAdd captures enum value "add"
	V1HTTPPatchOpAdd string = "add"

	// V1HTTPPatchOpRemove captures enum value "remove"
	V1HTTPPatchOpRemove string = "remove"

	// V1HTTPPatchOpReplace captures enum value "replace"
	V1HTTPPatchOpReplace string = "replace"

	// V1HTTPPatchOpMove captures enum value "move"
	V1HTTPPatchOpMove string = "move"

	// V1HTTPPatchOpCopy captures enum value "copy"
	V1HTTPPatchOpCopy string = "copy"
)

// prop value enum
func (m *V1HTTPPatch) validateOpEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1HttpPatchTypeOpPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1HTTPPatch) validateOp(formats strfmt.Registry) error {

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	// value enum
	if err := m.validateOpEnum("op", "body", *m.Op); err != nil {
		return err
	}

	return nil
}

func (m *V1HTTPPatch) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 Http patch based on context it is used
func (m *V1HTTPPatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1HTTPPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1HTTPPatch) UnmarshalBinary(b []byte) error {
	var res V1HTTPPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
