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

// V1CoxEdgeBaseUrls List of CoxEdge base urls
//
// swagger:model v1CoxEdgeBaseUrls
type V1CoxEdgeBaseUrls struct {

	// base urls
	// Required: true
	BaseUrls []string `json:"baseUrls"`
}

// Validate validates this v1 cox edge base urls
func (m *V1CoxEdgeBaseUrls) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseUrls(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CoxEdgeBaseUrls) validateBaseUrls(formats strfmt.Registry) error {

	if err := validate.Required("baseUrls", "body", m.BaseUrls); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CoxEdgeBaseUrls) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CoxEdgeBaseUrls) UnmarshalBinary(b []byte) error {
	var res V1CoxEdgeBaseUrls
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
