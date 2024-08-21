// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AppProfile Application profile response
//
// swagger:model v1AppProfile
type V1AppProfile struct {

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// spec
	Spec *V1AppProfileSpec `json:"spec,omitempty"`

	// status
	Status *V1AppProfileStatus `json:"status,omitempty"`
}

// Validate validates this v1 app profile
func (m *V1AppProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppProfile) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1AppProfile) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

func (m *V1AppProfile) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppProfile) UnmarshalBinary(b []byte) error {
	var res V1AppProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1AppProfileSpec v1 app profile spec
//
// swagger:model V1AppProfileSpec
type V1AppProfileSpec struct {

	// Application profile parent profile uid
	ParentUID string `json:"parentUid,omitempty"`

	// template
	Template *V1AppProfileTemplate `json:"template,omitempty"`

	// Application profile version
	Version string `json:"version,omitempty"`

	// Application profile versions list
	Versions []*V1AppProfileVersion `json:"versions"`
}

// Validate validates this v1 app profile spec
func (m *V1AppProfileSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppProfileSpec) validateTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "template")
			}
			return err
		}
	}

	return nil
}

func (m *V1AppProfileSpec) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	for i := 0; i < len(m.Versions); i++ {
		if swag.IsZero(m.Versions[i]) { // not required
			continue
		}

		if m.Versions[i] != nil {
			if err := m.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppProfileSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppProfileSpec) UnmarshalBinary(b []byte) error {
	var res V1AppProfileSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1AppProfileStatus Application profile status
//
// swagger:model V1AppProfileStatus
type V1AppProfileStatus struct {

	// Application profile apps array
	InUseApps []*V1ObjectResReference `json:"inUseApps"`
}

// Validate validates this v1 app profile status
func (m *V1AppProfileStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInUseApps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppProfileStatus) validateInUseApps(formats strfmt.Registry) error {

	if swag.IsZero(m.InUseApps) { // not required
		return nil
	}

	for i := 0; i < len(m.InUseApps); i++ {
		if swag.IsZero(m.InUseApps[i]) { // not required
			continue
		}

		if m.InUseApps[i] != nil {
			if err := m.InUseApps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status" + "." + "inUseApps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppProfileStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppProfileStatus) UnmarshalBinary(b []byte) error {
	var res V1AppProfileStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}