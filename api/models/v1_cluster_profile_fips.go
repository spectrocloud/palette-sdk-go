// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterProfileFips Cluster profile fips compliance status
//
// swagger:model v1ClusterProfileFips
type V1ClusterProfileFips struct {

	// mode
	Mode V1ClusterFipsMode `json:"mode,omitempty"`
}

// Validate validates this v1 cluster profile fips
func (m *V1ClusterProfileFips) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterProfileFips) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := m.Mode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterProfileFips) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterProfileFips) UnmarshalBinary(b []byte) error {
	var res V1ClusterProfileFips
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}