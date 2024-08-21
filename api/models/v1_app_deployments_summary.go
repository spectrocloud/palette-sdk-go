// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1AppDeploymentsSummary v1 app deployments summary
//
// swagger:model v1AppDeploymentsSummary
type V1AppDeploymentsSummary struct {

	// app deployments
	// Unique: true
	AppDeployments []*V1AppDeploymentSummary `json:"appDeployments"`

	// listmeta
	Listmeta *V1ListMetaData `json:"listmeta,omitempty"`
}

// Validate validates this v1 app deployments summary
func (m *V1AppDeploymentsSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppDeployments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListmeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppDeploymentsSummary) validateAppDeployments(formats strfmt.Registry) error {

	if swag.IsZero(m.AppDeployments) { // not required
		return nil
	}

	if err := validate.UniqueItems("appDeployments", "body", m.AppDeployments); err != nil {
		return err
	}

	for i := 0; i < len(m.AppDeployments); i++ {
		if swag.IsZero(m.AppDeployments[i]) { // not required
			continue
		}

		if m.AppDeployments[i] != nil {
			if err := m.AppDeployments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appDeployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AppDeploymentsSummary) validateListmeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Listmeta) { // not required
		return nil
	}

	if m.Listmeta != nil {
		if err := m.Listmeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listmeta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppDeploymentsSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppDeploymentsSummary) UnmarshalBinary(b []byte) error {
	var res V1AppDeploymentsSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}