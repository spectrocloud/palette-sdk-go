// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1IPPoolStatus IP Pool status
//
// swagger:model v1IpPoolStatus
type V1IPPoolStatus struct {

	// allotted ips
	// Unique: true
	AllottedIps []string `json:"allottedIps"`

	// associated clusters
	// Unique: true
	AssociatedClusters []string `json:"associatedClusters"`

	// in use
	InUse bool `json:"inUse"`
}

// Validate validates this v1 Ip pool status
func (m *V1IPPoolStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllottedIps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssociatedClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IPPoolStatus) validateAllottedIps(formats strfmt.Registry) error {
	if swag.IsZero(m.AllottedIps) { // not required
		return nil
	}

	if err := validate.UniqueItems("allottedIps", "body", m.AllottedIps); err != nil {
		return err
	}

	return nil
}

func (m *V1IPPoolStatus) validateAssociatedClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.AssociatedClusters) { // not required
		return nil
	}

	if err := validate.UniqueItems("associatedClusters", "body", m.AssociatedClusters); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 Ip pool status based on context it is used
func (m *V1IPPoolStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1IPPoolStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IPPoolStatus) UnmarshalBinary(b []byte) error {
	var res V1IPPoolStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
