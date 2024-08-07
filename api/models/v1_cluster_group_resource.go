// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterGroupResource Cluster group resource allocated and usage information
//
// swagger:model v1ClusterGroupResource
type V1ClusterGroupResource struct {

	// allocated
	Allocated float64 `json:"allocated"`

	// used
	Used float64 `json:"used"`
}

// Validate validates this v1 cluster group resource
func (m *V1ClusterGroupResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterGroupResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterGroupResource) UnmarshalBinary(b []byte) error {
	var res V1ClusterGroupResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
