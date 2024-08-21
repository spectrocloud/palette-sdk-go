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

// V1ClusterBackupLocationType Cluster backup location type
//
// swagger:model v1ClusterBackupLocationType
type V1ClusterBackupLocationType struct {

	// location type
	// Required: true
	LocationType *string `json:"locationType"`
}

// Validate validates this v1 cluster backup location type
func (m *V1ClusterBackupLocationType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterBackupLocationType) validateLocationType(formats strfmt.Registry) error {

	if err := validate.Required("locationType", "body", m.LocationType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterBackupLocationType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterBackupLocationType) UnmarshalBinary(b []byte) error {
	var res V1ClusterBackupLocationType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}