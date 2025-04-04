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

// V1WorkspaceNamespaceImage Workspace namespace image information
//
// swagger:model v1WorkspaceNamespaceImage
type V1WorkspaceNamespaceImage struct {

	// black listed images
	// Unique: true
	BlackListedImages []string `json:"blackListedImages"`
}

// Validate validates this v1 workspace namespace image
func (m *V1WorkspaceNamespaceImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlackListedImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1WorkspaceNamespaceImage) validateBlackListedImages(formats strfmt.Registry) error {
	if swag.IsZero(m.BlackListedImages) { // not required
		return nil
	}

	if err := validate.UniqueItems("blackListedImages", "body", m.BlackListedImages); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 workspace namespace image based on context it is used
func (m *V1WorkspaceNamespaceImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1WorkspaceNamespaceImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1WorkspaceNamespaceImage) UnmarshalBinary(b []byte) error {
	var res V1WorkspaceNamespaceImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
