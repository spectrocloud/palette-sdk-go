// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroClusterState Spectrocluster state entity
//
// swagger:model v1SpectroClusterState
type V1SpectroClusterState struct {

	// Spectrocluster state
	State string `json:"state,omitempty"`
}

// Validate validates this v1 spectro cluster state
func (m *V1SpectroClusterState) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterState) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}