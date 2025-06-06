// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OverlordMigrateEntity v1 overlord migrate entity
//
// swagger:model v1OverlordMigrateEntity
type V1OverlordMigrateEntity struct {

	// source Uid
	SourceUID string `json:"sourceUid,omitempty"`

	// target Uid
	TargetUID string `json:"targetUid,omitempty"`
}

// Validate validates this v1 overlord migrate entity
func (m *V1OverlordMigrateEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 overlord migrate entity based on context it is used
func (m *V1OverlordMigrateEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1OverlordMigrateEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OverlordMigrateEntity) UnmarshalBinary(b []byte) error {
	var res V1OverlordMigrateEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
