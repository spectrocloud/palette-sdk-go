// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterProfileMetadataImportEntity Cluster profile import metadata
//
// swagger:model v1ClusterProfileMetadataImportEntity
type V1ClusterProfileMetadataImportEntity struct {

	// Cluster profile description
	Description string `json:"description,omitempty"`

	// Cluster profile labels
	Labels map[string]string `json:"labels,omitempty"`

	// Cluster profile name
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 cluster profile metadata import entity
func (m *V1ClusterProfileMetadataImportEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterProfileMetadataImportEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterProfileMetadataImportEntity) UnmarshalBinary(b []byte) error {
	var res V1ClusterProfileMetadataImportEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
