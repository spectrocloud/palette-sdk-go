// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1BulkDeleteRequest v1 bulk delete request
//
// swagger:model v1BulkDeleteRequest
type V1BulkDeleteRequest struct {

	// uids
	// Required: true
	// Unique: true
	Uids []string `json:"uids"`
}

// Validate validates this v1 bulk delete request
func (m *V1BulkDeleteRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUids(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BulkDeleteRequest) validateUids(formats strfmt.Registry) error {

	if err := validate.Required("uids", "body", m.Uids); err != nil {
		return err
	}

	if err := validate.UniqueItems("uids", "body", m.Uids); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BulkDeleteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BulkDeleteRequest) UnmarshalBinary(b []byte) error {
	var res V1BulkDeleteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
