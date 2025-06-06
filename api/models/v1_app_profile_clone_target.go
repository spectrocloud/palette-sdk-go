// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AppProfileCloneTarget Application profile clone target
//
// swagger:model v1AppProfileCloneTarget
type V1AppProfileCloneTarget struct {

	// Application profile clone target project uid
	ProjectUID string `json:"projectUid,omitempty"`
}

// Validate validates this v1 app profile clone target
func (m *V1AppProfileCloneTarget) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 app profile clone target based on context it is used
func (m *V1AppProfileCloneTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AppProfileCloneTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppProfileCloneTarget) UnmarshalBinary(b []byte) error {
	var res V1AppProfileCloneTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
