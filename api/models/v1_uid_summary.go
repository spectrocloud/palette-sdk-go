// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UIDSummary v1 Uid summary
//
// swagger:model v1UidSummary
type V1UIDSummary struct {

	// name
	Name string `json:"name,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 Uid summary
func (m *V1UIDSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 Uid summary based on context it is used
func (m *V1UIDSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UIDSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UIDSummary) UnmarshalBinary(b []byte) error {
	var res V1UIDSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
