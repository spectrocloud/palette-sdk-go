// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1AuditMsgUpdate Audit user message update request payload
//
// swagger:model v1AuditMsgUpdate
type V1AuditMsgUpdate struct {

	// User message
	// Max Length: 255
	// Min Length: 3
	UserMsg string `json:"userMsg,omitempty"`
}

// Validate validates this v1 audit msg update
func (m *V1AuditMsgUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserMsg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AuditMsgUpdate) validateUserMsg(formats strfmt.Registry) error {

	if swag.IsZero(m.UserMsg) { // not required
		return nil
	}

	if err := validate.MinLength("userMsg", "body", string(m.UserMsg), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("userMsg", "body", string(m.UserMsg), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AuditMsgUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AuditMsgUpdate) UnmarshalBinary(b []byte) error {
	var res V1AuditMsgUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}