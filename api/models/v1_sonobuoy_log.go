// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SonobuoyLog Compliance Scan Sonobuoy Log
//
// swagger:model v1SonobuoyLog
type V1SonobuoyLog struct {

	// description
	Description string `json:"description,omitempty"`

	// msg
	Msg string `json:"msg,omitempty"`

	// output
	Output string `json:"output,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 sonobuoy log
func (m *V1SonobuoyLog) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SonobuoyLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SonobuoyLog) UnmarshalBinary(b []byte) error {
	var res V1SonobuoyLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}