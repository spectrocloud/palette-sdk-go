// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ResourceRolesUpdateEntity v1 resource roles update entity
//
// swagger:model v1ResourceRolesUpdateEntity
type V1ResourceRolesUpdateEntity struct {

	// filter refs
	FilterRefs []string `json:"filterRefs"`

	// project uids
	ProjectUids []string `json:"projectUids"`

	// roles
	Roles []string `json:"roles"`
}

// Validate validates this v1 resource roles update entity
func (m *V1ResourceRolesUpdateEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ResourceRolesUpdateEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ResourceRolesUpdateEntity) UnmarshalBinary(b []byte) error {
	var res V1ResourceRolesUpdateEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
