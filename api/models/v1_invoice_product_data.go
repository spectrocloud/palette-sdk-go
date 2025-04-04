// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1InvoiceProductData Product invoice data
//
// swagger:model v1InvoiceProductData
type V1InvoiceProductData struct {

	// Allocated credits
	AllocatedCredits int64 `json:"allocatedCredits,omitempty"`

	// Total amount
	Amount float64 `json:"amount,omitempty"`

	// Credits to be billed
	BillableCredits float64 `json:"billableCredits,omitempty"`

	// Credits that are exceeds the allocated credits
	BreachedCredits float64 `json:"breachedCredits,omitempty"`

	// Applied discount
	Discount int64 `json:"discount,omitempty"`

	// Allocated free credits
	FreeCredits int64 `json:"freeCredits,omitempty"`

	// Allowed overage limit in percentage
	OverageLimitPercentage int8 `json:"overageLimitPercentage,omitempty"`

	// Tier name
	TierName string `json:"tierName,omitempty"`

	// Tier price
	TierPrice float64 `json:"tierPrice,omitempty"`

	// Total used credits
	TotalUsedCredits float64 `json:"totalUsedCredits,omitempty"`

	// Used credits
	UsedCredits float64 `json:"usedCredits,omitempty"`
}

// Validate validates this v1 invoice product data
func (m *V1InvoiceProductData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 invoice product data based on context it is used
func (m *V1InvoiceProductData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1InvoiceProductData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1InvoiceProductData) UnmarshalBinary(b []byte) error {
	var res V1InvoiceProductData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
