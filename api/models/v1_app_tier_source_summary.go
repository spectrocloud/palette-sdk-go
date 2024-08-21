// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AppTierSourceSummary Application profile's tier source information
//
// swagger:model v1AppTierSourceSummary
type V1AppTierSourceSummary struct {

	// addon sub type
	AddonSubType string `json:"addonSubType,omitempty"`

	// addon type
	AddonType string `json:"addonType,omitempty"`

	// logo Url
	LogoURL string `json:"logoUrl,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 app tier source summary
func (m *V1AppTierSourceSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AppTierSourceSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppTierSourceSummary) UnmarshalBinary(b []byte) error {
	var res V1AppTierSourceSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}