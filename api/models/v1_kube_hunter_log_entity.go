// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1KubeHunterLogEntity KubeHunter log
//
// swagger:model v1KubeHunterLogEntity
type V1KubeHunterLogEntity struct {

	// description
	Description string `json:"description,omitempty"`

	// evidence
	Evidence string `json:"evidence,omitempty"`

	// reference
	Reference string `json:"reference,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`

	// test Id
	TestID string `json:"testId,omitempty"`

	// vulnerability
	Vulnerability string `json:"vulnerability,omitempty"`
}

// Validate validates this v1 kube hunter log entity
func (m *V1KubeHunterLogEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1KubeHunterLogEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1KubeHunterLogEntity) UnmarshalBinary(b []byte) error {
	var res V1KubeHunterLogEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
