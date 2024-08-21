// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UpdatedMsg Update response with message
//
// swagger:model v1UpdatedMsg
type V1UpdatedMsg struct {

	// msg
	Msg string `json:"msg,omitempty"`
}

// Validate validates this v1 updated msg
func (m *V1UpdatedMsg) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UpdatedMsg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UpdatedMsg) UnmarshalBinary(b []byte) error {
	var res V1UpdatedMsg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}