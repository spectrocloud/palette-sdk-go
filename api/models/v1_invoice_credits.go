// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1InvoiceCredits Invoice credits object
//
// swagger:model v1InvoiceCredits
type V1InvoiceCredits struct {

	// Credits allocated for import clusters
	AlloyFreeCredits int64 `json:"alloyFreeCredits,omitempty"`

	// Credits allocated for managed clusters
	PureFreeCredits int64 `json:"pureFreeCredits,omitempty"`
}

// Validate validates this v1 invoice credits
func (m *V1InvoiceCredits) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 invoice credits based on context it is used
func (m *V1InvoiceCredits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1InvoiceCredits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1InvoiceCredits) UnmarshalBinary(b []byte) error {
	var res V1InvoiceCredits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
