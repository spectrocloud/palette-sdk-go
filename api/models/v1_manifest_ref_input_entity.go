// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ManifestRefInputEntity Manifest request payload
//
// swagger:model v1ManifestRefInputEntity
type V1ManifestRefInputEntity struct {

	// Manifest content in yaml
	Content string `json:"content,omitempty"`

	// Manifest uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 manifest ref input entity
func (m *V1ManifestRefInputEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ManifestRefInputEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ManifestRefInputEntity) UnmarshalBinary(b []byte) error {
	var res V1ManifestRefInputEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
