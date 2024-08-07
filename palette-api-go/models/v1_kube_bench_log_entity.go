// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1KubeBenchLogEntity KubeBench log
//
// swagger:model v1KubeBenchLogEntity
type V1KubeBenchLogEntity struct {

	// description
	Description string `json:"description,omitempty"`

	// expected
	Expected string `json:"expected,omitempty"`

	// remediation
	Remediation string `json:"remediation,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// test Id
	TestID string `json:"testId,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this v1 kube bench log entity
func (m *V1KubeBenchLogEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1KubeBenchLogEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1KubeBenchLogEntity) UnmarshalBinary(b []byte) error {
	var res V1KubeBenchLogEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
