// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SpectroClusterPackStatusEntity v1 spectro cluster pack status entity
//
// swagger:model v1SpectroClusterPackStatusEntity
type V1SpectroClusterPackStatusEntity struct {

	// Pack deployment status conditions
	Condition *V1SpectroClusterPackCondition `json:"condition,omitempty"`

	// Pack deployment end time
	// Format: date-time
	EndTime V1Time `json:"endTime,omitempty"`

	// Pack name
	Name string `json:"name,omitempty"`

	// Cluster profile uid
	ProfileUID string `json:"profileUid,omitempty"`

	// Pack deployment start time
	// Format: date-time
	StartTime V1Time `json:"startTime,omitempty"`

	// type
	Type V1PackType `json:"type,omitempty"`

	// pack version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 spectro cluster pack status entity
func (m *V1SpectroClusterPackStatusEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterPackStatusEntity) validateCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	if m.Condition != nil {
		if err := m.Condition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("condition")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterPackStatusEntity) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := m.EndTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("endTime")
		}
		return err
	}

	return nil
}

func (m *V1SpectroClusterPackStatusEntity) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := m.StartTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("startTime")
		}
		return err
	}

	return nil
}

func (m *V1SpectroClusterPackStatusEntity) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterPackStatusEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterPackStatusEntity) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterPackStatusEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
