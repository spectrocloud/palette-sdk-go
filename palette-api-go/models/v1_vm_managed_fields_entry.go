// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMManagedFieldsEntry ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.
//
// swagger:model v1VmManagedFieldsEntry
type V1VMManagedFieldsEntry struct {

	// APIVersion defines the version of this resource that this field set applies to. The format is "group/version" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.
	APIVersion string `json:"apiVersion,omitempty"`

	// FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: "FieldsV1"
	FieldsType string `json:"fieldsType,omitempty"`

	// fields v1
	FieldsV1 *V1VMFieldsV1 `json:"fieldsV1,omitempty"`

	// Manager is an identifier of the workflow managing these fields.
	Manager string `json:"manager,omitempty"`

	// Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.
	Operation string `json:"operation,omitempty"`

	// Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource.
	Subresource string `json:"subresource,omitempty"`

	// time
	// Format: date-time
	Time V1Time `json:"time,omitempty"`
}

// Validate validates this v1 Vm managed fields entry
func (m *V1VMManagedFieldsEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldsV1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMManagedFieldsEntry) validateFieldsV1(formats strfmt.Registry) error {

	if swag.IsZero(m.FieldsV1) { // not required
		return nil
	}

	if m.FieldsV1 != nil {
		if err := m.FieldsV1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fieldsV1")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMManagedFieldsEntry) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := m.Time.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMManagedFieldsEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMManagedFieldsEntry) UnmarshalBinary(b []byte) error {
	var res V1VMManagedFieldsEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
