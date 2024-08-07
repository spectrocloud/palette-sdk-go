// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PackManifestRef v1 pack manifest ref
//
// swagger:model v1PackManifestRef
type V1PackManifestRef struct {

	// digest
	Digest string `json:"digest,omitempty"`

	// is overridden
	IsOverridden bool `json:"isOverridden"`

	// name
	Name string `json:"name,omitempty"`

	// parent Uid
	ParentUID string `json:"parentUid,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 pack manifest ref
func (m *V1PackManifestRef) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PackManifestRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PackManifestRef) UnmarshalBinary(b []byte) error {
	var res V1PackManifestRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
